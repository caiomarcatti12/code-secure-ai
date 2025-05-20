package main

import (
    "log"

    "github.com/yourusername/dependency-check-automation/internal/application"
    "github.com/yourusername/dependency-check-automation/internal/infrastructure/ai"
    "github.com/yourusername/dependency-check-automation/internal/infrastructure/git"
    "github.com/yourusername/dependency-check-automation/internal/infrastructure/parser"
    "github.com/yourusername/dependency-check-automation/internal/web"
)

func main() {
    dpParser := parser.NewDependencyCheckParser()
    openAi := ai.NewOpenAIService()
    gitHub := git.NewGitHubService()

    vulnService := application.NewVulnerabilityService(dpParser, openAi, gitHub)

    server := web.NewServer(vulnService)
    log.Println("Starting server on :8080")
    if err := server.Start(); err != nil {
        log.Fatal(err)
    }
}

