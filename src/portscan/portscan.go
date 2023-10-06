package portscan

import (
	"fmt"
	"net"
	"time"
	 cl "github.com/fatih/color"
)

func scanPort(target string, port int) bool {
	andress := fmt.Sprintf("%s:%v", target, port)
	conn, err := net.DialTimeout("tcp", andress, 2*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
func Scan() {
	red := cl.New(cl.FgRed, cl.Bold)
	var target string
	red.Print("Digite o host: ")
	fmt.Scanln(&target)
	ports := []int{21, 22, 2553, 80, 110, 143, 443, 3306, 5432, 8080, 3690}
	for _, port := range ports {
		if scanPort(target, port) {
			fmt.Printf("Porta aberta! %v\n", port)
		} else {
			fmt.Printf("Porta fechada! %v\n", port)
		}
	}
}
func main() {
	Scan()
}
