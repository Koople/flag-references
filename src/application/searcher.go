package application

import (
	"github.com/koople/flag-references/src/api"
	"github.com/koople/flag-references/src/git"
	"github.com/koople/flag-references/src/searcher"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func Run(repository string, projectPath string, apiKey string, baseUri string, logger *logrus.Logger) error {
	options := api.KPLOptions{
		BaseUri: baseUri,
		ApiKey:  apiKey,
		Logger:  logger,
	}
	client := api.NewClient(options)

	flags, err := client.GetListFlags()
	if err != nil {
		return err
	}

	founds := make([]api.FlagFound, 0)

	for _, flag := range flags {
		flagFounds, err := searcher.FileSearcher(projectPath, flag, 5)
		if err != nil {
			err := errors.Wrap(err, "Error while searching")
			return err
		}

		founds = append(founds, api.FlagFound{
			Flag:   flag,
			Founds: flagFounds,
		})
	}

	gitClient, err := git.NewGitClient(projectPath)
	if err != nil {
		return err
	}

	branch, err := gitClient.CurrentBranch()
	if err != nil {
		return err
	}

	references := api.RepositoryReferences{
		Repository: repository,
		Branch:     branch,
		References: founds,
	}

	return client.SaveFlagsInformation(references)
}
