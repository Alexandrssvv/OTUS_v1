package main

import "fmt"

func GenerateBoard(size int) {
	row := ""
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				row = row + "#"
			} else {
				row = row + " "
			}
		}
		fmt.Println(row)
		row = ""
	}
}

func main() {
	var size int
	fmt.Print("Введите размер доски: ")
	_, err := fmt.Scanln(&size)
	if err != nil || size <= 0 {
		fmt.Println("Ввод некорректен. Используется размер по умолчанию: 8")
		size = 8
	}
	GenerateBoard(size)
}
