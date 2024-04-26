package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade int    `json:"idade"`
}

func main() {
	c := cachorro{"Tobby", "Poodle", 5}
	fmt.Println(c)

	dadosCachorro, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(bytes.NewBuffer(dadosCachorro))

	cachorroEmJson := `{"nome":"Rex","raca":"Dalmanta","idade":5}`
	var dadosCachorro2 cachorro
	if erro := json.Unmarshal([]byte(cachorroEmJson), &dadosCachorro2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(dadosCachorro2)

}
