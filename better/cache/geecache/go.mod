module github.com/oushuifa/golang/better/cache/geecache

go 1.16

require (
	github.com/oushuifa/golang/better/cache/consistenthash v0.0.0
	github.com/oushuifa/golang/better/cache/geecachepb v0.0.0
	github.com/oushuifa/golang/better/cache/lru v0.0.0
	github.com/oushuifa/golang/better/cache/singleflight v0.0.0
	google.golang.org/protobuf v1.27.1
)

replace (
	github.com/oushuifa/golang/better/cache/consistenthash => ../consistenthash
	github.com/oushuifa/golang/better/cache/geecachepb => ../geecachepb
	github.com/oushuifa/golang/better/cache/lru => ../lru
	github.com/oushuifa/golang/better/cache/singleflight => ../singleflight
)
