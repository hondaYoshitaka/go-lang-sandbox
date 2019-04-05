package main

import (
	"encoding/json"
	"fmt"
)

type HumanTag struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
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
	var humans []HumanTag

	if err := json.Unmarshal(bytes, &humans); err != nil {
		panic(err)
	}

	for i := range humans {
		fmt.Printf("氏名: %v, 年齢: %v", humans[i].Name, humans[i].Age)
		fmt.Println()
	}
}
