package swordReferoffer

import "math/rand"

// RandomizedSet 每一次移除都最后一个填到前面被删除的index上
type RandomizedSet struct {
	nums    []int
	indices map[int]int
}

// Constructor /** Initialize your data structure here. */
func Constructor2() RandomizedSet {
	return RandomizedSet{
		nums:    make([]int, 0),
		indices: map[int]int{},
	}
}

// Insert /** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (ran *RandomizedSet) Insert(val int) bool {
	if _, ok := ran.indices[val]; ok {
		return false
	}
	le := len(ran.nums)
	ran.nums = append(ran.nums, val)
	ran.indices[val] = le
	return true
}

// Remove /** Removes a value from the set. Returns true if the set contained the specified element. */
func (ran *RandomizedSet) Remove(val int) bool {
	if _, ok := ran.indices[val]; !ok {
		return false
	}
	last := ran.nums[len(ran.nums)-1]
	index := ran.indices[val]
	ran.nums[index] = last
	ran.nums = ran.nums[:len(ran.nums)-1]
	ran.indices[last] = index
	delete(ran.indices, val)
	return true
}

// GetRandom /** Get a random element from the set. */
func (ran *RandomizedSet) GetRandom() int {
	x := rand.Intn(len(ran.nums))
	return ran.nums[x]
}
