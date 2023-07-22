package main

import (
	"fmt"
)

func main() {
	fmt.Println("--------- Start ---------")
	test()
	fmt.Println("---------- End ----------")
}

func test() {
	// do test
}

type IP []byte

func (ip IP) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}
