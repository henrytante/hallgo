package portscan

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	cl "github.com/fatih/color"
	"sync"
	"time"
)

func Port() {
	var host string
	red := cl.New(cl.FgRed, cl.Bold)
	green := cl.New(cl.FgHiGreen, cl.Bold)
	yellow := cl.New(cl.FgHiYellow, cl.Bold)
	red.Print("Digite seu host: ")
	fmt.Scanln(&host)
	target := host

	// Defina as portas a serem verificadas
	portsStr := "80,443,22,3306,1433,5432,6379,25,110,143,993,995,53,161,162,5900"
	ports := parsePorts(portsStr)

	red.Printf("\nIniciando varredura em %s...\n\n", target)

	var wg sync.WaitGroup
	openPorts := make(chan int)

	for _, port := range ports {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", target, p)
			conn, err := net.DialTimeout("tcp", address, 2*time.Second)
			if err == nil {
				defer conn.Close()
				openPorts <- p
			}
		}(port)
	}

	go func() {
		wg.Wait()
		close(openPorts)
	}()

	var openPortList []int
	for port := range openPorts {
		openPortList = append(openPortList, port)
	}

	yellow.Println("Portas abertas:\n")
	for _, port := range openPortList {
		serviceName := getServiceName(port)
		green.Printf("Porta %d aberta - Serviço: %s\n", port, serviceName)
	}
}

func parsePorts(portStr string) []int {
	portList := strings.Split(portStr, ",")
	var ports []int

	for _, port := range portList {
		p, err := strconv.Atoi(port)
		if err == nil {
			ports = append(ports, p)
		}
	}

	return ports
}

func getServiceName(port int) string {
	serviceMap := map[int]string{
		22:    "SSH",
		25:    "SMTP",
		53:    "DNS",
		80:    "HTTP",
		110:   "POP3",
		143:   "IMAP",
		161:   "SNMP",
		162:   "SNMP-TRAP",
		443:   "HTTPS",
		5900:  "VNC",
		6379:  "Redis",
		993:   "IMAPS",
		995:   "POP3S",
		1433:  "MSSQL",
		3306:  "MySQL",
		5432:  "PostgreSQL",
		// Adicione mais serviços e portas aqui conforme necessário
	}

	if serviceName, ok := serviceMap[port]; ok {
		return serviceName
	}

	return "Desconhecido"
}
