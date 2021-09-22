package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	arg_num := len(os.Args)
	fmt.Printf("the num of input is %d\n", arg_num)
	if arg_num < 2 {
		fmt.Printf("Usage %s 0.0.0.0:1223 0.0.0.0:1224\n", os.Args[0])
		return
	}

	for i := 1; i < arg_num; i++ {
		go listenOnUdpURL(os.Args[i])
	}

	for {
		time.Sleep(time.Duration(2) * time.Second)
	}
}

func listenOnUdpURL(url string) {
	fmt.Printf("Listening udp at[%s]\n", url)
	addr, err := net.ResolveUDPAddr("udp", url)
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		return
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer conn.Close()
	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	data := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Println("failed to read UDP msg because of", err.Error())
		return
	}
	fmt.Printf("Receved[%d] from client\n", n)
	conn.WriteToUDP(data[:n], remoteAddr)
}
