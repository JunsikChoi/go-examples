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
	line string
}

type FindInfo struct {
	filename string
	lines []LineInfo
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Error: Not enough arguments")
		fmt.Println("Usage: wordfinder king '*.txt'")
		return
	}

	word := os.Args[1]
	patterns := os.Args[2:]
	result := []FindInfo{}

	fmt.Println("Target Word: ", word)

	for _, pattern := range patterns {
		result = append(result, FindWordInAllFiles(word, pattern)...)
	}

	PrintResult(result)
}

func GetMatchingFiles(pattern string) ([]string, error){
	return filepath.Glob(pattern)
}

// Find word in all matching files
func FindWordInAllFiles(word string, pattern string) []FindInfo {
	findInfos := []FindInfo{}

	// Find matches with file path
	filenames, err := GetMatchingFiles(pattern)
	if err != nil {
		fmt.Println("Error occurs during finding files matching", pattern)
		return findInfos
	}
	
	ch := make(chan FindInfo)
	cnt := len(filenames)
	recvCnt := 0

	for _, filename := range filenames {
		go FindWordInFile(word, filename, ch)
	}

	for fileInfo := range ch {
		findInfos = append(findInfos, fileInfo)
		recvCnt++
		if cnt == recvCnt {
			break
		}
	}
	return findInfos
}

// Find word in file and return FindInfo
func FindWordInFile(word string, filename string, ch chan FindInfo) {
	findInfo := FindInfo{filename, []LineInfo{}}

	file, err := os.Open(filename)
	if (err != nil) {
		fmt.Println("Can not find file: ", filename)
		ch <- findInfo
	}
	// Close file handle before function ends
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNo := 1
	for scanner.Scan() {
		line := scanner.Text()
		// Search word line by line
		if strings.Contains(line, word) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}
	ch <- findInfo
}

func PrintResult(result []FindInfo) {
	numTotalMatchingLines := 0

	for _, findInfo := range result {
		numTotalMatchingLines += len(findInfo.lines)
	}
	
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Total # of matching lines: ", numTotalMatchingLines)
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println()
	for _, findInfo := range result {
		fmt.Println(findInfo.filename)
		fmt.Println("# of matching lines: ", len(findInfo.lines))
		fmt.Println()
		fmt.Println("----------------------------------------")
		for _, lineInfo := range findInfo.lines {
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("----------------------------------------")
		fmt.Println()
	}
}