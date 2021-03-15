package flags_searcher

import (
	"github.com/go-git/go-git/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

const directory = "testProject"

func createRepository(t *testing.T) {
	require.NoError(t, os.RemoveAll(directory))
	_, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL: "https://github.com/Koople/test.git",
	})
	require.NoError(t, err)
}

func TestGetCurrentBranchName(t *testing.T) {
	createRepository(t)
	gitClient, err := NewGitClient(directory)
	require.NoError(t, err)

	branch, err := gitClient.CurrentBranch()
	require.NoError(t, err)

	assert.Equal(t, branch, "main")
}
