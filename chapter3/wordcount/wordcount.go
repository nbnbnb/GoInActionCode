package main

import (
	"fmt"
	"io/ioutil"

	"goinaction.zhangjin.me/chapter3/words"
)

func main() {
	filename := "gowords.txt"

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("There was an error opening the file:", err)
		return
	}

	text := string(contents)

	count := words.CountWords(text)

	fmt.Printf("There are %d words in your text. \n", count)
}
