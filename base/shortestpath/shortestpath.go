package main

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

const N, M = 110, 6010

var vis [N]bool
var w [N][N]int

// 单源点的时候
var dist [N]int
var k int

// 1e9
const INF = 0x3f3f3f3f

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// 如果是邻接表
var he [N]int
var e [M]int
var ne [M]int
var weight [M]int
var idx int

func add(a, b, w int) {
	e[idx] = b
	// 边
	ne[idx] = he[a]
	// 节点
	he[a] = idx
	weight[idx] = w
	idx++
}
func Floyd() {
	// i到j的距离
	var dist [N][N]int
	for p := 1; p <= N; p++ {
		for i := 1; i <= N; i++ {
			for j := 1; j <= N; j++ {
				dist[i][j] = min(dist[i][j], dist[i][p]+dist[p][j])
			}
		}
	}
}

func dijkstra() {
	// k是起点
	k = 5
	vis = [N]bool{}
	for i := 0; i <= N; i++ {
		dist[i] = INF
	}
	dist[k] = 0
	for i := 1; i <= N; i++ {
		t := -1
		// 找出距离最小的一个节点
		for p := 1; p <= N; p++ {
			if !vis[p] && (t == -1 || dist[p] < dist[t]) {
				t = p
			}
		}
		vis[t] = true
		// 松弛操作
		for i := 1; i <= N; i++ {
			dist[i] = min(dist[i], dist[t]+w[t][i])
		}
	}
}

type pq [][]int

func (h pq) Len() int           { return len(h) }
func (h pq) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h pq) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h pq) Push(v interface{}) { h = append(h, v.([]int)) }
func (h pq) Pop() interface{}   { top := h[len(h)-1]; h = h[:len(h)-1]; return top }
func Cnt2() {
	q := priorityqueue.NewWith(func(a, b interface{}) int {
		x := a.([]int)
		y := b.([]int)
		return utils.IntComparator(x[1], y[1])
	})
	k = 5
	vis = [N]bool{}
	//q := &pq{}
	vis = [N]bool{}
	for i := 0; i <= N; i++ {
		dist[i] = INF
	}
	for i := 0; i <= N; i++ {
		he[i] = -1
	}
	dist[k] = 0
	//heap.Push(q, []int{k, 0})
	q.Enqueue([]int{k, 0})
	for q.Size() != 0 {
		//node := heap.Pop(q).([]int)
		node, _ := q.Dequeue()
		p := node.([]int)[0]
		if vis[p] {
			continue
		}
		vis[p] = true
		for i := he[p]; i != -1; i = ne[i] {
			// p-j 的边，j是顶点
			j := e[i]
			if dist[j] > dist[p]+weight[i] {
				dist[j] = dist[p] + weight[i]
				//heap.Push(q, []int{j, dist[j]})
				q.Enqueue([]int{j, dist[j]})
			}
		}
	}
}
