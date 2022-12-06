package shared

import "github.com/sbabiv/roundrobin"

type Balancer struct {
	inner *roundrobin.Balancer
}

func (p *Balancer) Pick() (interface{}, error) {
	return p.inner.Pick()
}

func (p *Balancer) MustPick() interface{} {
	item, err := p.inner.Pick()
	if err != nil {
		panic(err.Error())
	}

	return item
}

func (p *Balancer) MustPickString() string {
	if s, ok := p.MustPick().(string); ok {
		return s
	}

	panic("picked item is not a string")
}

func NewBalancer(items ...interface{}) *Balancer {
	return &Balancer{
		inner: roundrobin.New(items),
	}
}

func NewStringBalancer(items ...string) *Balancer {
	data := make([]interface{}, len(items))
	for idx, item := range items {
		data[idx] = item
	}
	return NewBalancer(data...)
}
