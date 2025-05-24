package codexcli

import (
	"bytes"
	"os/exec"
)

// Client represents a wrapper around the codex CLI.
type Client struct {
	// Path is the path to the codex executable. If empty, "codex" is used.
	Path string
}

// RequestFix sends the given text to the codex CLI asking for a correction
// and prepares a pull request in the repository at repoPath.
func (c *Client) RequestFix(repoPath, text string) error {
	cmdPath := c.Path
	if cmdPath == "" {
		cmdPath = "codex"
	}

	// Example codex CLI usage. Assumes the codex CLI supports flags:
	//   --repo: path to repository
	//   --message: prompt to send
	//   --create-pr: create pull request automatically
	cmd := exec.Command(cmdPath, "--repo", repoPath, "--message", text, "--create-pr")

	// Capture output for debugging purposes.
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
