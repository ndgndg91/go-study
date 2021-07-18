package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	printFile("example.txt")
}

func printFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다.", fileName)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file) // 스캐너를 통해서 한 줄 씩 읽기
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
