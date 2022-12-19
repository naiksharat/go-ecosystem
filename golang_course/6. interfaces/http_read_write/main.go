package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	const url = "http://google.com"
	resp, _ := http.Get(url)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	//Reader:
	//resp.Body.Read is a reader function
	// it has some info that it's keeping to itself
	// we can pass a buffer ([]byte) to the readeer fucntion and ask it to fill the buffer with info
	bs := make([]byte, 9)
	resp.Body.Read(bs)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(2)
	// }
	fmt.Println(string(bs))

	// read from resp.body -- READER
	// write to os.stdout -- WRITER
	io.Copy(os.Stdout, resp.Body)

	// Custom writer/custom reader // implement write/read function
}
