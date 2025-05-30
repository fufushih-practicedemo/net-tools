package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

func main() {
	var host string
	fmt.Print("Enter the host to scan (example: 192.168.1.1): ")
	fmt.Scanln(&host)

	var startPort, endPort int
	fmt.Print("Enter the start port (example: 1): ")
	fmt.Scanln(&startPort)
	fmt.Print("Enter the end port (example: 1024): ")
	fmt.Scanln(&endPort)

	fmt.Println("Scanning ports...")

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	openPorts := make(chan int, endPort-startPort+1) // Buffered channel to store open ports

	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done() // Mark the goroutine as done when it finishes

			address := net.JoinHostPort(host, strconv.Itoa(p))
			conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
			if err != nil {
				return
			}
			conn.Close()
			openPorts <- p // Send the open port to the channel
		}(port)
	}

	wg.Wait()
	close(openPorts) // Close the channel to signal that no more ports will be sent

	fmt.Println("------------- Scan complete -------------")
	foundOpenPorts := false
	for p := range openPorts {
		fmt.Printf("Port %d is open\n", p)
		foundOpenPorts = true
	}

	if !foundOpenPorts {
		fmt.Println("No open ports found")
	}

	fmt.Println("----------------------------------------")
}
