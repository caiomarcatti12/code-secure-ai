package parser

import (
    "encoding/json"
    "errors"
    "io"
    "os"

    "github.com/yourusername/dependency-check-automation/internal/domain"
)

// DependencyCheckReport models a simplified JSON structure from Dependency Check.
type DependencyCheckReport struct {
    Dependencies []struct {
        PackageName    string `json:"packageName"`
        Version        string `json:"version"`
        Vulnerabilities []struct {
            ID          int    `json:"id"`
            Severity    string `json:"severity"`
            Description string `json:"description"`
        } `json:"vulnerabilities"`
    } `json:"dependencies"`
}

type DependencyCheckParser interface {
    Parse(filePath string) ([]domain.Vulnerability, error)
}

type dependencyCheckParser struct{}

func NewDependencyCheckParser() DependencyCheckParser {
    return &dependencyCheckParser{}
}

func (p *dependencyCheckParser) Parse(filePath string) ([]domain.Vulnerability, error) {
    if filePath == "" {
        return nil, errors.New("filePath cannot be empty")
    }

    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    data, err := io.ReadAll(file)
    if err != nil {
        return nil, err
    }

    var report DependencyCheckReport
    if err = json.Unmarshal(data, &report); err != nil {
        return nil, err
    }

    var vulnerabilities []domain.Vulnerability

    for _, dep := range report.Dependencies {
        for _, v := range dep.Vulnerabilities {
            vuln, err := domain.NewVulnerability(
                v.ID,
                dep.PackageName,
                dep.Version,
                v.Severity,
                v.Description,
            )
            if err != nil {
                continue
            }
            vulnerabilities = append(vulnerabilities, vuln)
        }
    }

    return vulnerabilities, nil
}

