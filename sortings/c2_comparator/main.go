package main

import (
	"fmt"
	"sort"
)

type Points struct {
	L int
	R int
}

func main() {
	// Custom types
	points := []Points{
		{1, 2}, {2, 3}, {4, 5}, {6, 9}, {3, 7}, {2, 8}, {6, 7},
	}

	sort.SliceStable(points, func(i, j int) bool {
		return points[i].L < points[j].L || (points[i].L == points[j].L && points[i].R < points[j].R)
	})

	fmt.Println(points)

	// Base on priority array
	a := []int{1, 6, 4, 5, 5, 4, 3, 2, 7}
	prior := map[int]int{
		1: 7,
		2: 6,
		3: 1,
		4: 2,
		5: 4,
		6: 5,
		7: 3,
	}

	sort.SliceStable(a, func(i, j int) bool {
		return prior[a[i]] > prior[a[j]] || (a[i] == a[j] && i > j)
	})

	fmt.Println(a)

	// Sort scala descreasingly
	x := []int{1, 2, 6, 7, 3, 4, 2, 5, 1}
	sort.SliceStable(x, func(i, j int) bool {
		return x[i] > x[j]
	})
	fmt.Println(x) // [7 6 5 4 3 2 2 1 1]
}
