package main

import (
	"os"
	"os/exec"
	"syscall"
)

//注意 Go 没有提供 Unix 经典的 fork 函数。 一般来说，这没有问题，因为启动协程、生成进程和执行进程， 已经涵盖了 fork 的大多数使用场景。

func main() {
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-alh"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
