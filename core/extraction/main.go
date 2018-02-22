package main

import (
	"bufio"
	"github.com/yut-kt/high-q-TTS/env"
	"github.com/yut-kt/high-q-TTS/peripheral/file"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func fCreateOpen(path string) *os.File {
	fOut, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return fOut
}

func fClose(fp *os.File) {
	if err := fp.Close(); err != nil {
		panic(err)
	}
}

func concatExceptAppearSentence(scanner *bufio.Scanner, reg string) string {
	var sentencesStr string
	for scanner.Scan() {
		text := scanner.Text()
		if !regexp.MustCompile(reg).Match([]byte(scanner.Text())) {
			sentencesStr += text
		}
	}
	return sentencesStr
}

func extractSentences(finp *os.File) []string {
	scanner := bufio.NewScanner(transform.NewReader(finp, japanese.ShiftJIS.NewDecoder()))
	sentencesStr := concatExceptAppearSentence(scanner, `[：■]`)
	sentences := strings.Split(sentencesStr, "。")
	return sentences[:len(sentences)-1]
}

func mapStr(strs []string, f func(string) string) []string {
	r := make([]string, len(strs))
	for i, e := range strs {
		r[i] = f(e)
	}
	return r
}

func indexingSentences(sentences []string, index *int) []string {
	return mapStr(sentences, func(s string) string {
		*index++
		return strconv.Itoa(*index) + " " + s
	})
}

type sentenceMap map[string]int

func (m sentenceMap) register(sentences []string) {
	for _, sentence := range sentences {
		m[sentence] = 1
	}
}

func indexingSentenceMap(m sentenceMap) sentenceMap {
	index := 0
	newMap := make(sentenceMap, len(m))
	for sentence := range m {
		index++
		newMap[strconv.Itoa(index)+" "+sentence] = 0
		delete(m, sentence)
	}
	return newMap
}

func writeStrings(foutp *os.File, strs []string) {
	for _, str := range strs {
		_, err := foutp.WriteString(str + "\n")
		if err != nil {
			panic(err)
		}
	}
}

func writeSentenceMap(foutp *os.File, m sentenceMap) {
	for sentence := range m {
		_, err := foutp.WriteString(sentence + "\n")
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	fOutBody := fCreateOpen(env.AllBodyPath)
	fOutIndexedBody := fCreateOpen(env.AllIndexedBodyPath)
	fOutUniqueBody := fCreateOpen(env.AllUniqueBodyPath)
	fOutIndexedUniqueBody := fCreateOpen(env.AllIndexedUniqueBodyPath)
	defer func() {
		fClose(fOutBody)
		fClose(fOutIndexedBody)
		fClose(fOutUniqueBody)
		fClose(fOutIndexedUniqueBody)
	}()

	filePaths, err := file.FetchPaths(env.OrigDataDir)
	if err != nil {
		panic(err)
	}

	uniqueSentences := make(sentenceMap, 100)
	index := 0
	for _, filePath := range filePaths {
		finp, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}

		sentences := extractSentences(finp)
		writeStrings(fOutBody, sentences)
		indexedSentences := indexingSentences(sentences, &index)
		writeStrings(fOutIndexedBody, indexedSentences)

		uniqueSentences.register(sentences)

		if err := finp.Close(); err != nil {
			panic(err)
		}
	}

	writeSentenceMap(fOutUniqueBody, uniqueSentences)
	writeSentenceMap(fOutIndexedUniqueBody, indexingSentenceMap(uniqueSentences))

}
