package auth

import (
	"fmt"

	"golang.org/x/oauth2"
)

type FileTokenSource struct {
	Store       *FileTokenStore
	Underlying  oauth2.TokenSource
	cachedToken *oauth2.Token
}

func (ts *FileTokenSource) Token() (*oauth2.Token, error) {
	newToken, err := ts.Underlying.Token()
	if err != nil {
		return nil, err
	}

	if ts.cachedToken == nil || newToken.AccessToken != ts.cachedToken.AccessToken {
		if err := ts.Store.SaveToken(newToken); err != nil {
			// Log the error but continue; do not block the active web request
			fmt.Printf("Warning: failed to persist token: %v\n", err)
		}
		ts.cachedToken = newToken
	}

	return newToken, nil
}
