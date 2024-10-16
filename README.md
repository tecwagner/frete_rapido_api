# Frete Rápido API

## Descrição

API Rest desenvolvida em Go Lang para realizar cotações de frete utilizando a API da Frete Rápido.

## Como Executar

### Pré-requisitos

- **Docker**: Ferramenta de contêiner para empacotar a aplicação.
- **Docker Compose**: Utilizado para orquestrar a execução dos serviços, como o banco de dados.
- **GoLang (1.23)**: Linguagem de programação utilizada no projeto.
- **Testify**: Framework para realizar testes unitários.
- **Mock**: Biblioteca de mocking para simulação de dependências nos testes.
- **API Frete Rápido (Simulação)**: API externa para simular cotações de frete.
- **Banco de Dados (Teste: SQLite)**: Usado para testes locais.
- **Banco de Dados (Desenvolvimento: PostgreSQL)**: Utilizado em desenvolvimento.
- **ORM (Gorm)**: Mapeamento objeto-relacional utilizado no projeto.


### Executando a Aplicação

1. **Construir a imagem Docker:**

   ```bash
   docker build -t frete_rapido_api .

2. **Ajustar permissões para o diretório do Docker:**

   ```bash
   sudo chown -R $USER:$USER .docker/dbdata

3. **Executar o container Docker:**

   ```bash
   docker compose up --build   

4. **Acessar o Banco de Dados (Postgres) no PgAdmin:**

   - Acesse o URL: [http://localhost:9000/login?next=%2F](http://localhost:9000/login?next=%2F)
   - Credenciais de login:
     - **Usuário**: `admin@user.com`
     - **Senha**: `123456`

   - Após login, acesse [http://localhost:9000/browser/](http://localhost:9000/browser/) e registre a instância do banco:
     - **Nome**: `db`
     - **Hostname**: `database`
     - **Porta**: `5432`
     - **Banco de Dados**: `freterapido`
     - **Usuário**: `postgres`
     - **Senha**: `postgres`


5. **Testar as Rotas da API:**

Use ferramentas como **Postman** ou **cURL** para testar as rotas da API:

- **Extenção ResFul para VSCode na raiz do projeto:**:**

   - api > client.http

- **Solicitar Cotação:**
  - Método: `POST`
  - URL: `http://localhost:8081/api/v1/quote`
  - Exemplo de Payload:
    ```json
    {
      "shipper": {
         "registered_number": "25438296000158",
         "token": "1d52a9b6b78cf07b0886152459a5c90",
         "platform_code": "5AKVkHqCn"
      },
      "recipient": {
         "type": 1,
         "country": "BRA",
         "zipcode": 29161376
      },
      "dispatchers": [
         {
               "registered_number": "25438296000158",
               "zipcode": 29161376,
               "total_price": 0.0,
               "volumes": [
                  {
                     "amount": 1,
                     "category": "7",
                     "sku": "abc-teste-123",
                     "height": 0.2,
                     "width": 0.2,
                     "length": 0.2,
                     "unitary_price": 359,
                     "unitary_weight": 5
                  },
                  {
                     "amount": 2,
                     "category": "7",
                     "sku": "abc-teste-527",
                     "height": 0.4,
                     "width": 0.6,
                     "length": 0.15,
                     "unitary_price": 556,
                     "unitary_weight": 4
                  }
               ]
         }
      ],
      "simulation_type": [
         0
      ]
   }
    ```

- **Buscar Métricas de Cotações (com parâmetro):**
  - Método: `GET`
  - URL: `http://localhost:8081/api/v1/metrics?last_quotes=5`

- **Buscar Métricas de Cotações (sem parâmetro):**
  - Método: `GET`
  - URL: `http://localhost:8081/api/v1/metrics`


6. **Comando para Executar os Testes:**

   - **Remover o arquivo de cobertura existente:**

      - rm coverage.out

   - **Remover cache de build do Go:**

      - go clean -testcache

   - **Executar novamente os testes com cobertura:**

      - go test -coverprofile=coverage.out ./...

   - **Gerar um arquivo HTML com o relatório de cobertura:**

      - go tool cover -html=coverage.out
