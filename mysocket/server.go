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
	X1 string
	X2 string
	X3 string
	Y1 string
	Y2 string
	Y3 string
	E  int
}

type Response struct {
	Code  int
	Error string
	X     string
	Y     string
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
	prec := uint(req.E * 4)
	bx1, _ := big.NewFloat(0.0).SetPrec(prec).SetString(req.X1)
	bx2, _ := big.NewFloat(0.0).SetPrec(prec).SetString(req.X2)
	bx3, _ := big.NewFloat(0.0).SetPrec(prec).SetString(req.X3)
	by1, _ := big.NewFloat(0.0).SetPrec(prec).SetString(req.Y1)
	by2, _ := big.NewFloat(0.0).SetPrec(prec).SetString(req.Y2)
	by3, _ := big.NewFloat(0.0).SetPrec(prec).SetString(req.Y3)
	zero := big.NewFloat(0.0)
	if bx2.Cmp(zero) == 0 || by2.Cmp(zero) == 0 {
		return Response{Code: serverCodeErr, Error: "Integer devide by zero"}
	}
	bx1.Quo(bx1, bx2)
	bx1.Mul(bx1, bx3)
	by1.Quo(by1, by2)
	by1.Mul(by1, by3)
	var e string
	if e = "F"; bx1.Cmp(by1) == 0 {
		e = "T"
	}
	return Response{
		Code: serverCodeSuc,
		X:    bx1.Text('f', req.E),
		Y:    by1.Text('f', req.E),
		E:    e,
	}
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
