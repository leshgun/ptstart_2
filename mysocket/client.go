package mysocket

import (
	"net"
	"strconv"
)

const responseBuf = 1024

func ClientCreate(host string, port int) (net.Conn, error) {
	conn, _ := net.Dial("tcp", host+":"+strconv.Itoa(port))
	return conn, nil
}

func ClientSend(conn net.Conn, data string) (string, error) {
	conn.Write([]byte(data))
	buf := make([]byte, responseBuf)
	mLen, err := conn.Read(buf)
	if err != nil {
		return "", err
	}
	return string(buf[:mLen]), nil
}

func ClientClose(conn net.Conn) {
	conn.Close()
}
