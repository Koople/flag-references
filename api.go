package flags_searcher

import (
	"bytes"
	"encoding/json"
	"fmt"
	rhttp "github.com/hashicorp/go-retryablehttp"
	"github.com/pkg/errors"
)

const (
	BaseUri = "https://sdk.koople.io/"
)

type KPLApi struct {
	httpClient *rhttp.Client
	options    KPLOptions
}

type KPLOptions struct {
	BaseUri string
	ApiKey  string
}

func NewClient(opts KPLOptions) KPLApi {
	if opts.BaseUri == "" {
		opts.BaseUri = BaseUri
	}

	httpClient := rhttp.NewClient()

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

func (client KPLApi) SaveFlagsInformation(founds []FlagFound) error {
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
