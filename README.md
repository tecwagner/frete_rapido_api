# Frete Rápido API

## Descrição

API Rest desenvolvida em Go Lang para realizar cotações de frete utilizando a API da Frete Rápido.

## Como Executar

### Pré-requisitos

- Docker
- Docker Compose
- Go Lang: 1.23
- API REST
- API Frete Rápido Simulação
- Banco de Dados Teste: SQlite 
- Banco de Dados Desenvolvimento: PostgreSQL
- ORM: Gorm

### Executando a Aplicação

1. **Construir a imagem Docker:**

   ```bash
   docker build -t frete_rapido_api .

2. **Executar o comando para permitir acesso docker/dbdata:**

   ```bash
   sudo chown -R $USER:$USER .docker/dbdata   

3. **Executar o container Docker:**

   ```bash
   docker compose up --build   

4. **Ao Executar o projeto para testar as rotas: api > client.http**   

   - Solicitar cotação:
      - POST http://localhost:8081/api/v1/quote

