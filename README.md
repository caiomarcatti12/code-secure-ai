# code-secure-ai

This repository contains documentation and a simple Go script that interacts with the SonarQube API.

## SonarQube Fetcher

`sonarqube_fetcher.go` fetches project information, issues and security hotspots from a SonarQube instance and logs the results.

### Running

1. Install Go 1.19 or newer.
2. Export the following environment variables:
   - `SONAR_URL` - base URL of your SonarQube server
   - `SONAR_TOKEN` - authentication token
   - `SONAR_PROJECT_KEY` - key of the project to fetch
3. Run the script:

```bash
go run sonarqube_fetcher.go
```
