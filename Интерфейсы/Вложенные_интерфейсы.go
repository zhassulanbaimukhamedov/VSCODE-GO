package main

import "fmt"

type Reader interface {
	read()
}

type Writer interface {
	write(string)
}

type ReaderWriter interface {
	Reader
	Writer
}

func writeToStream(writer Writer, text string) {
	writer.write(text)
}

func readerFromStream(reader Reader) {
	reader.read()
}

type File struct {
	text string
}

func (f *File) read() {
	fmt.Println(f.text)
}

func (f *File) write(message string) {
	f.text = message
	fmt.Println("Запись в файл строки", message)
}

func main() {

	fmt.Println("Hello world")

	myFile := &File{}
	writeToStream(myFile, "hello world")
	readerFromStream(myFile)
	writeToStream(myFile, "lolly bomb")
	readerFromStream(myFile)
}
