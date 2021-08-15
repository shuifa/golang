package consistenthash

import (
	"strconv"
	"testing"
)

func TestMap_Get(t *testing.T) {

	hash := NewMap(3, func(data []byte) uint32 {
		atoi, _ := strconv.Atoi(string(data))
		return uint32(atoi)
	})

	hash.Add("6", "4", "2")

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
