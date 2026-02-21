package main

import (
	"fmt"
	"os"
)

func errCheck(err error){
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open("file.txt")

	errCheck(err)

	fileInfo, err := f.Stat()

	errCheck(err)

	// fmt.Println("File Name: ",fileInfo.Name())
	// fmt.Println("File or Folder: ",fileInfo.IsDir())
	// fmt.Println("File Size:",fileInfo.Size())
	// fmt.Println("File Mode:",fileInfo.Mode())
	// fmt.Println("File Modified Time:",fileInfo.ModTime())

	//Read File
	size := fileInfo.Size()
	buf := make([]byte, size)

	d,err := f.Read(buf)

	errCheck(err)

	fmt.Println(d)

	fmt.Print("data: ")
	for i:=range buf{
		print(string(buf[i]))
	}
	defer f.Close()


	//Another way of reading data
	
}