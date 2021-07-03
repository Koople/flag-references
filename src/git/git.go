package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/pkg/errors"
)

type GitClient struct {
	repository *git.Repository
}

func NewGitClient(directory string) (*GitClient, error) {
	repository, err := git.PlainOpen(directory)
	if err != nil {
		return nil, err
	}

	return &GitClient{
		repository: repository,
	}, nil
}

func (client *GitClient) CurrentBranch() (string, error) {
	branchName, err := client.repository.Head()
	if err != nil {
		return "", err
	}

	name := branchName.Name().Short()
	if name == plumbing.HEAD.Short() {
		return "", errors.New("Repository in detached HEAD. Use --branch option to set the branch name.")
	}
	return name, nil
}
