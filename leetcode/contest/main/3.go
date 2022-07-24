package main

import (
	"container/heap"
	"sort"
)

type pair struct {
	name string
	val  int
}

type FoodRatings struct {
	cuisines     map[string]*minHeap
	foodcuisines map[string]string
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	cuisine := map[string]*minHeap{}
	foodcuisines := map[string]string{}
	for i := range cuisines {
		if cuisine[cuisines[i]] == nil {
			cuisine[cuisines[i]] = &minHeap{pair: make([]pair, 0)}
		}
		heap.Push(cuisine[cuisines[i]], pair{name: foods[i], val: ratings[i]})
		foodcuisines[foods[i]] = cuisines[i]
	}
	return FoodRatings{cuisines: cuisine, foodcuisines: foodcuisines}
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	idx := sort.Search(this.cuisines[this.foodcuisines[food]].Len(), func(i int) bool {
		return this.cuisines[this.foodcuisines[food]].pair[i].name == food
	})
	this.cuisines[this.foodcuisines[food]].pair[idx].val = newRating
	heap.Fix(this.cuisines[this.foodcuisines[food]], idx)
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	return this.cuisines[cuisine].pair[0].name
}

type minHeap struct {
	pair []pair
}

func (h *minHeap) Len() int {
	return len(h.pair)
}
func (h *minHeap) Push(i interface{}) {
	h.pair = append(h.pair, i.(pair))
}
func (h *minHeap) Pop() interface{} {
	top := h.pair[len(h.pair)-1]
	h.pair = h.pair[:len(h.pair)-1]
	return top
}
func (x *minHeap) Less(i, j int) bool {
	if x.pair[i].val == x.pair[j].val {
		return x.pair[i].name < x.pair[j].name
	}
	return x.pair[i].val > x.pair[j].val
}
func (x *minHeap) Swap(i, j int) { x.pair[i], x.pair[j] = x.pair[j], x.pair[i] }

/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */
