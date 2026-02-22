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
	data,err := os.ReadFile("file.txt")
	errCheck(err)
	fmt.Println("\nData:",string(data))

	//Read Folders
	dir,err := os.Open(".")
	// dir,err := os.Open("../")

	errCheck(err)

	defer dir.Close()

	fileInfos, err := dir.ReadDir(-1)
	errCheck(err)

	fmt.Println()
	fmt.Println()
	for _,fi :=range fileInfos{
		fmt.Print(fi.Name())
		if(fi.IsDir()){
			fmt.Print(" -  This is a directory!!")
		}
		fmt.Println()
	}


	//Write a file
	file,err := os.Create("example.txt")	//creating a new file
	errCheck(err)
	defer file.Close()

	file.WriteString("Hello ! i m under water, please help!!")
}