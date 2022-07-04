package main

import "sort"

// 计数也可以，记录在arr1中每个数字出现的次数，然后append到一个答案里
func relativeSortArray1(arr1 []int, arr2 []int) []int {
	mp := map[int]int{}
	for i := 0; i < len(arr2); i++ {
		mp[arr2[i]] = i
	}
	sort.Slice(arr1, func(i, j int) bool {
		a, b := arr1[i], arr1[j]
		v1, ok1 := mp[a]
		v2, ok2 := mp[b]
		if ok1 && ok2 {
			return v1 < v2
		}
		// 返回在hash里面的那个
		// 如果是ok1 true，返回，如果是ok2 true， ok1就是false，加入ok2对应的v
		if ok1 || ok2 {
			return ok1
		}
		return a < b
	})
	return arr1
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	mp := map[int]int{}
	for i := 0; i < len(arr2); i++ {
		mp[arr2[i]] = i
	}
	nums1 := make([]int, 0)
	nums2 := make([]int, 0)
	for i := 0; i < len(arr1); i++ {
		if _, ok := mp[arr1[i]]; !ok {
			nums2 = append(nums2, arr1[i])
		} else {
			nums1 = append(nums1, arr1[i])
		}
	}
	sort.Slice(nums1, func(i, j int) bool {
		return mp[nums1[i]] < mp[nums1[j]]
	})
	sort.Ints(nums2)
	return append(nums1, nums2...)
}
