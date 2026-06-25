package main

import "fmt"

const LoginToken string = "super secret login token " // public

func main() {
	var username string = "ramanan"
	fmt.Println(username)
	fmt.Printf("variable type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable type: %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("variable type: %T \n", smallValue)

	var notSmallValue uint = 256
	fmt.Println(notSmallValue)
	fmt.Printf("variable type: %T \n", notSmallValue)

	var smallFloat float32 = 255.343423434534543543
	fmt.Println(smallFloat)
	fmt.Printf("variable type: %T \n", smallFloat)

	var notSmallFloat float64 = 255.343423434534543543
	fmt.Println(notSmallFloat)
	fmt.Printf("variable type: %T \n", notSmallFloat)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable type: %T \n", notSmallFloat)

	website := "google.com"
	fmt.Println(website)
	fmt.Printf("variable type: %T \n", notSmallFloat)
}
