package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const file = "testfile"

func main() {
	f, err := os.Create(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	str := strings.Repeat("testfile\n", 100)
	_, err = f.Write([]byte(str))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	f.Seek(0, 0)

	buf := make([]byte, 4096, 4096)
	cnt := 0
	for {
		n, err := f.Read(buf[0:])
		if n != 0 {
			fmt.Print(string(buf[:n]))
			cnt += n
		}

		if err != io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(string(buf[:n]))
	}

	fmt.Println("bytes:", cnt)
}
