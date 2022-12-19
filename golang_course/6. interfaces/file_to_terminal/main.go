package main

import (
	"fmt"
	"io"
	"os"
)

type NoCliArgError struct{}

func (ce NoCliArgError) Error() string {
	fmt.Println("Len < 2. No CLI Arg")
	return ""
}

type FileDoesntExist struct{}

func (ce FileDoesntExist) Error() string {
	fmt.Println("File doesn't exist")
	return ""
}

type customWriter struct{}

func (cw customWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs), "from this")
	return len(bs), nil
}

func main() {
	if len(os.Args) < 2 {
		_ = NoCliArgError{}.Error()
		os.Exit(1)
	}

	file := os.Args[1]

	f, err := os.Open(file)
	if err != nil {
		_ = FileDoesntExist{}.Error()
		os.Exit(2)
	}

	//buf := []byte{}
	//bytes_read, err := f.Read(buf)
	//fmt.Println(string(buf), bytes_read, err)
	io.Copy(customWriter{}, f)

	f.Seek(0, 0)
	io.Copy(customWriter{}, f)

	f.Seek(0, 0)
	io.Copy(os.Stdout, f)

}
