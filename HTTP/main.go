package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:- ",err)
		os.Exit(1)
	}

	// fmt.Print(resp)
	// bs := make([]byte,999999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	io.Copy(logWriter{}, resp.Body)
}

func (logWriter) Write(b []byte)(int,error){
	fmt.Println(string(b))
	fmt.Printf("Just wrote %v number of bytes",len(b))
	return len(b),nil
}
