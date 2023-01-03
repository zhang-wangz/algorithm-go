package _022_08

type OrderedStream struct {
	stream []string
	ptr    int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{make([]string, n+1), 1}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.stream[idKey] = value
	start := this.ptr
	for this.ptr < len(this.stream) && this.stream[this.ptr] != "" {
		this.ptr++
	}
	return this.stream[start:this.ptr]
}
