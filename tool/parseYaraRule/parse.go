package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ruleNum = 0

func replace() string {
	ruleNum++
	ret := "rule rule_replace_" + strconv.Itoa(ruleNum) + "{"
	return ret
}

func parse(srcFile, dstFile string) {
	file, err := os.Open(srcFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	outFile, err := os.Create(dstFile)
	if err != nil {
		fmt.Println("Error creating target file:", err)
		return
	}
	defer outFile.Close()
	writer := bufio.NewWriter(outFile)

	scanner := bufio.NewScanner(file)

	target := "rule "
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		if strings.Contains(line, target) && strings.Contains(line, "{") {
			str := replace()
			line = str
		}
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to target file:", err)
			return
		}
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}
}
