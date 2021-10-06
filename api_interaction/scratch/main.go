package main

import "fmt"

func main() {
	key_empty := make([]byte, 32)
	password := "this is test"
	println([]byte(password))
	fmt.Println(key_empty)
	copy(key_empty[:len(password)], []byte(password))
	fmt.Println(key_empty)
}
