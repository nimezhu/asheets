package asheets

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"golang.org/x/oauth2"
)

type GAgent struct {
	root string
}

func NewGAgent(root string) GAgent {
	return GAgent{root}
}
func (agent GAgent) GetClient(ctx context.Context, config *oauth2.Config) *http.Client {
	cacheFile, err := agent.tokenCacheFile()
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
	}
	tok, err := tokenFromFile(cacheFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(cacheFile, tok)
	}
	return config.Client(ctx, tok)
}
func (agent GAgent) HasCacheFile() bool {
	cacheFile, err := agent.tokenCacheFile()
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
		return false
	}
	if _, err := tokenFromFile(cacheFile); err != nil {
		return false
	}
	return true
}

func (agent GAgent) tokenCacheFile() (string, error) {
	tokenCacheDir := filepath.Join(agent.root, "credentials")
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir,
		url.QueryEscape("gsheets.json")), nil
}
