# Desafio Sistema de Temperatura por CEP

Este repositório foi criado exclusivamente para hospedar o código do desenvolvimento do Desfio do Sistema de temperatura por CEP da da **Pós Go Expert**, ministrado pela **Full Cycle**.

## Descrição do Desafio

A seguir estão os dados fornecidos na descrição do desafio.

### Objetivo

Desenvolver um sistema em Go que receba um CEP, identifica a cidade e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin). Esse sistema deverá ser publicado no Google Cloud Run.

### Requisitos

- O sistema deve receber um CEP válido de 8 digitos
- O sistema deve realizar a pesquisa do CEP e encontrar o nome da localização, a partir disso, deverá retornar as temperaturas e formata-lás em: Celsius, Fahrenheit, Kelvin.
- O sistema deve responder adequadamente nos seguintes cenários:
    - Em caso de sucesso:
        - Código HTTP: 200
        - Response Body: { "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }
    - Em caso de falha, caso o CEP não seja válido (com formato correto):
        - Código HTTP: 422
        - Mensagem: invalid zipcode
    - Em caso de falha, caso o CEP não seja encontrado:
        - Código HTTP: 404
        - Mensagem: can not find zipcode
- Deverá ser realizado o deploy no Google Cloud Run.

### Dicas

- Utilize a API viaCEP (ou similar) para encontrar a localização que deseja consultar temperatura: https://viacep.com.br/
- Utilize a API WeatherAPI (ou similar) para consultar as temperaturas desejadas: https://www.weatherapi.com/
- Para realizar a conversão de Celsius para Fahrenheit, utilize a seguinte fórmula: F = C * 1,8 + 32
- Para realizar a conversão de Celsius para Kelvin, utilize a seguinte fórmula: K = C + 273
    - Sendo F = Fahrenheit
    - Sendo C = Celsius
    - Sendo K = Kelvin

### Entrega

- O código-fonte completo da implementação.
- Testes automatizados demonstrando o funcionamento.
- Utilize docker/docker-compose para que possamos realizar os testes de sua aplicação.
- Deploy realizado no Google Cloud Run (free tier) e endereço ativo para ser acessado.


## Execução do Desafio

### Testes do Webserver

Para a realização dos testes, basta executar o seguinte comando `go test ./internal/infra/webserver/handlers -v` a partir da raiz do repositório. Abaixo segue o exemplo da execução:



```bash

wander@bsnote283:~/desafio-temperatura-cep$ go test ./internal/infra/webserver/handlers -v
=== RUN   TestBuscaTemperaturaHandlerOk
--- PASS: TestBuscaTemperaturaHandlerOk (1.11s)
=== RUN   TestBuscaTemperaturaHandlerOkCaractereEspecial
--- PASS: TestBuscaTemperaturaHandlerOkCaractereEspecial (0.72s)
=== RUN   TestBuscaTemperaturaHandlerCepInvalido
2024/06/18 21:45:09 invalid zipcode: 324500000
--- PASS: TestBuscaTemperaturaHandlerCepInvalido (0.00s)
=== RUN   TestBuscaTemperaturaHandlerCepNaoEncontrado
2024/06/18 21:45:09 can not find zipcode: 00000000
--- PASS: TestBuscaTemperaturaHandlerCepNaoEncontrado (0.20s)
PASS
ok  	github.com/wandermaia/desafio-temperatura-cep/internal/infra/webserver/handlers	(cached)
wander@bsnote283:~/desafio-temperatura-cep$ 


```

Também foi criado o arquivo `api/apis_temperatura_cep.http` para que os endpoints possam ser testados diretamente a partir do VScode. Para realizar o teste utilizando o arquivo, basta executar o comando abaixo para inicializar o webserver e, em seguida utilizar os endpoints cadastrados:


```bash

wander@bsnote283:~/desafio-temperatura-cep$ go mod tidy
wander@bsnote283:~/desafio-temperatura-cep$ 
wander@bsnote283:~/desafio-temperatura-cep$ go run cmd/server/main.go 
2024/06/18 22:32:45 Servidor iniciado na porta 8080!


```

Ao executar esses comandos, os módulos serão baixados e o webserver será inciado.


### Criação do Container


Para gerar um container da aplicação, basta executar os seguintes comandos a partir da raiz do projeto:


```bash

# Gerar uma nova imagem
docker build -t wandermaia/desafio-temperatura-cep:latest -f Dockerfile.prod .

# Verificar a imagem gerada
docker images | grep desafio-temperatura-cep

# Executar o container
docker run --rm -p 8080:8080 wandermaia/desafio-temperatura-cep:latest

```

Com isso, o webserver será iniciado na porta 8080. Para realizar testes no container, o próprio arquivo `api/apis_temperatura_cep.http` pode ser utilizado.


### Google Cloud Run


O deploy foi realizado uma conta gratuíta no Google Cloud utilizando o projeto default. Abaixo segue o print do deploy do projeto já em execução no Google Cloud Run:


![gcp-service.png](/.img/gcp-service.png)


Para acessar o projeto, pode ser utilizado o link https://desafio-temperatura-cep-u6wtscwu3a-uc.a.run.app . Abaixo segue o print de alguns testes utilizando o curl:


![gcp-test.png](/.img/gcp-test.png)


Os comandos utilizados para o teste acima estão abaixo:



```bash

# Cep válido
curl -i https://desafio-temperatura-cep-u6wtscwu3a-uc.a.run.app/32450000

# Cep inválido
curl -i https://desafio-temperatura-cep-u6wtscwu3a-uc.a.run.app/324500000

# Cep com formato válido, mas inexistente
curl -i https://desafio-temperatura-cep-u6wtscwu3a-uc.a.run.app/000000000

```


