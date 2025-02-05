package main

import "fmt"

/*P11 */

func main() {
	var s = make([]string, 0)
	var m = make(map[string]string)    // 类似python的dict

	s = append(s, "some string")
	m["some key"] = "some value"

	fmt.Println(s)
	fmt.Println(m)

}
