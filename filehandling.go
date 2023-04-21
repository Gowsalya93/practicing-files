package main

import (
	"fmt"
	"ioutil/io"
	"log"
	"os"
)
func createFile(){
	fmt.Println("writing a file to golang")
	File,err:=os.Create("test.text")
	if err!=nil{
		log.Fatalf(failed creating file:",err)
	}
	defer file.close()
	len,err:=file.WriteString("welcome to golang programs")
	if err!=nil{
		log.Fatalf(failed writing to file:",err)
	}
	fmt.Println("FileName:%s",file.name())
	fmt.Println("length:%dbytes",len)
func ReadFile(){
	fmt.Println("Reading a file in golang")
	fileName="test.txt"
	data,err:=ioutil.Readfile("test.txt")
	if err!=nil{
		log.Panicf("failed reading data from file:%s",err)
	}
fmt.Println("FileName:%s",fileName)
fmt.Println("size:%d bytes",len(data))
fmt.Printf("data:%s",data)
	}
	func main(){
	createFile()
	ReadFile()
	}
}