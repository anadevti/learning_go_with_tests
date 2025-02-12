### Escrever um teste é como escrever uma função, com algumas regras:

- Precisa estar em um arquivo com um nome parecido com `xxx_test.go`

- A função de teste precisa começar com a palavra `Test`

- A função de teste recebe um único argumento, que é `t *testing.T`

- o nosso `t` do tipo `*testing.T` é a nossa porta de entrada para a ferramenta de testes e assim você poderá utilizar o `t.Fail()` quando precisar relatar um erro.

- com o comando `go install golang.org/x/tools/cmd/godoc@latest` podemos rodar um servidor de doc da nossa aplicação, basta dar o comando após a instalação:
  `godoc -http :8000`