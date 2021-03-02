package flags_searcher

import (
	"github.com/pkg/errors"
	"os"
)

type FlagFound struct {
	Flag   string `json:"flag"`
	Founds []File `json:"founds"`
}

func Run(projectPath string, apiKey string) error {
	env, exists := os.LookupEnv("API_URL")
	baseUri := "https://sdk.koople.io"
	if exists {
		baseUri = env
	}

	options := KPLOptions{BaseUri: baseUri, ApiKey: apiKey}
	client := NewClient(options)

	flags, err := client.GetListFlags()
	if err != nil {
		return err
	}

	founds := make([]FlagFound, 0)

	for _, flag := range flags {
		flagFounds, err := FileSearcher(projectPath, flag, 5)
		if err != nil {
			err := errors.Wrap(err, "Error while searching")
			return err
		}

		founds = append(founds, FlagFound{
			Flag:   flag,
			Founds: flagFounds,
		})
	}

	return client.SaveFlagsInformation(founds)
}
