package main

import "fmt"

func main() {
	cache := make(map[string]string)

	cache["name"] = "TesterCC"
	cache["job"] = "Developer"

	fmt.Println(cache)
}
