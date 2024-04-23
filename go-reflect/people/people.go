package people

type People struct {
	Name any `json:"name"`
	Nick any `json:"nick"`
	like any `json:"like"`
}

func (p *People) SetLike(k any) {
	p.like = k
}
func (p *People) GetLike() any {
	return p.like
}
