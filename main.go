package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: netfwd [listen] [target]")
		return
	}

	list, err := net.Listen("tcp", os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		con, err := list.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go func(c net.Conn) {
			other, err := net.Dial("tcp", os.Args[2])
			if err != nil {
				fmt.Println(err)
				return
			}

			go io.Copy(other, c)
			go io.Copy(c, other)
		}(con)
	}
}
