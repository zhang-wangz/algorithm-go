package _022_06

//  https://leetcode.cn/problems/duplicate-zeros/
// 给你一个长度固定的整数数组arr，请你将该数组中出现的每个零都复写一遍，并将其余的元素向右平移。

// 注意：请不要在超过该数组长度的位置写入元素。

// 要求：请对输入的数组就地进行上述修改，不要从函数返回任何东西。

// 模拟栈
func DuplicateZeros(arr []int) {
	stack := make([]int, 0)
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			stack = append(stack, 0)
		}
		stack = append(stack, arr[i])
		if len(stack) == n {
			break
		}
	}
	for i := n - 1; i >= 0; i-- {
		arr[i] = stack[i]
	}
}
