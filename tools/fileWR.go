package tools

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func WriteFile(name string, content string, batch int) {
	//for each 50 lines, dump results into a single file
	writtingFile, openErr := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if openErr != nil {
		fmt.Printf("open file %d error: ", batch)
		panic(openErr)
	} else {
		defer writtingFile.Close()
	}
	_, newLineErr := writtingFile.WriteString(content)
	if newLineErr != nil {
		fmt.Printf("writting file %d error: ", batch)
	}
	fmt.Println(`written successfully`)
}

func ReadWholeFile(name string) (result string) {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := make([]string, 100)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	result = strings.Join(lines, "\n")
	return
}
