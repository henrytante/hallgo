package consulta

import (
	"fmt"
	"net/http"
	"encoding/json"
	cl "github.com/fatih/color"
	
)

type Cpf struct{
	Status string `json:"status"`
	Results struct{
		Nome string `json:"NOME"`
		Sexo string `json:"SEXO"`
		Nasc string `json:"NASC`
		Nmae string `json:"NOME_MAE"`
		Npai string `json:"NOME_PAI"`
		Rg string `json:"Rg"`
	} `json:"results"`
}

func Ccpf(){
	var green = cl.New(cl.FgHiGreen, cl.Bold)
	var red = cl.New(cl.FgRed, cl.Bold)
	var cpf string
	red.Print("Digite o cpf: ")
	fmt.Scanln(&cpf)
	url := fmt.Sprintf("https://apisdedicado.nexos.dev/SerasaCpf/cpf?token=2ae274ad75c45b657547631a82358dbc&cpf=%s", cpf)
	resp, err := http.Get(url)
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()
	var cpfData Cpf
	err = json.NewDecoder(resp.Body).Decode(&cpfData)
	if err != nil{
		panic(err)
	}
	if cpfData.Status == "success"{
		green.Printf("\nNome: %s\n", cpfData.Results.Nome)
		green.Printf("Sexo: %s\n", cpfData.Results.Sexo)
		green.Printf("Nascimento: %s\n", cpfData.Results.Nasc)
		green.Printf("Mãe: %s\n", cpfData.Results.Nmae)
		green.Printf("Pai: %s\n", cpfData.Results.Npai)
		green.Printf("RG: %s\n", cpfData.Results.Rg)
		
	}
	
}