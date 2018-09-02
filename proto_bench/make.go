package main

import (
	"io/ioutil"
	"math/rand"
)

func main() {
	array := make([]byte, 500 * 1000000)
	rand.Read(array)
	ioutil.WriteFile("bin", array, 0644)
}
