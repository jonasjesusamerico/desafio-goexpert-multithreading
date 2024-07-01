package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Struct para armazenar os dados do endereço
type Address struct {
	Cep          string `json:"cep"`
	State        string `json:"state,omitempty"`
	Uf           string `json:"uf,omitempty"`
	City         string `json:"city,omitempty"`
	Localidade   string `json:"localidade,omitempty"`
	Neighborhood string `json:"neighborhood,omitempty"`
	Bairro       string `json:"bairro,omitempty"`
	Street       string `json:"street,omitempty"`
	Logradouro   string `json:"logradouro,omitempty"`
	Service      string `json:"service,omitempty"`
	Complemento  string `json:"complemento,omitempty"`
	Unidade      string `json:"unidade,omitempty"`
	Ibge         string `json:"ibge,omitempty"`
	Gia          string `json:"gia,omitempty"`
	Ddd          string `json:"ddd,omitempty"`
	Siafi        string `json:"siafi,omitempty"`
}

func main() {
	cep := "01153000"
	ch := make(chan *http.Response)
	timeout := time.After(1 * time.Second)

	go requestBrasilAPI(cep, ch)
	go requestViaCEP(cep, ch)

	select {
	case resp := <-ch:
		if resp == nil {
			fmt.Println("Erro ao fazer a requisição")
			return
		}
		defer resp.Body.Close()

		var address Address
		if err := json.NewDecoder(resp.Body).Decode(&address); err != nil {
			fmt.Println("Erro ao decodificar a resposta:", err)
			return
		}

		// Ajustar campos caso necessário
		if address.State == "" {
			address.State = address.Uf
		}
		if address.City == "" {
			address.City = address.Localidade
		}
		if address.Neighborhood == "" {
			address.Neighborhood = address.Bairro
		}
		if address.Street == "" {
			address.Street = address.Logradouro
		}

		// Verificar de qual API veio a resposta
		var api string
		if strings.Contains(resp.Request.URL.Host, "brasilapi.com.br") {
			api = "BrasilAPI"
		} else {
			api = "ViaCEP"
		}

		// Exibir os dados do endereço e a API utilizada
		fmt.Println("---------------------------------------")
		fmt.Printf("Resposta da %s\n", api)
		fmt.Printf("  CEP: %s\n", address.Cep)
		fmt.Printf("  Logradouro: %s\n", address.Street)
		fmt.Printf("  Bairro: %s\n", address.Neighborhood)
		fmt.Printf("  Localidade: %s\n", address.City)
		fmt.Printf("  UF: %s\n", address.State)
		fmt.Println("---------------------------------------")
	case <-timeout:
		fmt.Println("Erro de timeout: tempo limite de 1 segundo excedido")
	}
}

// Função para fazer a requisição à BrasilAPI
func requestBrasilAPI(cep string, ch chan<- *http.Response) {
	url := "https://brasilapi.com.br/api/cep/v1/" + cep
	resp, err := http.Get(url)
	if err != nil {
		ch <- nil
		return
	}
	ch <- resp
}

// Função para fazer a requisição à ViaCEP
func requestViaCEP(cep string, ch chan<- *http.Response) {
	url := "http://viacep.com.br/ws/" + cep + "/json/"
	resp, err := http.Get(url)
	if err != nil {
		ch <- nil
		return
	}
	ch <- resp
}
