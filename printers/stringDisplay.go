package printers

import "fmt"

// Display takes a stirng
func Display(value, title string) {
	fmt.Printf("====================== %s ======================", title)
	fmt.Println(value)
	fmt.Println("========================================================")
}
