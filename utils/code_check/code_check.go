package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	// 调用 Simian 命令行工具，并传递要比较的文件路径
	cmd := exec.Command("simian", "-reportDuplicateText", "-reportFile=report.txt", "path/to/file1", "path/to/file2")

	// 捕获命令的标准输出
	output, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	// 启动命令
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	// 读取命令的输出
	scanner := bufio.NewScanner(output)
	for scanner.Scan() {
		line := scanner.Text()

		// 解析 Simian 的输出，获取相似度分数
		if strings.HasPrefix(line, "Found a duplicate") {
			scoreStr := strings.Split(line, " ")[4]
			score, err := strconv.ParseFloat(strings.TrimSuffix(scoreStr, "%"), 64)
			if err != nil {
				panic(err)
			}

			// 判断相似度分数是否超过阈值
			if score >= 80 {
				fmt.Println("Files are similar")
			} else {
				fmt.Println("Files are different")
			}
			break
		}
	}

	// 等待命令结束
	if err := cmd.Wait(); err != nil {
		panic(err)
	}
}
