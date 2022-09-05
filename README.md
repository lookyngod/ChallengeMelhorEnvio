# Challenge Melhor Envio
## Descrição do desafio

O arquivo [logs.txt](https://drive.google.com/file/d/1b9mpy5fXb3yQwcRDu03-pTBA2QCcLLTK/) contém informações de log geradas por um sistema de API Gateway.
Cada solicitação foi registrada em um objeto JSON separado por uma nova linha \n

## Requisitos
- Processar o arquivo de log, extrair informações e salvá-las no banco de dados.
- Gerar um relatório para cada descrição abaixo, em formato csv:
- Requisições por consumidor;
- Requisições por serviço;
- Tempo médio de request , proxy e gateway por serviço.
- Documentar passo a passo de como executar o teste através de um arquivo README.md.
- Efetue o commit de todos os passos do desenvolvimento em um git público de sua preferência e disponibilize apenas o link para o repositório.

## Decisões Técnicas

Utilizei algumas das boas práticas de clean code e as premissas do [Uber - Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md).

Utilizei também algumas premissas da arquitetura "go-scaffold" onde tenho familiaridade que pode ser vista aqui: [go-scaffold](https://pkg.go.dev/github.com/facily-tech/go-scaffold)

## Instalação

A aplicação necessita de um ambiente com [Golang](https://go.dev/doc/install) 1.17+ para rodar.

Necessita também do [Docker Compose](https://docs.docker.com/compose/install/)

Instale as dependências e para rodar a aplicação use o passo-a-passo abaixo:

### Passo 1:
Faça o clone deste projeto, baixe o arquivo [logs.txt](https://drive.google.com/file/d/1b9mpy5fXb3yQwcRDu03-pTBA2QCcLLTK/) e mova para a pasta raiz do projeto.

### Passo 2:
Entre na pasta raiz do projeto, pois subiremos nosso banco de dados com nosso docker-compose, rode o comando:
```sh
docker-compose up -d
```

Caso ocorra o erro: "Error while fetching server API version: ('Connection aborted.', PermissionError(13, 'Permission denied'))" rode o comando:
```sh
sudo chmod 777 /var/run/docker.sock
```

### Passo 3:
Certifique-se que o arquivo logs.txt esteja na raiz do projeto e rode o seguinte comando:
```sh
go run cmd/cli/main.go < logs.txt
```

### Passo 4:
Aguarde a finalização do processo, após finalização seus arquivos CSV serão gerados e os logs estarão no banco de dados.

Poderá conferir via PhpMyAdmin os lançamento dos logs no banco de dados > localhost:8081 usando as credenciais
- utilizador: root
- palavra-passe: Melhor1#

## Referências

- Medium/Youtube
- StackOverflow
- Udemy
- [GoPlayGround](https://go.dev/play/p/7N4sxD-Bai)
- [gorm.io](https://gorm.io/docs/connecting_to_the_database.html)

## Notas 

Desafio muito legal e que me deixou bastante empenhado em entregar da melhor forma.
