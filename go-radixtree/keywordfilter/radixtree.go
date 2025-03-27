package keywordfilter

type node struct {
	category int
	children map[rune]*node
}

type KeywordFilter struct {
	root *node
}

func New() *KeywordFilter {
	return &KeywordFilter{root: &node{children: make(map[rune]*node)}}
}

func (t *KeywordFilter) Insert(str string, category int) {
	if len(str) == 0 || category == 0 {
		return
	}
	n := t.root
	for _, one := range str {
		if _, ok := n.children[one]; !ok {
			n.children[one] = &node{children: make(map[rune]*node)}
		}
		n = n.children[one]
	}
	n.category = category
}

func (t *KeywordFilter) Contains(str string) (string, int, bool) {
	for k := range str {
		key := str[k:]
		if keyword, category, ok := t.prefixSearch(key); ok {
			return keyword, category, ok
		}
	}
	return "", 0, false
}

func (t *KeywordFilter) prefixSearch(str string) (string, int, bool) {
	n := t.root
	k := 0
	for i, one := range str {
		if _, ok := n.children[one]; !ok {
			k += i
			break
		}
		n = n.children[one]
	}
	if n.category == 0 {
		return "", 0, false
	}
	if k == 0 {
		k = len(str)
	}
	return str[0:k], n.category, true
}
