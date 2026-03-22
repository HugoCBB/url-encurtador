# 🔗 URL Encurtador

Serviço simples e eficiente de encurtamento de URLs construído em **Go**, utilizando **Redis** como armazenamento temporário. Os links encurtados expiram automaticamente após **24 horas**.

## ✨ Funcionalidades

- Encurta qualquer URL via requisição `POST`
- Redireciona automaticamente para a URL original via `GET`
- Links com expiração automática de 24 horas
- Health check endpoint
- Infraestrutura containerizada com Docker

## 🛠️ Tecnologias

| Tecnologia | Uso |
|---|---|
| [Go](https://golang.org/) | Linguagem principal |
| [Redis](https://redis.io/) | Armazenamento das URLs |
| [go-redis/v9](https://github.com/redis/go-redis) | Cliente Redis para Go |
| [Docker](https://www.docker.com/) | Containerização |

## 📁 Estrutura do Projeto

```
url-encurtador/
├── cmd/
│   ├── api/v1/
│   │   └── routes.go         # Definição das rotas HTTP
│   └── server/
│       └── main.go           # Ponto de entrada da aplicação
├── internal/
│   ├── config/
│   │   └── redis-client.go   # Inicialização do cliente Redis
│   ├── controllers/
│   │   └── url.controller.go # Handlers das requisições HTTP
│   ├── models/
│   │   └── url.go            # Definição dos modelos de dados
│   └── repository/
│       └── url.repository.go # Acesso ao Redis e geração de short codes
├── Dockerfile
├── docker-compose.yaml
├── Makefile
└── go.mod
```

## 🚀 Rodando o Projeto

### Pré-requisitos

- [Go 1.21+](https://golang.org/dl/)
- [Docker](https://www.docker.com/) e [Docker Compose](https://docs.docker.com/compose/)
- [Make](https://www.gnu.org/software/make/)

### Com Make (recomendado)

```bash
# Sobe o Redis e inicia o servidor Go
make run

# Para os containers Docker
make stop

# Limpa binários e volumes Docker
make clean
```

O servidor estará disponível em `http://localhost:8080`.

## 📡 Endpoints

### Health Check

```
GET /
```

**Resposta:**
```json
{"status": "ok"}
```

---

### Encurtar URL

```
POST /
Content-Type: application/json
```

**Body:**
```json
{
  "url": "https://www.exemplo.com/um-link-muito-longo"
}
```

**Resposta:**
```json
{
  "url_converted": "http://localhost:8080/a1b2c3d4",
  "exp": "24h0m0s"
}
```

---

### Redirecionar

```
GET /{code}
```

Redireciona automaticamente (`301 Moved Permanently`) para a URL original correspondente ao código.

**Exemplo:**
```
GET /a1b2c3d4  →  302 redirect para https://www.exemplo.com/um-link-muito-longo
```

## ⚙️ Como Funciona

1. A URL recebida é processada com **MD5** e os primeiros 8 caracteres do hash hex são usados como `short code`.
2. O mapeamento `{short_code → dados da URL}` é salvo no Redis com TTL de **24 horas**.
3. Ao acessar o link encurtado, o servidor busca o `short code` no Redis e redireciona para a URL original.

## 🐳 Docker

O `Dockerfile` usa **multi-stage build** para gerar um binário mínimo baseado na imagem `scratch`:

```bash
# Build da imagem
docker build -t url-encurtador .

# Rodar o container (requer Redis acessível)
docker run -p 8080:8080 url-encurtador
```

Para subir a stack completa (Redis + API), descomente o serviço `backend` no `docker-compose.yaml` e rode:

```bash
docker compose up
```
