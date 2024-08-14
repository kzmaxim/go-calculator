package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("----------------------Go Calculator Client ---------------------------")
	fmt.Println("-----Введите число, действие и число, чтобы получить ответ------------")
	fmt.Println("----------------------------------------------------------------------")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter calculation (e.g., 2 + 3): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		if n, err := conn.Write([]byte(input)); err != nil || n == 0 {
			fmt.Println(err)
			break
		}
		buf := make([]byte, 1024*4)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		res := string(buf[:n])
		if res == "Bye bye!" {
			return
		}
		fmt.Println("\nОтвет:")
		fmt.Println(res)
	}
}
