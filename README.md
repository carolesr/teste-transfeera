# Teste Trasfeera API

[Especificações do projeto](https://docs.google.com/document/d/1cSVj8EK7x2tfhr0IXn9shEAMRVRzRe5VHC-AjJUTWlk/edit#)

## Dependências

Para a execução desse projeto, é necessário ter instalado:

- Go versão 1.20
- Docker

## Setup

1- Instalar dependências

```
go mod tidy
```

2- Configurar container do banco de dados

```
docker-compose up
```

3- Inserção em massa de "seed" de dados iniciais no banco de dados

```
go run cmd/cli/main.go seed
```

4- Rodar API

```
go run cmd/server/main.go
```

## Testes

```
go test ./... -coverprofile coverage.out # coverage is optional
```

## Documentação da API

Os endpoints da API GraphQL estão listados em ```postman_collection.json```.

As queries e mutations podem ser executadas no Playground, acessando ```localhost:8080/api/v1/playground```, o qual também contém a documentação do schema graphql do projeto.

### listReceivers

Este endpoint retorna a lista paginada de receivers existentes do banco de dados.

A paginação apresenta 10 registros por padrão, mas esse número pode ser personalizado através do parâmetro ```first``` no input da query.
Para avançar na paginação, deve-se incluir o parâmetro ```after``` com o cursor do último registro apresentado na página.

É possível filtrar os registros por Nome, Status, Tipo de chave e Valor de chave através dos parâmetros ```name```, ```status```, ```key_type``` e ```type```.

### receiver

Este endpoint retorna o receiver correspondente ao campo ```id``` enviado na query.

### createReceiver

Este endpoint cria um novo receiver contendo os dados de Nome, Email, Identificador, Tipo de chave Pix e Valor de chave Pix, 
```name```, ```email```, ```identifier```, ```pixKeyType``` e ```pixKey```, enviados no input da mutation.

O campo ```pixKeyType``` aceita as opções "CPF", "CNPJ", "EMAIL", "TELEFONE" e "CHAVE_ALEATORIA".

O campo ```pixKey``` é validado conforme o valor do ```pixKeyType```.

O campo ```identifier```, aceita tanto dados de CPF quanto de CNPJ.

O campo ```email```, aceita o mesmo formato de email que a chave Pix de tipo Email, e tem um limite de 250 caracteres.

O receiver é criado com o campo Status com valor ```Draft``` (Rascunho).

### updateReceiver

Este endpoint atualiza os dados do receiver correspondente ao campo ```id``` enviado na mutation.

É possível atualizar os campos ```name```, ```email```, ```identifier```, ```pixKeyType``` e ```pixKey```, e é necessário enviar ao menos um deles para a execução da atualização.

Todos esses campos possuem as mesmas validações aplicadas na mutation ```createReceiver```.

Para atualizar o campo ```pixKeyType```, é necessário atualizar também o campo ```pixKey```.

Receivers com Status ```Draft``` podem ter todos esses campos atualizados, mas receivers com Status ```Validated``` podem terão somente o campo ```email``` atualizado.

### deleteReceiver

Este endpoint exclui um ou mais receivers, correspondentes ao campo ```ids``` enviados na mutation.

A exclusão é feita com soft delete, de modo que os dados ainda existem no banco, mas não são mais exibidos nas queries de listagem, e não são mais passíveis de atualização.
