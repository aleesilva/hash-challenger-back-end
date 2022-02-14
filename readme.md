# hash-challenger-checkout Desenvolvido por Aleson França #

## Tecnologias Utilizadas Go, Grpc, Gin Framework ##

<hr>

## Base
Para criação do projeto utilizei alguns conceitos de clean architeture por mais que o projeto fosse básico

queria deixar as responsabilidade bem divididas então caso a gente queira utilizar outras libs no futuro isso esteja

desacoplado ao máximo.

<hr>

### Estrutura do Projeto
```
📦project
  ┗ 📦docker
      ┗ 📜Dockerfile
      📦config
        ┣ 📜env_vars.go
        ┣ 📜gin_server.go
        ┗ 📦routes
          ┣📜routes.go
      📦controllers
        ┗ cart_controller.go
      📦db
        ┗📜products.json
      📦helpers
        ┣📜black_friday.go
        ┣📜json_parse.go
      📦integrations
        ┣📜discount.go
      📦pb
        ┣📜discount_grpc.pb.go
        ┣📜discount.pb.go
      📦proto
        ┗ 📜discount.proto
      📦proto
        ┗ 📜product.go
      📦test
        ┣ 📜apply_discount_test.go
        ┣ 📜get_product_is_gift_test.go
        ┗ 📜product_details_test.go
      📦usecases
        ┣ 📜apply_discount.go
        ┣ 📜get_product_is_gift.go
        ┗ 📜product_details.go
      📜 .env.example
      📜.gitignore
      📜 docker-compose.yml
      📜 go.mod
      📜 main.go
      📜 readme.md
```


## Para montar  o ambiente, basta seguir os passos abaixo:

1 - Clonar o repositório: git clone https://github.com/aleesilva/hash-challenger-back-end.git

2 - Entrar na pasta do projeto: ``` cd hash-chacllenger-back-end ```

3 - Executar docker compose para subir os serviços: ``` docker-compose up ```

4 - Com isso, a aplicação estará rodando e escutando em localhost na porta 9001.


## Executar A API

*URL da aplicação:* http://localhost:9001/api/cart/checkout

Método da requisição é *POST*

```
{
  "products": [
    {
      "id": 4,
      "quantity": 8
    },
    {
      "id": 1,
      "quantity": 2
    }
  ]
}
```
Enviar o payload a cima utilizando o postman ou insomnia.
preenchendo o body com o payload no quadro a cima e utlizando a *URL da aplicação:* http://localhost:9001/api/cart/checkout
<hr>

## Execução dos testes

Alguns testes foram criados para os cenarios de usecases, porém poderiam ser mais completos mas para rodar os testes existentes basta rodar o seguinte comando : ``` cd test && go test -v ```

## Libs utilizadas
http : [Gin-Gonic](https://github.com/gin-gonic/gin)

grpc : [GRPC](https://grpc.io/docs/languages/go/quickstart/)
