package main

import "fmt"

func main() {
	todos := Todos{}
	todos.add("Buy milk")
	todos.add("Buy bread")
	fmt.Printf("%+v\n\n", todos)
	err := todos.delete(0)
	if err != nil {
		return
	}
	fmt.Printf("%+v\n\n", todos)
}
