package babble

import (
	"math"
	"regexp"
	"strings"

	"gopkg.in/neurosnap/sentences.v1"
	"gopkg.in/neurosnap/sentences.v1/english"
)

var sentenceDetector *sentences.DefaultSentenceTokenizer
var sentenceOk *regexp.Regexp

func init() {
	s, err := english.NewSentenceTokenizer(nil)
	if err != nil {
		panic(err)
	}

	r, err := regexp.Compile("(^')|('$)|\\s'|'\\s|[\"(\\(\\)\\[\\])]")
	if err != nil {
		panic(err)
	}

	sentenceDetector = s
	sentenceOk = r
}

type TextModel struct {
	chain Chain
	text  string

	Tries           int
	MaxOverlapTotal int
	MaxOverlapRatio float64
}

func (t *TextModel) testSentenceSimilarity(s string) bool {
	if s == "" {
		return false
	}

	words := strings.Split(s, " ")
	overlapFromRatio := t.MaxOverlapRatio * float64(len(words))
	maxOverlap := int(math.Min(float64(t.MaxOverlapTotal), overlapFromRatio))
	gramCount := int(math.Max(float64(len(words)-maxOverlap), 1.0))

	for i := 0; i < gramCount; i++ {
		gram := strings.Join(words[i:i+maxOverlap+1], " ")

		if strings.Contains(t.text, gram) {
			return false
		}
	}

	return true
}

func (t *TextModel) MakeSentence() string {
	for x := 0; x < t.Tries; x++ {
		s := t.chain.Generate()

		if t.testSentenceSimilarity(s) {
			return s
		}
	}

	return ""
}

func (t *TextModel) MakeShortSentence(lengthLimit int) string {
	for x := 0; x < t.Tries; x++ {
		s := t.chain.Generate()

		if len(s) < lengthLimit {
			if t.testSentenceSimilarity(s) {
				return s
			}
		}
	}

	return ""
}

func NewTextModel(inputText string, stateSize int) TextModel {
	nl := regexp.MustCompile("[\r\n]+")
	sp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	inputText = nl.ReplaceAllString(inputText, " ")
	inputText = sp.ReplaceAllString(inputText, " ")

	s := splitIntoSentences(inputText)
	ws := splitSentencesIntoWords(s)

	return TextModel{
		chain: NewChain(ws, stateSize),
		text:  inputText,

		Tries:           DEFAULT_SENTENCE_TRIES,
		MaxOverlapRatio: DEFAULT_MAX_OVERLAP_RATIO,
		MaxOverlapTotal: DEFAULT_MAX_OVERLAP_TOTAL,
	}
}

func splitIntoSentences(input string) []string {
	pos := sentenceDetector.SentencePositions(input)
	s := []string{}

	for i, v := range pos {
		if i == 0 {
			s = append(s, input[0:v])
			continue
		}

		if sentenceIsFeasible(input[pos[i-1]:v]) {
			s = append(s, string(input[pos[i-1]:v]))
		}
	}

	return s
}

func splitSentencesIntoWords(s []string) [][]string {
	ws := [][]string{}

	for _, v := range s {
		ws = append(ws, strings.Split(v, " "))
	}

	return ws
}

func sentenceIsFeasible(s string) bool {
	return !sentenceOk.Match([]byte(s))
}
