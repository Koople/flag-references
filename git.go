package flags_searcher

import "github.com/go-git/go-git/v5"

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

	return branchName.Name().Short(), nil
}
