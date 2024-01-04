package main

import (
	// пакет используется для проверки ответа, не удаляйте его
	"encoding/json"
	"fmt" // пакет используется для проверки ответа, не удаляйте его
	"os"
)

func main() {
	a, b, c := readTask() // исходные данные получаются с помощью этой функции

	if _, ok := a.(float64); !ok {
		fmt.Printf("value=%v: %T", a, a)
		return
	}

	if _, ok := b.(float64); !ok {
		fmt.Printf("value=%v: %T", b, b)
		return
	}

	if val, ok := c.(string); ok {
		switch val {
		case "+":
			fmt.Printf("%.4f", a.(float64)+b.(float64))
		case "-":
			fmt.Printf("%.4f", a.(float64)-b.(float64))
		case "*":
			fmt.Printf("%.4f", a.(float64)*b.(float64))
		case "/":
			fmt.Printf("%.4f", a.(float64)/b.(float64))
		default:
			fmt.Println("неизвестная операция")
		}
	} else {
		fmt.Println("неизвестная операция")
	}
}

func readTask() (interface{}, interface{}, interface{}) {
	var a, b, c interface{}

	decoder := json.NewDecoder(os.Stdin)

	if err := decoder.Decode(&a); err != nil {
		fmt.Println("Ошибка чтения первого значения:", err)
		os.Exit(1)
	}

	if err := decoder.Decode(&b); err != nil {
		fmt.Println("Ошибка чтения второго значения:", err)
		os.Exit(1)
	}

	if err := decoder.Decode(&c); err != nil {
		fmt.Println("Ошибка чтения третьего значения:", err)
		os.Exit(1)
	}

	return a, b, c
}
