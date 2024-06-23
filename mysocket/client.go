package mysocket

import (
	"fmt"
	"net"
	"strconv"
)

func ClientCreate(host string, port int) (net.Conn, error) {
	conn, _ := net.Dial("tcp", host+":"+strconv.Itoa(port))
	return conn, nil
}

func ClientSend(conn net.Conn, data string) string {
	conn.Write([]byte(data))
	buf := make([]byte, 1024)
	mLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	return string(buf[:mLen])
}

func ClientClose(conn net.Conn) {
	conn.Close()
}
