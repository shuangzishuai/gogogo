package main

import (
	"fmt"
	"testing"
)

// testing 包为提供了编写单元测试所需的工具，写好单元测试后，我们可以通过 go test 命令运行测试
// 际上，单元测试的代码可以位于任何包下。测试代码通常与需要被测试的代码位于同一个包下。
//go test -v
//go test -bench=.

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b

}

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		//t.Error* 会报告测试失败的信息，然后继续运行测试。
		//t.Fatal* 会报告测试失败的信息，然后立即终止测试。
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{2, -2, -2},
		{-2, 2, -2},
		{2, 2, 2},
		{-2, -2, -2},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d, %d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

}

func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(2, -2)
	}
}

func main() {

}
