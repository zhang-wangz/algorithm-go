package main

// 1146
type SnapshotArray struct {
	id   int
	snap [][][]int
}

func Constructor1(length int) SnapshotArray {
	return SnapshotArray{id: 0, snap: make([][][]int, length)}
}

func (this *SnapshotArray) Set(index int, val int) {
	cur := this.id
	this.snap[index] = append(this.snap[index], []int{cur, val})
}

func (this *SnapshotArray) Snap() int {
	cur := this.id
	this.id++
	return cur
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	snap := this.snap[index]
	l, r := 0, len(snap)
	for l < r {
		mid := l + (r-l)>>1
		if snap[mid][0] <= snap_id {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if l-1 < 0 {
		return 0
	}
	return snap[l-1][1]
}
