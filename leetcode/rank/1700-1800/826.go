package main

import "sort"

// 826. 安排工作以达到最大收益
func maxProfitAssignment(difficulty []int, profit []int, worker []int) (ans int) {
	type pair struct {
		hard, pro int
	}
	jobs := []pair{}
	for i := range difficulty {
		jobs = append(jobs, pair{difficulty[i], profit[i]})
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].hard < jobs[j].hard
	})
	sort.Ints(worker)
	i, best := 0, 0
	for _, skill := range worker {
		for i < len(jobs) && jobs[i].hard <= skill {
			if best < jobs[i].pro {
				best = jobs[i].pro
			}
			i++
		}
		ans += best
	}
	return
}
