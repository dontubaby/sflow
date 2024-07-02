package main

import (
	"fmt"
	"os"
)

func main() {

	for {
		fmt.Println("Введите число")
		data := make([]byte, 8)
		n, err := os.Stdin.Read(data)
		if err == nil && n > 0 {
			process(data)
		} else {
			break
		}
	}

}

func process(data []byte) {
	fmt.Printf("Получено: %X 	%s\n", data, string(data))
}
