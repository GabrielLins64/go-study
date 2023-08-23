<h1>Repositório de Estudo da Linguagem Go</h1>

Este é um repositório dedicado ao estudo da linguagem de programação Go. Aqui você encontrará recursos, exemplos e informações para ajudá-lo a aprender e dominar a linguagem Go.

---

- [Sobre a Linguagem Go](#sobre-a-linguagem-go)
- [Instalação (Linux)](#instalação-linux)
- [Módulos e Pacotes do Go](#módulos-e-pacotes-do-go)
  - [Pacotes externos](#pacotes-externos)
- [Go Workspaces](#go-workspaces)
- [Comandos Básicos do Go](#comandos-básicos-do-go)
- [Referências](#referências)

---

## Sobre a Linguagem Go

Go, também chamado de Golang, é uma linguagem de programação *open source* desenvolvida pelo Google. Desenvolvedores de software utilizam o Go em uma variedade de sistemas operacionais e estruturas para criar aplicações web, serviços em nuvem, redes e outros tipos de software.

O Go é estaticamente tipado, explícito e modelado com base na linguagem de programação C. Devido ao rápido tempo de inicialização, baixa sobrecarga em tempo de execução e capacidade de rodar sem uma máquina virtual (VM), a linguagem Go se tornou muito popular para escrever microsserviços e outros usos. Além disso, o Go é utilizado para programação concorrente — uma estratégia para executar múltiplas tarefas ao mesmo tempo, fora de ordem ou em ordem parcial.

A linguagem Go foi inspirada pela produtividade e relativa simplicidade do Python. Ela utiliza goroutines, ou processos leves, e uma coleção de pacotes para gerenciamento eficiente de dependências. Foi projetada para resolver diversos problemas, incluindo tempo lento de compilação, dependências descontroladas, duplicação de esforço, dificuldade na criação de ferramentas automáticas e desenvolvimento entre diferentes linguagens.

---

## Instalação (Linux)

Para instalar o Go no Linux, siga estas etapas:

1. Faça o download da versão mais recente do Go em: [Página de Downloads do Go](https://golang.org/dl/)

2. Remova qualque instalação prévia do Go:

```shell
sudo rm -rf /usr/local/go 
```

3. Extraia o arquivo compactado usando o seguinte comando (substitua `versao` pelo número da versão baixada):

```shell
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
```

4. Edite o arquivo de execução ao fazer login:

```shell
vim $HOME/.profile
```

5. E ao final do arquivo adicione a seguinte linha para adicionar o binário do go ao PATH:

```shell
export PATH=$PATH:/usr/local/go/bin
```

6. Verifique que foi instalado corretamente com:

```shell
go version
```

---

## Módulos e Pacotes do Go

Go possui um sistema de gerenciamento de dependências integrado chamado de módulos. Módulos permitem que você organize e gerencie as dependências do seu projeto de forma eficiente.

### Pacotes externos

Para importar e utilizar pacotes de terceiros, primeiro procure pelo pacote em [pkg.go.dev](https://pkg.go.dev), importe o pacote no código com:

```go
import "external/package/name"
```

*Substitua "external/package/name" pelo nome do pacote.

E, finalmente, atualize as dependências do seu módulo com:

```shell
go mod tidy
```

---

## Go Workspaces

Com espaços de trabalho de vários módulos &ndash; *multi-module workspaces* &ndash; você pode informar ao comando Go que está escrevendo código em vários módulos ao mesmo tempo e facilmente construir e executar código nesses módulos.

Você pode inicializer um workspace com:

```shell
go work init
```

Para usar um módulo específico dentro do workspace, como o *hello*, por exemplo, faça:

```shell
go work use hello
```

Então você consegue executar o arquivo `hello.go` no diretório do workspace com:

```shell
go run ./hello
```

---

## Comandos Básicos do Go

Aqui estão alguns comandos básicos que você usará frequentemente ao trabalhar com Go:

- Crie um novo módulo:

```shell
go mod init nome-do-modulo
```

- Compilar um projeto/script:

```shell
go build
```

- Executar um projeto/script:

```shell
go run <filepath/projectpath>
```

- Rodar testes:

```shell
go test
```

---

## Referências

[1] https://www.techtarget.com/

[2] https://go.dev/doc/

[3] https://www.codigofluente.com.br/
