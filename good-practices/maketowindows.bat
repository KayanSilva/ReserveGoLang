@echo off
REM Script para rodar testes e lint no Windows

REM Executa o teste no container Docker
echo Executando test...
docker compose exec api go test ./test -v

REM Executa o lint
echo Executando lint...
docker run -t --rm -v "%cd%":/app -w /app golangci/golangci-lint:v2.6.1 golangci-lint run

REM Executa o formatting
echo Executando formatting...
docker run -t --rm -v "%cd%":/app -w /app golangci/golangci-lint:v2.6.1 golangci-lint fmt

echo Todos os comandos foram executados.
