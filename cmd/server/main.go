package main

import "fmt"

// Run is going ot be responsible for the instantitation and startup of our go application
func Run() error {
	fmt.Println("Starting up our application")
	return nil
}

func main() {
	fmt.Println("GO REST API Course")
	if err := Run(); err != nil{
		fmt.Println(err)
	}
}