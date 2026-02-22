package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//Cheaks error by checking not nill. Receiving parameter error(error type)
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
	file.WriteString("I am drwaning....")


	bytes := []byte("hello go lang!")

	n,err := file.Write(bytes)

	errCheck(err)

	fmt.Println("Returend n: ",n)


	//read and write to another file (streaming fashion)

	sourceFile,err := os.Open("example.txt")
	errCheck(err)

	defer sourceFile.Close()

	destFile, err := os.Create("example2.txt")

	errCheck(err)
	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for{
		b,err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}
		err2 := writer.WriteByte(b)
		errCheck(err2)
	}

	writer.Flush()
	fmt.Println("Written to new file successfully!")


	//Another way of copying
	sourceFile2, err := os.Open("file.txt")
	errCheck(err)
	defer sourceFile2.Close()

	destFile2,err := os.Create("example3.txt")
	errCheck(err)
	defer destFile2.Close()

	res,e := io.Copy(destFile2,sourceFile2)
	errCheck(e)

	fmt.Println("Response: ",res)
	fmt.Println("Coping Complete!")
	
}