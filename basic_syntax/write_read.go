package main

import (
	"fmt"
	"os"
)

func RandW() {
	var file *os.File
	var err error

	// open the file in read-write mode
	file, err = os.OpenFile("./file.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("File opened successfully")
	}
	defer file.Close() // defer is used to close the file after the function ends

	// write to the file
	_, err = file.WriteString("Hello World\n")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data written successfully")
	}
}
