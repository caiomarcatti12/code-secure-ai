# Dependency Check Automation (Example)

Este repositório apresenta um exemplo conceitual de aplicação em Go para ler relatórios do **OWASP Dependency Check**, gerar sugestões de correção via serviço de IA (simulado) e criar Pull Requests automáticos. O projeto segue uma organização inspirada em Clean Architecture.

## Estrutura de Pastas

```
.
├── cmd
│   └── main.go
├── internal
│   ├── domain
│   ├── application
│   ├── infrastructure
│   │   ├── ai
│   │   ├── git
│   │   └── parser
│   └── web
└── go.mod
```

- **cmd**: ponto de entrada da aplicação.
- **internal/domain**: entidades de negócio, como `Vulnerability`.
- **internal/application**: coordena o fluxo principal do caso de uso.
- **internal/infrastructure**: implementações de parser, IA e Git.
- **internal/web**: servidor HTTP para disparar o processamento.

## Executando

1. Gere um relatório do Dependency Check em JSON.
2. Execute o binário (ou `go run cmd/main.go`) e envie um POST para `http://localhost:8080/process` informando o caminho do relatório:

```bash
curl -X POST -H "Content-Type: application/json" \
     -d '{"report_path":"/caminho/para/dependency-check-report.json"}' \
     http://localhost:8080/process
```

O serviço processará o arquivo, gerará sugestões fictícias de correção e exibirá (no console) a criação de um Pull Request.

## Observações

- As chamadas à API da OpenAI e ao GitHub estão simuladas para fins didáticos.
- Ajustes de segurança, tratamento de erros e validação de entradas são recomendados para uso real.

