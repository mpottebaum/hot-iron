package main

import (
	"fmt"
)

func main() {
	fmt.Printf("\n\n\n\n")
	fmt.Printf("what up my dude\n")
	fmt.Printf("welcome to hot iron\n\n")
	fmt.Printf("made in loving memory of jeremy irons\n\n")
	fmt.Printf("where's the beef???\n\n")
	fmt.Printf("\n\n\n\n")
	var tmplFilePath string
	fmt.Printf("enter the file path for your home page --> ")
	fmt.Scan(&tmplFilePath)
	fmt.Printf("\n\n")
	fmt.Printf("%s\n\n", tmplFilePath)
}
