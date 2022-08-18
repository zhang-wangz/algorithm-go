package _022_07

import (
	"math/rand"
)

const MAXLEVEL, P = 32, 0.25

type SkipNode struct {
	Next []*SkipNode
	Val  int
}
type Skiplist struct {
	Head  *SkipNode
	Level int
}

func Constructor10() Skiplist {
	return Skiplist{&SkipNode{make([]*SkipNode, MAXLEVEL), -1}, 0}
}

func (this *Skiplist) Search(target int) bool {
	cur := this.Head
	for i := this.Level - 1; i >= 0; i-- {
		for cur.Next[i] != nil && cur.Next[i].Val < target {
			cur = cur.Next[i]
		}
	}
	cur = cur.Next[0]
	return cur != nil && cur.Val == target
}
func (this *Skiplist) randomLevel() int {
	lv := 1
	for lv < MAXLEVEL && rand.Float64() < P {
		lv++
	}
	return lv
}

func (this *Skiplist) Add(num int) {
	update := make([]*SkipNode, MAXLEVEL)
	for i := 0; i < len(update); i++ {
		update[i] = this.Head
	}
	cur := this.Head
	for i := this.Level - 1; i >= 0; i-- {
		for cur.Next[i] != nil && cur.Next[i].Val < num {
			cur = cur.Next[i]
		}
		update[i] = cur
	}
	lv := this.randomLevel()
	this.Level = max(lv, this.Level)
	newNode := &SkipNode{Val: num, Next: make([]*SkipNode, lv)}
	for i := lv - 1; i >= 0; i-- {
		newNode.Next[i] = update[i].Next[i]
		update[i].Next[i] = newNode
	}
}

func (this *Skiplist) Erase(num int) bool {
	// 都设置为最大值，因为生成的level不知道有多大
	update := make([]*SkipNode, MAXLEVEL)
	for i := 0; i < len(update); i++ {
		update[i] = this.Head
	}
	cur := this.Head
	for i := this.Level - 1; i >= 0; i-- {
		for cur.Next[i] != nil && cur.Next[i].Val < num {
			cur = cur.Next[i]
		}
		update[i] = cur
	}
	cur = cur.Next[0]
	if cur == nil || cur.Val != num {
		return false
	}
	for i := 0; i < len(update) && update[i].Next[i] == cur; i++ {
		update[i].Next[i] = cur.Next[i]
	}
	// 最小设置为1， 不然后面就成负数了
	for this.Level > 1 && this.Head.Next[this.Level-1] == nil {
		this.Level--
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
