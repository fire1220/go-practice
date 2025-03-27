package go_radixtree

import (
	"fmt"
	"go-radixtree/keywordfilter"
	"testing"
)

func TestKeywordFilter(t *testing.T) {
	tree := keywordfilter.New()
	tree.Insert("你好", 1)
	tree.Insert("你是好人", 1)
	tree.Insert("中国我的国", 2)
	tree.Insert("好的", 1)
	tree.PrintTree()
	if key, val, found := tree.Contains("你你好k中国好的a"); found {
		fmt.Println("true:", key, val)
	} else {
		fmt.Println("false")
	}
}
