package consulta

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	cl "github.com/fatih/color"
)

type Nome struct {
	Status  string `json:"status"`
	Results []struct {
		Nome string `json:"NOME"`
		Cpf  string `json:"CPF"`
		Sexo string `json:"SEXO"`
		Nasc string `json:"NASC"`
		Mae  string `json:"MAE"`
		Pai  string `json:"PAI"`
	} `json:"results"`
}

func Cnome() {
	var nome string
	var red = cl.New(cl.FgRed, cl.Bold)
	var green = cl.New(cl.FgHiGreen, cl.Bold)
	red.Print("Digite o nome: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		nome = scanner.Text()
	}
	url := fmt.Sprintf("https://apisdedicado.nexos.dev/SerasaNome/nome?token=2ae274ad75c45b657547631a82358dbc&nome=%s", nome)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var nomeData Nome
	err = json.NewDecoder(resp.Body).Decode(&nomeData)
	if err != nil {
		panic(fmt.Sprintf("Erro ao decodificar o json: %v", err))
	}
	if nomeData.Status == "success" {
		for _, result := range nomeData.Results {
			green.Printf("\nNome: %s\n", result.Nome)
			green.Printf("Sexo: %s\n", result.Sexo)
			green.Printf("Nascimento: %s\n", result.Nasc)
			green.Printf("Mãe: %s\n", result.Mae)
			green.Printf("Pai: %s\n", result.Pai)
		}
	} else {
		red.Println("Nome não encontrado")
	}
}
