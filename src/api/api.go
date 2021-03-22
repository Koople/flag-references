package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	rhttp "github.com/hashicorp/go-retryablehttp"
	"github.com/koople/flag-references/src/searcher"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	BaseUri = "https://sdk.koople.io"
)

type KPLApi struct {
	httpClient *rhttp.Client
	options    KPLOptions
}

type KPLOptions struct {
	BaseUri string
	ApiKey  string
	Logger  *logrus.Logger
}

type FlagFound struct {
	Flag   string          `json:"flag"`
	Founds []searcher.File `json:"founds"`
}

type RepositoryReferences struct {
	Repository string      `json:"repository"`
	Branch     string      `json:"branch"`
	References []FlagFound `json:"references"`
}

func NewClient(opts KPLOptions) KPLApi {
	if opts.Logger == nil {
		opts.Logger = logrus.New()
	}

	if opts.BaseUri == "" {
		opts.BaseUri = BaseUri
	}

	env, exists := os.LookupEnv("API_URL")
	if exists {
		opts.BaseUri = env
	}

	httpClient := rhttp.NewClient()
	httpClient.Logger = opts.Logger

	return KPLApi{
		httpClient: httpClient,
		options:    opts,
	}
}

func (client KPLApi) GetListFlags() ([]string, error) {
	flagsPath := "/proxy/server/flags"
	request, err := rhttp.NewRequest("GET", fmt.Sprintf("%s%s", client.options.BaseUri, flagsPath), nil)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating request GetFlags")
	}
	request.Header.Set("x-api-key", client.options.ApiKey)
	request.Header.Set("Content-Type", "application/json")

	resp, err := client.httpClient.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "Error doing the request GetFlags")
	}

	defer resp.Body.Close()

	var flags []string
	err = json.NewDecoder(resp.Body).Decode(&flags)
	return flags, err
}

func (client KPLApi) SaveFlagsInformation(founds RepositoryReferences) error {
	flagsPath := "/proxy/server/flags"
	jsonValue, err := json.Marshal(founds)
	if err != nil {
		return errors.Wrap(err, "Error marshaling founds list")
	}

	request, err := rhttp.NewRequest("PUT", fmt.Sprintf("%s%s", client.options.BaseUri, flagsPath), bytes.NewReader(jsonValue))
	if err != nil {
		return errors.Wrap(err, "Error creating request UpdateCodeReferences")
	}

	request.Header.Set("x-api-key", client.options.ApiKey)
	request.Header.Set("Content-Type", "application/json")

	_, err = client.httpClient.Do(request)
	if err != nil {
		return errors.Wrap(err, "Error doing the request UpdateCodeReferences")
	}

	return nil
}
