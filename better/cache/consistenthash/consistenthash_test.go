package consistenthash

import (
	"fmt"
	"sort"
	"testing"
)

func TestMap_Get(t *testing.T) {

	hash := NewMap(200, nil)

	nodes := []string{
		"172.168.92.13:9090",
		"195.125.78.222:18625",
		"200.8.96.48:53698",
		"123.10.231.189:65512",
		"189.20.10.13:5678",
		"125.234.45.73:19456",
		"45.189.177.179:23564",
		"222.127.122.219:36547",
		"173.163.60.59:45263",
		"243.145.78.67:59874",
		"200.236.203.153:29846",
		"99.8.165.58:57891",
	}


	hash.Add(nodes...)

	var relations = make(map[string][]int)

	for _, key := range hash.keys {
		relations[hash.hashMap[key]] = append(relations[hash.hashMap[key]], key)
	}

	var balances = make(map[string]map[string]int)

	for node, keys := range relations {
		tmp := make(map[string]int, 10)
		for i := 0; i < len(keys); i++ {
			idx := sort.SearchInts(hash.keys, keys[i])
			// fmt.Println(node, hash.hashMap[hash.keys[(idx + 1)%len(hash.keys)]])
			for hash.hashMap[hash.keys[(idx + 1)%len(hash.keys)]] == node {
				idx++
			}
			tmp[hash.hashMap[hash.keys[(idx + 1)%len(hash.keys)]]]++
		}
		// return
		balances[node] = tmp
	}

	for s, i := range balances["195.125.78.222:18625"] {
		fmt.Println(s, i)
	}
	// fmt.Println(balances["172.168.92.13:9090"])



	testCases := map[string]string{
		"2":  "2",
		"11": "2",
		"23": "4",
		"27": "2",
	}

	t.Run("consisitenthash_getnode", func(t *testing.T) {
		for k, tc := range testCases {
			if hash.Get(k) != tc {
				t.Errorf("Asking for %s, should have yielded %s", k, tc)
			}
		}
	})

	hash.Add("8")
	testCases["27"] = "8"

	t.Run("consisitenthash_addnode", func(t *testing.T) {
		for k, tc := range testCases {
			if hash.Get(k) != tc {
				t.Errorf("Asking for %s, should have yielded %s", k, tc)
			}
		}
	})

}
