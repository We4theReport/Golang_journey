package main

import "fmt"

func main(){
	//switch statement
	switch {
	case false:
			fmt.Println("this should not print")
	case (2==4):
			fmt.Println("this should not print2")
	case (3==3):
			fmt.Println("prints")
			fallthrough
	case (4==4):
			fmt.Println("this is also true, does it print?!")
	}
}