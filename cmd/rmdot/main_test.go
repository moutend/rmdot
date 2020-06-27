package main

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCandidatePaths(t *testing.T) {
	t.Parallel()

	expected := []string{
		filepath.Join("testdata", ".bar"),
		filepath.Join("testdata", ".foo"),
		filepath.Join("testdata", ".foobar"),
	}

	actual, err := getCandidatePaths("xxxxxxxx")

	require.Error(t, err)
	require.Empty(t, actual)

	actual, err = getCandidatePaths("testdata")

	require.NoError(t, err)
	require.Equal(t, len(expected), len(actual))

	for i, _ := range expected {
		require.Equal(t, expected[i], actual[i])
	}
}

func TestNormalizeCandidatePaths(t *testing.T) {
	t.Parallel()

	require.Equal(t, normalizeCandidatePath(""), "")
	require.Equal(t, normalizeCandidatePath(filepath.Join("path", "to", ".file")), filepath.Join("path", "to", ".file"))
	require.Equal(t, normalizeCandidatePath(filepath.Join("path", "to", ".dir", "file")), filepath.Join("path", "to", ".dir"))
	require.Equal(t, normalizeCandidatePath(filepath.Join("path", "to", ".dir1", "dir2", ".dir3")), filepath.Join("path", "to", ".dir1"))
}
