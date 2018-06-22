package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Provide your passwords as a list of arguments.")
		fmt.Println("eg: bcrypt_hash_password hunter2 password 123456 qwerty")
		return
	}

	for _, arg := range args {
		hash, err := bcrypt.GenerateFromPassword([]byte(arg), 10)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(arg)
		fmt.Println(string(hash))
	}
}
