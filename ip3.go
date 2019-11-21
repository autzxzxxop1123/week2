package main 

import "fmt"

func main() {
	n, e := fmt.Print("Hello", "Word", 123, 456, 123, "\n")
	fmt.Print("number of bytes written :", n, "\n")
	fmt.Print("write error encountered :", e, "\n")


	

	a, b := fmt.Println("Hello","Word", 123, 456, )
	fmt.Println("number of bytes written :", a,)
	fmt.Println("write error encountered :", b,)
}