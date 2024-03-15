# fiber-graphql

## config Environment

```bash
cp .env.example .env
```

- then update environments in .env

## Run Project With Docker

```bash
docker network create Heepoke
```

```bash
docker compose up -d && docker compose logs api --follow
```

## Clear

```bash
go mod tidy
```
