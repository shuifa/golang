module github.com/oushuifa/golang/better/cache

go 1.16

require (
	github.com/oushuifa/golang/better/cache/geecache v0.0.0
	github.com/oushuifa/golang/better/cache/lru v0.0.0
	github.com/oushuifa/golang/better/cache/consistenthash v0.0.0
)

replace (
	github.com/oushuifa/golang/better/cache/geecache => ./geecache
	github.com/oushuifa/golang/better/cache/lru => ./lru
	github.com/oushuifa/golang/better/cache/consistenthash => ./consistenthash
)