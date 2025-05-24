package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Project struct {
	Key       string `json:"key"`
	Name      string `json:"name"`
	Qualifier string `json:"qualifier"`
}

type projectComponent struct {
	Component Project `json:"component"`
}

type Issue struct {
	Key       string `json:"key"`
	Rule      string `json:"rule"`
	Severity  string `json:"severity"`
	Component string `json:"component"`
	Message   string `json:"message"`
}

type issuesResponse struct {
	Total  int     `json:"total"`
	Issues []Issue `json:"issues"`
}

type Hotspot struct {
	Key       string `json:"key"`
	Component string `json:"component"`
	Message   string `json:"message"`
}

type hotspotsResponse struct {
	Paging struct {
		Total int `json:"total"`
	} `json:"paging"`
	Hotspots []Hotspot `json:"hotspots"`
}

func fetch(baseURL, token, path string, params url.Values, target interface{}) error {
	reqURL := fmt.Sprintf("%s%s?%s", baseURL, path, params.Encode())
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return err
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}
	return json.NewDecoder(resp.Body).Decode(target)
}

func main() {
	baseURL := os.Getenv("SONAR_URL")
	token := os.Getenv("SONAR_TOKEN")
	projectKey := os.Getenv("SONAR_PROJECT_KEY")
	if baseURL == "" || projectKey == "" {
		log.Fatal("SONAR_URL and SONAR_PROJECT_KEY must be set")
	}

	var proj projectComponent
	if err := fetch(baseURL, token, "/api/components/show", url.Values{"component": {projectKey}}, &proj); err != nil {
		log.Fatalf("fetch project: %v", err)
	}
	log.Printf("Project: %s (%s)", proj.Component.Name, proj.Component.Key)

	var issues issuesResponse
	if err := fetch(baseURL, token, "/api/issues/search", url.Values{"componentKeys": {projectKey}}, &issues); err != nil {
		log.Fatalf("fetch issues: %v", err)
	}
	for _, issue := range issues.Issues {
		log.Printf("Issue %s: %s [%s] - %s", issue.Key, issue.Component, issue.Severity, issue.Message)
	}

	var hotspots hotspotsResponse
	if err := fetch(baseURL, token, "/api/hotspots/search", url.Values{"projectKey": {projectKey}}, &hotspots); err != nil {
		log.Fatalf("fetch hotspots: %v", err)
	}
	for _, hs := range hotspots.Hotspots {
		log.Printf("Hotspot %s: %s - %s", hs.Key, hs.Component, hs.Message)
	}
}
