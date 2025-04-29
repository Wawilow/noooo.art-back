package config

import (
	"context"
	"os"

	infisical "github.com/infisical/go-sdk"
)

type Config struct {
	INFISICAL_CLIENT_ID     string
	INFISICAL_CLIENT_SECRET string
	INFISICAL_PROJECT_ID    string
	INFISICAL_SITE_URL      string

	SUPABASE_DSN *string
}

func getEnvVariables() (string, string, string, string) {
	INFISICAL_CLIENT_ID := os.Getenv("INFISICAL_CLIENT_ID")
	INFISICAL_CLIENT_SECRET := os.Getenv("INFISICAL_CLIENT_SECRET")
	INFISICAL_PROJECT_ID := os.Getenv("INFISICAL_PROJECT_ID")
	INFISICAL_SITE_URL := os.Getenv("INFISICAL_SITE_URL")

	return INFISICAL_SITE_URL, INFISICAL_CLIENT_ID, INFISICAL_CLIENT_SECRET, INFISICAL_PROJECT_ID
}

func NewConfig() (*Config, error) {
	INFISICAL_SITE_URL, INFISICAL_CLIENT_ID, INFISICAL_CLIENT_SECRET, INFISICAL_PROJECT_ID := getEnvVariables()
	res := Config{}
	res.INFISICAL_SITE_URL = INFISICAL_SITE_URL
	res.INFISICAL_CLIENT_ID = INFISICAL_CLIENT_ID
	res.INFISICAL_CLIENT_SECRET = INFISICAL_CLIENT_SECRET
	res.INFISICAL_PROJECT_ID = INFISICAL_PROJECT_ID

	client := infisical.NewInfisicalClient(context.Background(), infisical.Config{
		SiteUrl:          INFISICAL_SITE_URL,
		AutoTokenRefresh: true,
	})
	_, err := client.Auth().UniversalAuthLogin(INFISICAL_CLIENT_ID, INFISICAL_CLIENT_SECRET)
	if err != nil {
		return nil, err
	}

	secret, err := client.Secrets().Retrieve(infisical.RetrieveSecretOptions{
		SecretKey:   "SUPABASE_DSN",
		Environment: "prod",
		ProjectID:   INFISICAL_PROJECT_ID,
		SecretPath:  "/",
	})
	if err != nil {
		return nil, err
	}
	res.SUPABASE_DSN = &secret.SecretValue

	return &res, nil
}
