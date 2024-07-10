package main

import (
	"fmt"
)

func showHelp() {

}

func showInfo() {
	fmt.Println("这是一个yara规则格式化工具")
	fmt.Println("  1.将yara规则中的中文替换")
	fmt.Println("  2.将字符数超过128字符的进行替换")
}

func main() {
	showInfo()

}
