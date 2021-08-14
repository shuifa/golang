package geecache

type ByteView struct {
    b []byte
}

func (v ByteView) Len() int {
    return len(v.b)
}

func (v ByteView) ByteSlice() []byte {
    return v.cloneBytes()
}

func(v ByteView) String() string {
    return string(v.b)
}

func(v ByteView) cloneBytes() []byte {
    c := make([] byte, len(v.b))
    copy(c, v.b)
    return c
}