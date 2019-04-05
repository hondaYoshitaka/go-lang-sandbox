package main

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	Name string
	Age  int
}

func main() {
	text := `[
		{
			"name": "Taro",
			"age": 20
		},
		{
			"name": "Jiro",
			"age": 15
		}
	]`
	bytes := []byte(text)
	var humans []Human

	if err := json.Unmarshal(bytes, &humans); err != nil {
		panic(err)
	}

	for i := range humans {
		fmt.Printf("氏名: %v, 年齢: %v", humans[i].Name, humans[i].Age)
		fmt.Println()
	}
}
