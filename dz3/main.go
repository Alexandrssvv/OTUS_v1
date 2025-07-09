package main

import "fmt"

func ScoringResult(score int) {
	x := ""
	for i := 0; i < score; i++ {
		for j := 0; j < score; j++ {
			if (i+j)%2 == 0 {
				x = x + "#"
			} else {
				x = x + " "
			}
		}
		fmt.Println(x)
		x = ""
	}
}

func main() {
	ScoringResult(8)
}
