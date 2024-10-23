# Planos de Estudos para Go

Este guia de estudos foi criado para ajudar a aprender Go de forma prática e eficiente. O repositório contém exemplos de código, exercícios e projetos práticos para ajudar a entender os conceitos fundamentais da linguagem Go.

## Semana 1: Fundamentos da Linguagem

1. Instalação e Ambiente

    - Instale o Go e configure seu ambiente.
    - Familiarize-se com o GOPATH e com a estrutura de diretórios.
    - Utilize um editor com suporte a Go (VS Code com a extensão Go é uma boa escolha).

2. Primeiros Passos

    - "Hello, World!" em Go.
    - Conheça as convenções da linguagem, como a organização dos arquivos .go.

3. Sintaxe Básica

    - Tipos de dados primitivos: int, float, string, bool.
    - Declaração de variáveis (usando var e :=).
    - Controle de fluxo: if, else, switch.
    - Laços de repetição: for (Go não possui while).

4. Funções

    - Declaração e chamada de funções.
    - Parâmetros, valores de retorno e funções nomeadas.
    - Funções anônimas e closures.

## Semana 2: Estruturas de Dados e Controle de Fluxo

1. Arrays, Slices e Maps

    - Arrays: definição e manipulação.
    - Slices: conceito, criação e operações básicas (slicing, append).
    - Maps: definição, inserção, remoção e iteração.

2. Estruturas e Métodos

    - Criação de structs.
    - Métodos em structs.
    - Uso de struct para representar entidades.

3. Ponteiros

    - Conceito de ponteiros em Go.
    - Referência e desreferência (* e &).
    - Passagem de ponteiros como argumentos de funções.

## Semana 3: Programação Orientada a Objetos e Interfaces

1. Interfaces

    - Conceito de interfaces.
    - Definição e implementação de interfaces.
    - Uso do polimorfismo em Go.

2. Composição

    - Em Go, não há herança, mas há composição.
    - Criação de estruturas compostas.
    - Embedding e métodos promovidos.

3. Pacotes e Módulos

    - Como organizar o código em pacotes.
    - Importação de pacotes.
    - Uso de módulos (go mod init, go.mod).

## Semana 4: Concorrência

1. Goroutines

    - Conceito de goroutines.
    - Criação de goroutines para executar tarefas em paralelo.

2. Canais (Channels)

    - Comunicação entre goroutines usando canais (chan).
    - Tipos de canais: unidirecionais e bidirecionais.
    - Buffers em canais.

3. Select Statement

    - Utilização do select para sincronização de canais.
    - Controle de múltiplas goroutines.

## Semana 5: Manipulação de Erros e Ferramentas

1. Tratamento de Erros

    - Convenção de tratamento de erros em Go (error interface).
    - Criação de erros personalizados (fmt.Errorf).

2. Pacotes Úteis

    - fmt (formatação e impressão).
    - os (manipulação de arquivos e variáveis de ambiente).
    - strconv, math, time.

3. Ferramentas do Ecossistema Go

    - `go fmt`: formatação do código.
    - `go ve`t: análise estática.
    - `golint`: análise de estilo.
    - `go test`: testes unitários.

## Semana 6: Testes e Documentação

1. Testes em Go

    - Escrever testes unitários usando testing.
    - Estrutura básica dos testes (TestXxx).
    - Testes de benchmark (BenchmarkXxx).

2. Documentação

    - Adicionar documentação ao seu código.
    - Gerar documentação com godoc.

3. Cobertura de Código

    - Como medir a cobertura dos testes.
    - Uso de `go test -cover`.

## Semana 7: Práticas Avançadas

1. Context Package

    - Uso do pacote context para gerenciar o tempo de vida de goroutines.
    - Cancelamento e deadlines.

2. Sincronização Avançada

    - Uso de sync.Mutex e sync.WaitGroup.
    - Quando usar sync/atomic para manipulação de variáveis compartilhadas.

## Semana 8: Aplicação Prática

1. Projeto Prático

    - Desenvolver uma API simples utilizando Go e o pacote net/http.
    - Criação de handlers para rotas.
    - Trabalhar com JSON (encoding/json).

2. Deploy de uma Aplicação Go

    - Compilar um binário para distribuição (go build).
    - Criar um Dockerfile para sua aplicação.
    - Deploy em um serviço como Heroku ou AWS.

## Semana 9: Go na Web

1. Servidor HTTP

    - Criação de um servidor HTTP com net/http.
    - Rotas e handlers.

2. Frameworks Web

    - Introdução ao framework Gin.
    - Criação de uma API RESTful utilizando Gin.

## Semana 10: Boas Práticas e Estudo Contínuo

1. Boas Práticas de Go

    - Convenções de código (idiomatic Go).
    - Organização de projetos Go.
    - Uso de ferramentas de análise estática.

2. Leitura de Código Fonte

    - Ler o código de bibliotecas populares.
    - Contribuir com projetos open source.

3. Materiais Recomendados
    - Livro: "The Go Programming Language" por Alan Donovan e Brian Kernighan.
    - Curso: Go (Golang) Programming - Complete Guide (Udemy ou Coursera).
    - Documentação Oficial: <https://golang.org/doc/>
    - Comunidade: Participe de fóruns como Reddit (/r/golang) ou Slack/Discord de desenvolvedores Go.
