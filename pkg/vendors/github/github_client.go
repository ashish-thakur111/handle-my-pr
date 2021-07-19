package github

import (
	"context"

	"github.com/google/go-github/v37/github"
	"golang.org/x/oauth2"
)

func NewGithubClientWithAuth(token string) *github.Client {
	ctx := context.Background()
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})

	tokenClient := oauth2.NewClient(ctx, tokenSource)

	return github.NewClient(tokenClient)
}

func NewGithubClientWithoutAuth() *github.Client {
	return github.NewClient(nil)
}
