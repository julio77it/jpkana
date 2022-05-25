package main

import (
	"fmt"
	"io/ioutil"
	"jpkana"
)

func main() {
	hbytes, err := ioutil.ReadFile("../resources/hiragana.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	kbytes, err := ioutil.ReadFile("../resources/katakana.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	generator, err := jpkana.New(hbytes, kbytes)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(generator)
}
