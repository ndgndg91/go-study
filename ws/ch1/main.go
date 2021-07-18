package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행인수가 필요합니다.")
		return
	}

	word := os.Args[1]
	files := os.Args[2:]
	fmt.Println("찾으려는 단어 : ", word)
	printAllFiles(files)
}

func printAllFiles(files []string) {
	for _, path := range files {
		fileList, err := getFiles(path) // 파일목록 가져오기
		if err != nil {
			fmt.Println("파일을 찾을 수 없습니다. err : ", err)
			return
		}

		fmt.Println("찾으려는 파일 리스트")
		for _, name := range fileList {
			fmt.Println(name)
		}
	}
}

func getFiles(path string) ([]string, error) {
	return filepath.Glob(path)
}
