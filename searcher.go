package flags_searcher

import (
	"github.com/pkg/errors"
)

type FlagFound struct {
	Flag   string `json:"flag"`
	Founds []File `json:"founds"`
}

type RepositoryReferences struct {
	Repository string      `json:"repository"`
	Branch     string      `json:"branch"`
	References []FlagFound `json:"references"`
}

func Run(repository string, projectPath string, apiKey string) error {
	options := KPLOptions{ApiKey: apiKey}
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

	gitClient, err := NewGitClient(projectPath)
	if err != nil {
		return err
	}

	branch, err := gitClient.CurrentBranch()
	if err != nil {
		return err
	}

	references := RepositoryReferences{
		Repository: repository,
		Branch:     branch,
		References: founds,
	}

	return client.SaveFlagsInformation(references)
}
