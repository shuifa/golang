module github.com/oushuifa/golang/better/cache

go 1.16

require github.com/oushuifa/golang/better/cache/geecache v0.0.0

replace (
	github.com/oushuifa/golang/better/cache/consistenthash => ./consistenthash
	github.com/oushuifa/golang/better/cache/geecache => ./geecache
	github.com/oushuifa/golang/better/cache/geecachepb => ./geecachepb
	github.com/oushuifa/golang/better/cache/lru => ./lru
	github.com/oushuifa/golang/better/cache/singleflight => ./singleflight
)
