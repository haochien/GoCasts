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
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)
	/*
		break down io.Copy(lw, resp.Body):

		1. similar simple idea:
		bs := make([]byte, 99999)
		resp.Body.Read(bs)   //hear Read assign resp html into empty bs
		fmt.Println(string(bs))

		2. using stander Writer os.Stdout
		io.Copy(os.Stdout, resp.Body)
		// for func Copy(dst Writer, src Reader): it takes src (Reader interface)
		   and exports via dst (Writer interface)
	*/

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
