# code-secure-ai

This repository contains a simple Go service that interacts with the OpenAI API. It accepts a full description of a SonarQube issue and returns additional guidance for fixing it. The service exposes a `/details` endpoint that expects a JSON body with the field `issue`.

## Building

Ensure you have Go installed. Then, run:

```bash
go build
```

## Running

Set the `OPENAI_API_KEY` environment variable with your API key and start the service:

```bash
OPENAI_API_KEY=yourkey go run .
```

The service listens on port `8080`.

## Request Example

Send a POST request to `/details` with a JSON payload:

```json
{
  "issue": "<descricao completa da issue do SonarQube, com referencia de arquivo e trecho de codigo>"
}
```

The response will contain a `details` field with the explanation returned by OpenAI.
