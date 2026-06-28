package auth

import (
	"encoding/json"
	"os"

	"golang.org/x/oauth2"
)

type DBTokenStore interface {
	SaveToken(token *oauth2.Token) error
	GetToken() (*oauth2.Token, error)
}

type FileTokenStore struct {
	Path string
}

func (fs *FileTokenStore) SaveToken(token *oauth2.Token) error {
	file, err := os.Create(fs.Path)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(token); err != nil {
		return err
	}

	return nil
}

func (fs *FileTokenStore) GetToken() (*oauth2.Token, error) {
	content, err := os.ReadFile(fs.Path)
	if err != nil {
		return nil, err
	}

	var token *oauth2.Token
	err = json.Unmarshal(content, token)
	if err != nil {
		return nil, err
	}

	return token, nil
}
