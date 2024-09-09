package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func BeautifulMatrix() {
	scanner := bufio.NewScanner(os.Stdin)
	input := ""
	inputCount := 0
	idxOne := 0

	for i := 0; i < 5; i++ {
		scanner.Scan()
		input = scanner.Text()
		input = strings.TrimSpace(input)

		if strings.Contains(input, "1") {
			nums := strings.Split(input, " ")
			for j, num := range nums {
				if num == "1" {
					idxOne = j
					inputCount = i
					break
				}
			}
			break
		}
	}

	f := math.Abs(2 - float64(inputCount))
	s := math.Abs(2 - float64(idxOne))
	r, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", f+s), 64)
	fmt.Println(r)
}
