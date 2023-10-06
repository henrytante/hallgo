package dirscan

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	

	cl "github.com/fatih/color"
)

func DirScan()  {
	var file string
	var url string
	green := cl.New(cl.FgHiGreen, cl.Bold)
	yellow := cl.New(cl.FgHiYellow, cl.Bold)
	red := cl.New(cl.FgRed, cl.Bold)
	green.Print("Digite a sua url (http://exemple.com/): ")
	fmt.Scanln(&url)
	
	red.Print("Digite o caminho da sua wordlist: ")
	fmt.Scanln(&file)
	rd , err := os.Open(file)
	if err != nil{
		panic(err.Error())
	}	
	defer rd.Close()
	scanner := bufio.NewScanner(rd)
	for scanner.Scan(){
		dirs := scanner.Text()
		urladress := url + dirs
		test, err :=  http.Get(urladress)
		if err != nil{
			panic("Erro na conexão com o host")
		}
		defer test.Body.Close()
		if test.StatusCode == 200 || test.StatusCode == 301{
			green.Printf("Diretorio encontrado: %v 200\n", url + dirs)
		}else if test.StatusCode == 401{
			yellow.Printf("Diretorio encontrado: %v 401\n", url + dirs)
		}else{
			red.Printf("%v 404\n", url+dirs)
		}
	}
	if err := scanner.Err(); err != nil{
		errorM := fmt.Sprintf("Erro no %v", err)
		panic(errorM)
	}
}