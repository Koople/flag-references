package flags_searcher

import (
	"bytes"
	"github.com/monochromegane/the_platinum_searcher"
	"github.com/pkg/errors"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type LineCode struct {
	Code       string `json:"code"`
	LineNumber int    `json:"lineNumber"`
}

type Chunk struct {
	Chunk LineCodes `json:"chunk"`
}

type File struct {
	Filename string  `json:"filename"`
	Chunks   []Chunk `json:"chunks"`
}

func FileSearcher(projectPath string, flag string, lineContext int) ([]File, error) {
	if lineContext <= 0 {
		lineContext = 5
	}

	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		return nil, errors.New("project path does not exist")
	}

	buf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)
	pt := the_platinum_searcher.PlatinumSearcher{Out: buf, Err: errBuf}
	pattern := "[\\\"\\'\\`]" + flag + "[\\\"\\'\\`]"
	exitCode := pt.Run([]string{"--nocolor", "--nogroup", "-e", "-C" + strconv.Itoa(lineContext), pattern, projectPath})
	result := buf.String()
	err := errBuf.String()
	if err != "" || exitCode == 1 {
		return nil, errors.New(err)
	}
	results := strings.Split(result, "\n")
	founds := make([]File, 0)

	fm := make(map[string]LineCodes)

	for _, line := range results {
		if line == "" {
			continue
		}

		splitted := strings.SplitN(line, ":", 2)
		fm[splitted[0]] = append(fm[splitted[0]], NewLineCode(splitted[1]))
	}

	for filename, linecodes := range fm {
		chunks := groupsByConsequtiveLines(linecodes)

		founds = append(founds, File{
			Filename: filename,
			Chunks:   chunks,
		})
	}

	return founds, nil
}

func NewLineCode(line string) LineCode {
	splitted := regexp.MustCompile("[\\:\\-]+").Split(line, 2)
	lineNumber, err := strconv.Atoi(splitted[0])
	if err != nil {
		return LineCode{
			Code:       "",
			LineNumber: -1,
		}
	}

	return LineCode{
		Code:       splitted[1],
		LineNumber: lineNumber,
	}
}

type LineCodes []LineCode

func (a LineCodes) Len() int           { return len(a) }
func (a LineCodes) Less(i, j int) bool { return a[i].LineNumber < a[j].LineNumber }
func (a LineCodes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func groupsByConsequtiveLines(lineCodes LineCodes) []Chunk {
	currentGroup := 0
	chunks := make(map[int]Chunk, 1)

	sort.Sort(lineCodes)

	currentLine := lineCodes[0].LineNumber

	for _, lineCode := range lineCodes {
		if lineCode.LineNumber == -1 {
			continue
		}

		if lineCode.LineNumber != currentLine {
			currentGroup++
			chunks[currentGroup] = Chunk{Chunk: []LineCode{}}
		}

		chunks[currentGroup] = Chunk{
			Chunk: append(chunks[currentGroup].Chunk, lineCode),
		}

		currentLine = lineCode.LineNumber + 1
	}

	var keys []Chunk
	for _, value := range chunks {
		keys = append(keys, value)
	}

	return keys
}
