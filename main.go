package main

import (
	"fmt"
	"hallgo/src/dirscan"
	"hallgo/src/portscan"
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
			main()
		case "n":
			red.Println(logo)
			red.Println(diz)
			red.Println(ate)
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


func main() {
	clearTerminal()
	var op int
	red.Println(logo)
	red.Println(menu)
	red.Print("> ")
	fmt.Scanln(&op)
	
	switch op {
	case 1:
		clearTerminal()
		portscan.Scan()
		verify()
	case 2:
		clearTerminal()
		dirscan.DirScan()
		verify()
	
	default:
		main()
	}
}
