package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

func createRepository(t *testing.T, detached bool) string {
	var directory = "fsearchertmpdir"
	dir, err := ioutil.TempDir(".", directory)

	require.NoError(t, os.RemoveAll(dir))
	opts := &git.CloneOptions{
		URL: "https://github.com/Koople/flag-references.git",
	}

	if detached {
		opts.ReferenceName = "refs/tags/0.0.2"
	}

	_, err = git.PlainClone(dir, false, opts)

	require.NoError(t, err)
	return dir
}

func TestGetCurrentBranchName(t *testing.T) {
	directory := createRepository(t, false)
	gitClient, err := NewGitClient(directory)
	require.NoError(t, err)

	branch, err := gitClient.CurrentBranch()
	require.NoError(t, err)

	assert.Equal(t, branch, "master")
	require.NoError(t, os.RemoveAll(directory))
}

func TestGetCurrentBranchNameWhenDetachedHead(t *testing.T) {
	directory := createRepository(t, true)
	gitClient, err := NewGitClient(directory)
	require.NoError(t, err)

	branch, err := gitClient.CurrentBranch()
	require.Empty(t, branch)
	assert.Error(t, err, "Repository in detached HEAD. Use --branch option to set the branch name.")
}
