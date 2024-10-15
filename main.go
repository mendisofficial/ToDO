package main

func main() {
	todos := Todos{}
	todos.add("Buy milk")
	todos.add("Buy bread")
	err := todos.toggle(0)
	if err != nil {
		return
	}
	todos.print()
}
