package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readEntry() { //(float64,[][]float64) {
	file, err := os.Open("./entrada.txt")
	if err != nil {
		panic("Error to open file: entrada.txt")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		lineReader := strings.NewReader(fileScanner.Text())
		var f float64
		for n, err := fmt.Fscan(lineReader, "%f", &f); err == nil; n, err == fmt.Fscan(lineReader, "%f", &f) {
			fmt.Println(f)
		}
	}
}
