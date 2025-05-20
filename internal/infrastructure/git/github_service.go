package git

import "fmt"

// GitService defines a contract for creating pull requests.
type GitService interface {
    CreatePullRequest(branch, title, description string) (bool, error)
}

type githubService struct{}

func NewGitHubService() GitService {
    return &githubService{}
}

func (g *githubService) CreatePullRequest(branch, title, description string) (bool, error) {
    fmt.Printf("Creating PR on branch %s:\nTitle: %s\nDescription:\n%s\n\n", branch, title, description)
    return true, nil
}

