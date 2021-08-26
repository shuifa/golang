package geecache

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/oushuifa/golang/better/cache/consistenthash"
	pb "github.com/oushuifa/golang/better/cache/geecachepb"
	"google.golang.org/protobuf/proto"
)

const (
	defaultBasePath = "/_geecache/"
	defaultReplicas = 50
)

type HttpPool struct {
	self        string
	basePath    string
	mtx         sync.Mutex
	peers       *consistenthash.Map
	HttpGetters map[string]*HttpGetter
}

type HttpGetter struct {
	baseUrl string
}

func NewHttpPool(self string) *HttpPool {
	return &HttpPool{
		self:     self,
		basePath: defaultBasePath,
	}
}

func (p *HttpPool) Log(format string, value ...interface{}) {
	log.Printf("[server]%s, %s", p.self, fmt.Sprintf(format, value...))
}

func (p *HttpPool) Set(peers ...string) {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	p.peers = consistenthash.NewMap(defaultReplicas, nil)

	p.peers.Add(peers...)

	p.HttpGetters = make(map[string]*HttpGetter, len(peers))

	for _, peer := range peers {
		p.HttpGetters[peer] = &HttpGetter{baseUrl: peer + p.basePath}
	}
}

func (p *HttpPool) PickPeer(key string) (peer PeerGetter, ok bool) {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	if node := p.peers.Get(key); node != "" && node != p.self {
		p.Log("Pick peer %s", node)
		return p.HttpGetters[node], true
	}

	return nil, false
}

func (p *HttpPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if !strings.HasPrefix(r.URL.Path, p.basePath) {
		panic("HTTPPool serving unexpected path: " + r.URL.Path)
	}

	p.Log("method=%s, path=%s", r.Method, r.URL.Path)

	parts := strings.SplitN(r.URL.Path[len(p.basePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	group := GetGroup(parts[0])
	key := parts[1]
	if group == nil {
		http.Error(w, "no such group: "+parts[0], http.StatusNotFound)
		return
	}

	byteView, err := group.Get(key)

	body, err := proto.Marshal(&pb.Response{Value: byteView.ByteSlice()})
	if err != nil {
		p.Log(http.StatusText(http.StatusInternalServerError)+", when get value, key=%s, err=%s", key, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")

	_, err = w.Write(body)
	if err != nil {
		p.Log("write response bytes err, err=%s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func (h *HttpGetter) Get(in *pb.Request, out *pb.Response) error {

	uri := fmt.Sprintf("%v%v/%v", h.baseUrl, url.QueryEscape(in.GetGroup()), url.QueryEscape(in.GetKey()))

	response, err := http.Get(uri)
	if err != nil {
		log.Fatalf("curl get err, err=%s", err.Error())
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Printf("read body err, err=%s", response.Status)
		return fmt.Errorf("server return %v", response.Status)
	}

	dataBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("read body err, err=%s", err.Error())
		return err
	}

	if err := json.Unmarshal(dataBytes, &response); err != nil {
		return fmt.Errorf("decoding response body: %v", err)
	}

	return nil
}


