package scanner

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

type Scanner struct {
	Timeout time.Duration
	Workers int
}

type ScanResult struct {
	Host      string
	OpenPorts []int
}

// NewScanner creates a new scanner with default timeout and workers
func NewScanner() *Scanner {
	return &Scanner{
		Timeout: 500 * time.Millisecond,
		Workers: 100, // Default using 100 workers
	}
}

// SetTimeout sets the connection timeout for the scanner
func (s *Scanner) SetTimeout(timeout time.Duration) {
	s.Timeout = timeout
}

// SetWorkers sets the number of concurrent workers
func (s *Scanner) SetWorkers(workers int) {
	if workers <= 0 {
		workers = 1
	}
	s.Workers = workers
}

func (s *Scanner) ScanPorts(host string, startPort, endPort int) (*ScanResult, error) {
	fmt.Printf("Scanning ports %d-%d on %s...\n", startPort, endPort, host)

	var wg sync.WaitGroup
	openPorts := make(chan int, endPort-startPort+1)
	ports := make(chan int, endPort-startPort+1)

	// Start workers
	for i := 0; i < s.Workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for port := range ports {
				if s.isPortOpen(host, port) {
					openPorts <- port
				}
			}
		}()
	}

	// Send ports to channel
	go func() {
		for port := startPort; port <= endPort; port++ {
			ports <- port
		}
		close(ports)
	}()

	// Wait for all workers to complete
	wg.Wait()
	close(openPorts)

	var resultPorts []int
	for p := range openPorts {
		resultPorts = append(resultPorts, p)
	}

	return &ScanResult{
		Host:      host,
		OpenPorts: resultPorts,
	}, nil
}

func (s *Scanner) isPortOpen(host string, port int) bool {
	address := net.JoinHostPort(host, strconv.Itoa(port))
	conn, err := net.DialTimeout("tcp", address, s.Timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func (s *Scanner) RunInteractiveMode() {
	var host string
	fmt.Print("Enter the host to scan (example: 192.168.1.1): ")
	fmt.Scanln(&host)

	var startPort, endPort int
	fmt.Print("Enter the start port (example: 1): ")
	fmt.Scanln(&startPort)
	fmt.Print("Enter the end port (example: 1024): ")
	fmt.Scanln(&endPort)

	result, err := s.ScanPorts(host, startPort, endPort)
	if err != nil {
		fmt.Printf("Error scanning ports: %v\n", err)
		return
	}

	s.PrintResults(result)
}

func (s *Scanner) PrintResults(result *ScanResult) {
	fmt.Println("------------- Scan complete -------------")

	if len(result.OpenPorts) == 0 {
		fmt.Println("No open ports found")
	} else {
		for _, port := range result.OpenPorts {
			fmt.Printf("Port %d is open\n", port)
		}
	}

	fmt.Println("----------------------------------------")
}
