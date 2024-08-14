package calc

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func HandleCalc(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024*4)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Incorrect input")
			break
		}
		source := string(buf[:n])
		source = strings.Trim(source, " \n\r")
		if source == "exit" {
			conn.Write([]byte("Bye bye!"))
			os.Exit(1)
		}

		m := strings.Split(source, " ")
		if len(m) < 3 {
			res := "Invalid input format. Expected format: num1 action num2\n"
			conn.Write([]byte(res))
			continue
		}

		num1Str := m[0]
		action := m[1]
		num2Str := m[2]

		num1, err1 := strconv.ParseFloat(num1Str, 32)
		num2, err2 := strconv.ParseFloat(num2Str, 32)

		if err1 != nil || err2 != nil {
			res := "Invalid numbers provided. Please provide valid integers.\n"
			conn.Write([]byte(res))
			continue
		}

		var res string
		switch action {
		case "+":
			res = fmt.Sprintf("%.0f + %.0f = %.3f\n", num1, num2, num1+num2)
		case "-":
			res = fmt.Sprintf("%.0f - %.0f = %.3f\n", num1, num2, num1-num2)
		case "*":
			res = fmt.Sprintf("%.0f * %.0f = %.3f\n", num1, num2, num1*num2)
		case "/":
			if num2 == 0 {
				res = "Error: Division by zero\n"
			} else {
				res = fmt.Sprintf("%.0f / %.0f = %.3f\n", num1, num2, num1/num2)
			}
		default:
			res = fmt.Sprintf("Unknown action: %s\n", action)
		}
		conn.Write([]byte(res))
	}
}
