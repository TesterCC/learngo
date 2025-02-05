package main

import (
	"fmt"
	"strings"
)

// ref: http://govspy.peterbe.com/#lists

func main() {
	// initialized array
	var numbers [5]int // becomes [0, 0, 0, 0, 0]
	// change one of them
	numbers[2] = 100
	//fmt.Println(numbers) // [0, 0, 100, 0, 0]

	// create a new slice from an array
	some_numbers := numbers[1:3]
	fmt.Println(some_numbers) // [0, 100]

	// length of it
	fmt.Println(len(numbers))

	// initialize a slice
	var scores []float64
	scores = append(scores, 1.1) // recreate to append
	scores[0] = 2.2              // change your mind
	fmt.Println(scores)          // prints [2.2]

	// when you don't know for sure how much you're going
	// to put in it, one way is to
	var things [100]string
	things[0] = "Peter"
	things[1] = "Anders"
	fmt.Println(len(things)) // 100
	//fmt.Println(things) // [Peter Anders            ]                                                                                               ]

	// custom test write list
	fmt.Println("-----------------------")
	var scores2 []float64
	scores2 = append(scores2, 1.1)
	fmt.Println(scores2)
	scores2 = append(scores2, 2.2)
	fmt.Println(scores2)
	scores2 = append(scores2, 3.3)
	fmt.Println(scores2)
	scores2[2] = 6.6
	fmt.Println(scores2)

	longStr := strings.Repeat("-",33)
	fmt.Println(longStr)
	var things2 [3]string
	things2[0] = "Alice"
	things2[1] = "Bob"
	things2[2] = "Chris"
	fmt.Println(len(things2))
	fmt.Println(things2)   // [Alice Bob Chris] not [Alice Bob Chris      ]

	var things3 []string
	things3 = append(things3, "David")
	things3 = append(things3, "Frank")
	fmt.Println(len(things3))
	fmt.Println(things3)
}
