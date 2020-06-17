package main

import (
	"encoding/json"
	"fmt"
)

//P19
type Foo struct {
	Bar string `xml:"id,attr"`
	Baz string `xml:"parent>child"`
}

func main() {
	f := Foo{"Joe Junior", "Hello Senior"}
	b, _ := json.Marshal(f) // Marshal() method encodes the struct to JSON, return a byte slice
	fmt.Println(string(b))  // a JSON-encoded string representation
	json.Unmarshal(b, &f)
}
