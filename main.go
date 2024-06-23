package main

import (
	"fmt"
	"int18/floatmath"
	"int18/mysocket"
	"strconv"
	"strings"
	"time"
)

func main() {
	sumStringFloat()
	fmt.Println()
	testWebsocketServer()
}

func sumStringFloat() {
	var timeBegin time.Time

	println(
		strings.Repeat("=", 25),
		"Test sum of string floats",
		strings.Repeat("=", 25),
	)

	actSum := "0.333333333333333333333333333333"
	nums, err := floatmath.ReadFile("numbers.txt")
	if err != nil {
		return
	}
	fmt.Println("Numbers:", nums)
	fmt.Println("The actually sum is:", actSum)

	println(strings.Repeat("-", 20), "Simple sum of floats", strings.Repeat("-", 20))
	timeBegin = time.Now()
	sum, _ := floatmath.SumStringArrayFloat(nums, false)
	timeSum := time.Since(timeBegin)
	fmt.Printf("The sum is:%.30f\n", sum)
	presTest := strconv.FormatFloat(sum, 'f', 30, 64) == actSum
	fmt.Println("Precision test:", presTest)
	fmt.Printf("It took: %d nanoseconds\n", timeSum.Nanoseconds())

	println(strings.Repeat("-", 20), "Sum of big floats", strings.Repeat("-", 20))
	timeBegin = time.Now()
	sumBig, _ := floatmath.SumStringArrayBigFloat(nums, 30)
	timeBigSum := time.Since(timeBegin)
	fmt.Println("The big sum is:", sumBig)
	fmt.Println("Precision test:", sumBig == actSum)
	fmt.Printf("It took: %d nanoseconds\n", timeBigSum.Nanoseconds())
}

func testWebsocketServer() {
	println(
		strings.Repeat("=", 23),
		"Test server and connections to him",
		strings.Repeat("=", 23),
	)

	host := "localhost"
	port := 11111
	clientNum := 10
	limit := 5
	interval := 1

	server, _ := mysocket.ServerCreate(host, port)
	go mysocket.ServerStart(server, limit, interval)
	fmt.Printf("Server started at: %s:%d\n", host, port)

	timeBegin := time.Now()
	for i := 0; i < clientNum; i++ {
		request := "{"
		request += fmt.Sprintf(`"X1": "0.%s",`, strconv.Itoa(i))
		request += fmt.Sprintf(`"X2": "0.%s",`, strconv.Itoa(i*2))
		request += fmt.Sprintf(`"X3": "0.%s",`, strconv.Itoa(i*3))
		request += fmt.Sprintf(`"Y1": "0.%s",`, strconv.Itoa(i))
		request += fmt.Sprintf(`"Y2": "0.%s",`, strconv.Itoa(i*2))
		request += fmt.Sprintf(`"Y3": "0.%s",`, strconv.Itoa(i*i*3))
		request += fmt.Sprintf(`"E": %d`, i+2)
		request += "}"
		fmt.Printf("Client #%d [%s]\n", i+1, time.Since(timeBegin).Round(time.Millisecond))
		fmt.Printf("-- Request: %s\n", request)
		client, _ := mysocket.ClientCreate(host, port)
		response, _ := mysocket.ClientSend(client, request)
		fmt.Printf("-- Response:%s\n", response)
		time.Sleep(time.Millisecond * time.Duration(500/(i+1)))
	}

	mysocket.ServerStop(server)
}
