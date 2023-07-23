package main

import (
	"fmt"
	"net"
	"sync"
)

func scanPort(host string, port int, wg *sync.WaitGroup) {
	defer wg.Done()

	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("Port %d on %s is closed\n", port, host)
		return
	}
	defer conn.Close()

	fmt.Printf("Port %d on %s is open\n", port, host)
}

func main() {
	host := "localhost"
	var wg sync.WaitGroup

	for port := 1; port <= 65535; port++ {
		wg.Add(1)
		go scanPort(host, port, &wg)
	}

	wg.Wait()
}
