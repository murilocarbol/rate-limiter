# Desafio 1 - Go Expert -> Rate Limiter
Sistema de rate limiter que realiza o controle de quantidade de requisições de acordo com Client IP ou Token informado.

Desafio Pós Go Expert - 2024 Desafios Técnicos - Rate Limiter - FullCycle

### Como Utilizar localmente:
#### Requisitos:
    - Certifique-se de ter o Go instalado em sua máquina.
    - Certifique-se de ter o Docker instalado em sua máquina.

  1. Clonar o Repositório:~
  ```git clone https://github.com/murilocarbol/rate-limiter.git```

  2. Rode o docker para buildar a imagem gerando o container com a aplicação e o redis:
  ```docker-compose up```

### Como testar localmente:
Porta: HTTP server on port :8080

#### Execute o curl abaixo ou use um aplicação client REST para realizar a requisição:

    curl --request GET \
    --url http://localhost:8080/ \
    --header 'API_KEY: abc123'