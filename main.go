package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("Listening on port :6379")
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	//infinite loop to receive commands from clients and respond to them
	for {
		resp := NewResp(conn)
		value, err := resp.Read()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(value)
		if value.typ != "array" {
			fmt.Println("Invalid request ,expected array")
			continue
		}
		if len(value.array) == 0 {
			fmt.Println("Invalid request,expected array length > 0")
			continue
		}
		command := strings.ToUpper(value.array[0].bulk)
		args := value.array[1:]

		writer := NewWriter(conn)
		handler, ok := Handlers[command]
		if !ok {
			fmt.Println("Invalid command: ", command)
			writer.Write(Value{typ: "error", str: "unknown command"})
			continue
		}
		result := handler(args)
		writer.Write(result)
		// writer.Write(Value{typ: "string", str: "OK"})
		// conn.Write([]byte("+OK\r\n"))
	}

}
