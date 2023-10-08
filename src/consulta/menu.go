package consulta

import (
	"fmt"
	"os"
	"os/exec"

	cl "github.com/fatih/color"
	
)
func verify() {
	resp := "n"
	for {
		red.Print("\nDeseja continuar(s/n): ")
		fmt.Scanln(&resp)
		switch resp {
		case "s":
			MenuConsulta()
		case "n":
			return
		default:
			fmt.Printf("%s não é uma opção valida\n", resp)
			continue

		}
	}
}
func clearTerminal()  {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

var red = cl.New(cl.FgRed, cl.Bold)
func MenuConsulta(){
var op int
clearTerminal()
var menu = `MENU:
1- Consultar CPF
2- Consultar Nome
99 - Voltar
`



red.Println(menu)
red.Print("> ")
fmt.Scanln(&op)
switch op{
case 1:
	clearTerminal()
	Ccpf()
	verify()
case 2:
	clearTerminal()
	Cnome()
	verify()
case 99:
	return
default:
	MenuConsulta()

}
}















