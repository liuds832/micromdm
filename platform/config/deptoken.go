package config

import (
	"time"

	"github.com/liuds832/micromdm/dep"
)

const DEPTokenTopic = "mdm.TokenAdded"

type DEPToken struct {
	ConsumerKey       string    `json:"consumer_key"`
	ConsumerSecret    string    `json:"consumer_secret"`
	AccessToken       string    `json:"access_token"`
	AccessSecret      string    `json:"access_secret"`
	AccessTokenExpiry time.Time `json:"access_token_expiry"`
}

// create a DEP client from token.
func (tok DEPToken) Client() (*dep.Client, error) {
	conf := dep.OAuthParameters{
		ConsumerKey:    tok.ConsumerKey,
		ConsumerSecret: tok.ConsumerSecret,
		AccessSecret:   tok.AccessSecret,
		AccessToken:    tok.AccessToken,
	}
	client := dep.NewClient(conf)
	return client, nil
}
