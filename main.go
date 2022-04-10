package main

import (
	"fmt"
	"gompiler/repl"
	"os"
	"os/user"
)

func main(){
	user, err := user.Current()		// current user for os
	if err != nil{
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.				Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}