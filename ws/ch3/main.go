package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type LineInfo struct {
	lineNo int
	line   string
}

type FindInfo struct {
	filename string
	lines    []LineInfo
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행 인수가 필요합니다.")
		return
	}

	word := os.Args[1]   // 찾으려는 단어
	files := os.Args[2:] // 특정 파일 혹은 패턴 ex) *.txt
	var findInfos []FindInfo
	for _, path := range files {
		findInfos = append(findInfos, findWordInAllFiles(word, path)...)
	}

	for _, findInfo := range findInfos {
		fmt.Println(findInfo.filename)
		fmt.Println("------------------------------")
		for _, lineInfo := range findInfo.lines {
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("------------------------------")
		fmt.Println()
	}

}

func getFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func findWordInAllFiles(word string, path string) []FindInfo {
	var findInfos []FindInfo

	fileList, err := getFileList(path)
	if err != nil {
		fmt.Println("파일을 찾을 수 업습니다. Err : ", err)
		return findInfos
	}

	recvCnt := 0
	c := make(chan FindInfo)
	for _, filename := range fileList {
		go findWordInFile(word, filename, c)
	}

	for info := range c {
		findInfos = append(findInfos, info)
		recvCnt++
		if recvCnt == len(fileList) {
			break
		}
	}

	return findInfos
}

func findWordInFile(word string, filename string, c chan FindInfo) {
	info := FindInfo{filename: filename, lines: []LineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. err : ", err)
		c <- info
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	lineNo := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			info.lines = append(info.lines, LineInfo{lineNo: lineNo, line: line})
		}
		lineNo++
	}

	c <- info
}
