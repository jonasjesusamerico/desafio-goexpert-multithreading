# desafio-goexpert-multithreading


Esta aplicação em Go faz requisições simultâneas a duas APIs de CEP (BrasilAPI e ViaCEP) e exibe o resultado mais rápido no terminal. Caso nenhuma resposta seja recebida dentro de 1 segundo, um erro de timeout é exibido.

## Pré-requisitos

- Go 1.16 ou superior instalado. Você pode baixar o Go [aqui](https://golang.org/dl/).

## Como Executar

1. Clone o repositório ou copie os arquivos para um diretório local:

    ```sh
    git clone https://github.com/jonasjesusamerico/desafio-goexpert-multithreading.git
    ```

2. Execute o comando abaixo para iniciar a aplicação:

    ```sh
    go run main.go
    ```

3. A aplicação irá fazer requisições às APIs e exibirá o resultado no terminal.

## O Que a Aplicação Faz

- Faz requisições simultâneas para:
  - BrasilAPI: `https://brasilapi.com.br/api/cep/v1/01153000`
  - ViaCEP: `http://viacep.com.br/ws/01153000/json/`
- Recebe a resposta mais rápida entre as duas APIs.
- Exibe os dados do endereço retornado pela API mais rápida no terminal.
- Limita o tempo de resposta em 1 segundo. Caso contrário, um erro de timeout é exibido.

## Exemplo de Saída

```sh
---------------------------------------
Resposta da ViaCEP
  CEP: 01153-000
  Logradouro: Rua Vitorino Carmilo
  Bairro: Barra Funda
  Localidade: São Paulo
  UF: SP
---------------------------------------
