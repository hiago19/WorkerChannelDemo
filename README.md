# appWorkerChannelDemo

## Sobre o Projeto

O **appWorkerChannelDemo** é uma aplicação escrita em Golang que demonstra o uso de routines e canais para processamento concorrente. Ele cria múltiplos workers que processam dados recebidos de um canal.

## Estrutura do Projeto

- **main.go:** Este arquivo contém a implementação principal do projeto. Ele define uma função worker que processa dados de um canal e a função main que cria vários workers e envia dados para o canal.

- **main.exe:** Arquivo executável gerado após a compilação do código Go.

## Tecnologias Utilizadas

O projeto foi desenvolvido em Golang, utilizando apenas as bibliotecas padrão da linguagem.

## Rodando Localmente
Para executar o projeto localmente, siga estas etapas:

1. Certifique-se de ter o Golang instalado em sua máquina.

2. Clone o repositório para sua máquina local:

   ```bash
   git clone https://github.com/hiago19/appWorkerChannelDemo.git
   ```

3. Navegue até o diretório do projeto:

   ```bash
   cd appWorkerChannelDemo
   ```

4. Compile o projeto executando o comando:

  ```bash
  go build -o main.exe main.go
  ```

5. Execute o arquivo `main.go`:

   ```bash
   ./main.exe
   ```

5. Siga as instruções no terminal para inserir a quantidade de workers e execuções.

## Funcionalidades

- Criação de múltiplos workers para processamento concorrente.
- Envio de dados para os workers através de um canal.
- Interação com o usuário para definir a quantidade de workers e execuções.

## Demonstração

Aqui está um exemplo de execução do programa:

![WorkerChannelDemo](https://github.com/hiago19/appWorkerChannelDemo/assets/81202387/9b300830-45b2-434e-b90e-a9bbf0373d84)




---
