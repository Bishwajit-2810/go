package file_handling

import (
	"fmt"
	"io"
	"os"
)

func FileHandling() {

	//write in file

	file, err := os.Create("file/welcome.txt")
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer file.Close()
	content := "good morning everyone\ntoday is 22-08-2024"
	// byte,error := io.WriteString(file, content+"\n")
	io.WriteString(file, content+"\n")
	fmt.Println("file created")

	// read from file
	file, error := os.Open("file/welcome.txt")
	if error != nil {
		fmt.Println("error opening file")
		return
	}
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("error while reading", err)
		}
		fmt.Print(string(buffer[:n]))
	}

	content1, error1 := os.ReadFile("file/welcome.txt")
	if error1 != nil {
		fmt.Println("error reading file", err)
		return
	}
	fmt.Print(string(content1))

}

// read happens in 2 way
// 1. buffer
// 2. os.ReadFile/ioutill.ReadFile   read entire file into byte slice
// os/ioutill loads the file in memory which cause huge memory useage
