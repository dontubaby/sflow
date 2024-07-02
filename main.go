package main
 
import (
    "fmt"
    "log"
)
 
func main() {
    n := 0
    fmt.Print("Введите целое число: ")
    _, err := fmt.Scan(&n)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Вы ввели число: %d\n", n)
}

func process(data []byte) {
	fmt.Printf("Получено: %X 	%s\n", data, string(data))
}
