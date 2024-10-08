# Frete Rápido API

## Descrição

API Rest desenvolvida em Go Lang para realizar cotações de frete utilizando a API da Frete Rápido.

## Como Executar

### Pré-requisitos

- Docker

### Executando a Aplicação

1. **Construir a imagem Docker:**

   ```bash
   docker build -t frete_rapido_api .

2. **Executar o container Docker:**

   ```bash
   docker run -d -p 8081:8080 frete_rapido_api   
