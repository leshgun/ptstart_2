package mysocket

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net"
	"strconv"
	"time"
)

type Request struct {
	X1 float64
	X2 float64
	X3 float64
	Y1 float64
	Y2 float64
	Y3 float64
	E  int
}

type Response struct {
	Code  int
	Error string
	X     float64
	Y     float64
	E     string
}

const (
	clientBuf               = 1024
	serverCodeSuc           = 200
	serverCodeErr           = 400
	serverCodeErrServerBusy = 402
)

func ServerCreate(host string, port int) (net.Listener, error) {
	// Bind socket for server
	listener, err := net.Listen("tcp", host+":"+strconv.Itoa(port))
	if err != nil {
		return nil, err
	}
	return listener, nil
}

func ServerStart(listener net.Listener, limit int, interval int) error {
	var prevConnTime time.Time
	var errMsg string
	// minimum time between connections (depends on limit and interval)
	var minTime time.Duration
	defer ServerStop(listener)
	// if 'limit' <= 0, then there is no limit for connections
	if limit > 0 {
		minTime = time.Duration(interval) * time.Second / time.Duration(limit)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		if minTime > 0 && !prevConnTime.IsZero() {
			if time.Since(prevConnTime) < minTime {
				errMsg = fmt.Sprintf(
					"Server is busy! (%s / %s)",
					time.Since(prevConnTime).String(),
					minTime.String(),
				)
				sendResponse(conn, serverError(serverCodeErrServerBusy, errMsg))
				conn.Close()
				continue
			}
		}
		prevConnTime = time.Now()
		go clientHandler(conn)
	}
}

func ServerStop(listener net.Listener) {
	listener.Close()
}

func clientHandler(conn net.Conn) {
	var err error
	var mesLen int
	var res Response
	defer conn.Close()

	buf := make([]byte, clientBuf)
	mesLen, err = conn.Read(buf)
	if err != nil {
		res = serverError(serverCodeErr, err.Error())
		sendResponse(conn, res)
		return
	}

	req := Request{}
	err = json.Unmarshal(buf[:mesLen], &req)
	if err != nil {
		res = serverError(serverCodeErr, err.Error())
		sendResponse(conn, res)
	}

	res = calcResult(req)
	sendResponse(conn, res)
}

func serverError(code int, err string) Response {
	return Response{
		Code:  code,
		Error: err,
	}
}

func calcResult(req Request) Response {
	if req.X2 == 0 || req.Y2 == 0 {
		return Response{Code: serverCodeErr, Error: "Integer devide by zero"}
	}
	bx := big.NewFloat(req.X1).SetPrec(uint(req.E))
	bx.Quo(bx, big.NewFloat(req.X2))
	bx.Mul(bx, big.NewFloat(req.X3))
	by := big.NewFloat(req.Y1).SetPrec(uint(req.E))
	by.Quo(by, big.NewFloat(req.Y2))
	by.Mul(by, big.NewFloat(req.Y3))
	var e string
	if e = "F"; bx.Cmp(by) == 0 {
		e = "T"
	}
	x, _ := bx.Float64()
	y, _ := by.Float64()
	return Response{Code: serverCodeSuc, X: x, Y: y, E: e}
}

func responseToBytes(res Response) ([]byte, error) {
	resMar, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	return resMar, nil
}

func sendResponse(conn net.Conn, res Response) {
	data, _ := responseToBytes(res)
	conn.Write(data)
}
