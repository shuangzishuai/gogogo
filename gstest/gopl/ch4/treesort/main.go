package main

import (
	"fmt"
	"time"
)

type tree struct {
	value int
	left  *tree
	right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t == nil {
		return values
	}
	values = appendValues(values, t.left)
	values = append(values, t.value)
	values = appendValues(values, t.right)
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {
	t := daysAgo(time.Now().AddDate(0, 0, -5))     //5天之前
	t1 := daysAgo(time.Now().Add(-time.Hour * 48)) // 2天前
	fmt.Printf("%d days ago\n", t)
	fmt.Printf("%d days ago\n", t1)
}
