package main

import (
	"fmt"
	"hallgo/src/consulta"
	"hallgo/src/dirscan"
	"hallgo/src/paramscan"
	"hallgo/src/portscan"
	"hallgo/src/webcrawler"
	"os"
	"os/exec"

	cl "github.com/fatih/color"
)

func clearTerminal()  {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func verify() {
	resp := "n"
	for {
		red.Print("\nDeseja continuar(s/n): ")
		fmt.Scanln(&resp)
		switch resp {
		case "s":
			Main()
		case "n":
			green.Println(logo)
			green.Println(diz)
			green.Println(ate)
			return
		default:
			fmt.Printf("%s não é uma opção valida\n", resp)
			continue

		}
	}
}

var ate = ` $$$$$$\    $$\                                             $$\           
$$  __$$\   $$ |                                            \__|          
$$ /  $$ |$$$$$$\    $$$$$$\        $$$$$$\$$$$\   $$$$$$\  $$\  $$$$$$$\ 
$$$$$$$$ |\_$$  _|  $$  __$$\       $$  _$$  _$$\  \____$$\ $$ |$$  _____|
$$  __$$ |  $$ |    $$$$$$$$ |      $$ / $$ / $$ | $$$$$$$ |$$ |\$$$$$$\  
$$ |  $$ |  $$ |$$\ $$   ____|      $$ | $$ | $$ |$$  __$$ |$$ | \____$$\ 
$$ |  $$ |  \$$$$  |\$$$$$$$\       $$ | $$ | $$ |\$$$$$$$ |$$ |$$$$$$$  |
\__|  \__|   \____/  \_______|      \__| \__| \__| \_______|\__|\_______/ 
                                                                          
                                                                          
                                                                          

`

var diz = ` ____  ____  ____ 
(  _ \(_  _)(_   )
 )(_) )_)(_  / /_ 
(____/(____)(____)

`

var menu = `MENU:
1- PortScan
2- DirScan
3- ParamScan
4- WebCrawler
5- Consultar Dados
99 - Sair
`
var logo = `██╗  ██╗ █████╗ ██╗     ██╗     
██║  ██║██╔══██╗██║     ██║     
███████║███████║██║     ██║     
██╔══██║██╔══██║██║     ██║     
██║  ██║██║  ██║███████╗███████╗
╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝╚══════╝
                                

`
var red = cl.New(cl.FgRed, cl.Bold)
var green = cl.New(cl.FgHiGreen, cl.Bold)

func Main() {
	clearTerminal()
	var op int
	red.Println(logo)
	red.Println(menu)
	red.Print("> ")
	fmt.Scanln(&op)
	
	switch op {
	case 1:
		clearTerminal()
		portscan.Port()
		verify()
	case 2:
		clearTerminal()
		dirscan.DirScan()
		verify()
	case 3:
		clearTerminal()
		paramscan.ParamScan()
		verify()
	case 4:
		clearTerminal()
		webcrawler.Crawler()
		verify()
	case 5:
		clearTerminal()
		consulta.MenuConsulta()
		Main()
		
	case 99:
		clearTerminal()
		green.Println(logo,diz,ate)
		return
		
		
	default:
		Main()
	}
}
func main()  {
	Main()
}