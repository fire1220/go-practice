package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// 要执行的 shell 命令
	cmd := "ls -l"

	// 使用 exec.Command 函数创建一个 Cmd 结构体
	// 第一个参数是要执行的命令，后面的参数是命令的参数列表（如果有）
	// 这里的示例是执行 "ls -l" 命令
	command := exec.Command("sh", "-c", cmd)

	// 使用 CombinedOutput 方法执行命令并等待其完成
	// CombinedOutput 方法返回命令的标准输出和标准错误输出合并后的字节切片
	output, err := command.CombinedOutput()
	if err != nil {
		fmt.Println("命令执行失败:", err)
		return
	}

	// 打印命令输出结果
	fmt.Println(string(output))
}
