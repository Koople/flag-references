package flags_searcher

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func Test_find_flags_in_directory(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var projectPath = dir + "/example"

	founds, errorMessage := FileSearcher(projectPath, "someFlag", 2)

	if errorMessage != nil {
		t.Fail()
	}

	expected := []File{
		{
			Filename: projectPath + "/index.tsx",
			Chunks: []Chunk{
				{
					Chunk: []LineCode{
						{LineNumber: 3, Code: "// NotEnabled Component"},
						{LineNumber: 4, Code: "const NotEnabled = () => {"},
						{LineNumber: 5, Code: "  const flag1 = useIsEnabled(\"someFlag\");"},
						{LineNumber: 6, Code: "  const flag2 = useIsEnabled(\"anotherFlag\");"},
						{LineNumber: 7, Code: "  return flag1 ? <h1>flag1 enabled</h1> : flag2 ? <h1>flag2 enabled</h1> : <h1>all features disabled</h1>"},
					},
				},
				{
					Chunk: []LineCode{
						{LineNumber: 9, Code: ""},
						{LineNumber: 10, Code: "const App = () => {"},
						{LineNumber: 11, Code: "  const flag = useIsEnabled(\"someFlag\");"},
						{LineNumber: 12, Code: "  return <div>"},
						{LineNumber: 13, Code: "    {flag ? <h1>Hello</h1> : <NotEnabled/>}"},
					},
				},
			},
		},
	}

	assert.Equal(t, expected, founds)
}

func Test_non_existing_flags(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var projectPath = dir + "/example"

	founds, errorMessage := FileSearcher(projectPath, "nonExistingFlag", 3)

	if errorMessage != nil {
		t.Fail()
	}

	expected := make([]File, 0)
	assert.Equal(t, expected, founds)
}

func Test_non_existing_directory(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var projectPath = dir + "/nonExisting"

	_, errorCode := FileSearcher(projectPath, "someFlag", 3)

	assert.Equal(t, errorCode.Error(), "project path does not exist")
}

func Test_split_line_into_a_code(t *testing.T) {
	code := NewLineCode("5:   someCode := \"here\"")
	expected := LineCode{
		Code:       "   someCode := \"here\"",
		LineNumber: 5,
	}

	assert.Equal(t, expected, code)
}

func Test_grouping_into_different_chunks_based_on_linenumbers(t *testing.T) {
	firstBlock := []LineCode{
		{LineNumber: 3, Code: "// NotEnabled Component"},
		{LineNumber: 4, Code: "const NotEnabled = () => {"},
		{LineNumber: 5, Code: "  const flag1 = useIsEnabled(\"someFlag\");"},
	}

	secondBlock := []LineCode{
		{LineNumber: 10, Code: "const App = () => {"},
		{LineNumber: 11, Code: "  const flag = useIsEnabled(\"someFlag\");"},
		{LineNumber: 12, Code: "  return <div>"},
	}

	unorderedLineCodes := []LineCode{
		{LineNumber: 3, Code: "// NotEnabled Component"},
		{LineNumber: 12, Code: "  return <div>"},
		{LineNumber: 5, Code: "  const flag1 = useIsEnabled(\"someFlag\");"},
		{LineNumber: 10, Code: "const App = () => {"},
		{LineNumber: 4, Code: "const NotEnabled = () => {"},
		{LineNumber: 11, Code: "  const flag = useIsEnabled(\"someFlag\");"},
	}

	chunks := groupsByConsequtiveLines(unorderedLineCodes)

	expected := []Chunk{
		{Chunk: firstBlock},
		{Chunk: secondBlock},
	}

	assert.Equal(t, expected, chunks)
}
