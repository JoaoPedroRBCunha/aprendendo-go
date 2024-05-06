package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	monitoramentos = 3
	delay          = 5
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
			imprimeLog()
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
	fmt.Println()

	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(time.Second * delay)
		fmt.Println()
	}
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site: ", site, "está com problemas. Status code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt") // os.Open - é utilizado para abrir um arquivo
	defer arquivo.Close()                // É uma boa prática fechar os arquivos no final de seu uso.
	// arquivo, err := os.ReadFile("sites.txt") - os.ReadFile - serve para ler todo arquivo
	// fmt.Println(strings(arquivo)) - strings(arquivo) - converte os bytes do arquivo para strings legíveis

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	/*
		bufio.NewReader() - te retorna um leitor onde vão ter umas funções que vão
		te ajudar a passear pelo arquivo lendo linha a linha
	*/
	for {
		linha, err := leitor.ReadString('\n')
		/*
			leitor.ReadString - faz com que leia uma linha inteira e nos retorne uma string dela
			e ele tem que receber um valor em byte (utilizando aspas simples) que vai delimitar até onde ele vai ler naquela linha
			utilizando um fmt.Println(linha) ele vai retornar apenas a primeira linha do arquivo
		*/
		linha = strings.TrimSpace(linha) // strings.TrimSpace - corta os espaços, os \n e os tabs do final da sua string
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	return sites

}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	/*
		os.OpenFile - abre o arquivo
		os.O_RDWR - onde consigo escrever e ler os dados dentro do arquivo
		os.O_CREATE - cria o arquivo caso ele não tenha sido criado
		os.O_APPEND - possibilita que você adicione vários tipos de dados no arquivo
	*/
	defer arquivo.Close() // É uma boa prática fechar os arquivos no final de seu uso.
	if err != nil {
		fmt.Println(err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	/*
		arquivo.WriteString - serve para escrever no arquivo
		time.Now().Format - faz com que você formate o tempo para o jeito que você quiser
		time.Now().Format - verificar site para saber como formatar (https://go.dev/src/time/format.go)
	*/
}

func imprimeLog() {
	fmt.Println("Exibindo logs...")
	arquivo, err := os.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
