package main

import (
	"GoCalculator/pkg/calc"
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("---------------------------------------------------")
	fmt.Println("---------------Go Calculator Server ---------------")
	fmt.Println("---------------------------------------------------")
	fmt.Println("---------------------------------------------------")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go calc.HandleCalc(conn)
	}
}
