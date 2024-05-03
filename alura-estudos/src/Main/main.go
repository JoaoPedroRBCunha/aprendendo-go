package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Douglas"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)

	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := [4]string{
		"https://httpbin.org/status/404",
		"https://www.alura.com.br",
		"https://httpbin.org/status/200",
		"https://www.caelum.com.br",
	}
	fmt.Println(sites)
	for _, site := range sites {
		resp, _ := http.Get(site)
		if resp.StatusCode == 200 {
			fmt.Println("Site: ", site, "foi carregado com sucesso!")
		} else {
			fmt.Println("Site: ", site, "está com problemas. Status code:", resp.StatusCode)
		}
	}

}
