package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type request struct {
	Issue string `json:"issue"`
}

type openaiMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type openaiRequest struct {
	Model    string          `json:"model"`
	Messages []openaiMessage `json:"messages"`
}

type openaiChoice struct {
	Message openaiMessage `json:"message"`
}

type openaiResponse struct {
	Choices []openaiChoice `json:"choices"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		http.Error(w, "missing OPENAI_API_KEY", http.StatusInternalServerError)
		return
	}

	oreq := openaiRequest{
		Model: "gpt-3.5-turbo",
		Messages: []openaiMessage{
			{Role: "system", Content: "Voce auxilia desenvolvedores a corrigir problemas de seguranca no codigo."},
			{Role: "user", Content: "Forneca mais detalhes para corrigir o seguinte problema detectado pelo SonarQube:\n\n" + req.Issue},
		},
	}
	body, err := json.Marshal(oreq)
	if err != nil {
		http.Error(w, "error encoding request", http.StatusInternalServerError)
		return
	}

	httpReq, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, "error creating request", http.StatusInternalServerError)
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		http.Error(w, "error calling openai", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(resp.Body)
		http.Error(w, "openai error: "+string(b), http.StatusInternalServerError)
		return
	}

	var oresp openaiResponse
	if err := json.NewDecoder(resp.Body).Decode(&oresp); err != nil {
		http.Error(w, "error decoding response", http.StatusInternalServerError)
		return
	}
	if len(oresp.Choices) == 0 {
		http.Error(w, "openai returned no choices", http.StatusInternalServerError)
		return
	}

	result := map[string]string{"details": oresp.Choices[0].Message.Content}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/details", handler)
	log.Println("listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
