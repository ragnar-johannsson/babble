package babble

import (
	"io/ioutil"
	"testing"
)

var text1 string = "Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae (dicta sunt explicabo.) Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?"

var text2 string = "But I must explain to you how all this mistaken idea of denouncing pleasure and praising pain was born and I will give you a complete account of the system, and expound the actual teachings of the great explorer of the truth, the master-builder of human happiness. No one rejects, dislikes, or avoids pleasure itself, because it is pleasure, but because those who do not know how to pursue pleasure rationally encounter consequences that are extremely painful. Nor again is there anyone who loves or pursues or desires to obtain pain of itself, because it is pain, but because occasionally circumstances occur in which toil and pain can procure him some great pleasure. To take a trivial example, which of us ever undertakes laborious physical exercise, except to obtain some advantage from it? But who has any right to find fault with a man who chooses to enjoy a pleasure that has no annoying consequences, or one who avoids a pain that produces no resultant pleasure?"

func TestSplitIntoSentences(t *testing.T) {
	sentences := splitIntoSentences(text1)

	if len(sentences) != 5 {
		t.Fail()
	}
}

func TestSplitSentencesIntoWords(t *testing.T) {
	sentences := splitIntoSentences(text1)
	sentencesInWords := splitSentencesIntoWords(sentences)

	if len(sentencesInWords) != 5 {
		t.Fail()
	}

	if len(sentencesInWords[1]) != 24 {
		t.Fail()
	}

}

func TestNewText(t *testing.T) {
	m := NewTextModel(text2, DEFAULT_STATE_SIZE)

	if len(m.chain.model) != 161 {
		t.Fail()
	}
}

func TestMakeSentence(t *testing.T) {
	b, _ := ioutil.ReadFile("./texts/pg5827.txt")
	m := NewTextModel(string(b), DEFAULT_STATE_SIZE)
	s := m.MakeSentence()

	if s == "" {
		t.Fail()
	}
}

func TestMakeShortSentence(t *testing.T) {
	b, _ := ioutil.ReadFile("./texts/pg5827.txt")
	m := NewTextModel(string(b), DEFAULT_STATE_SIZE)
	s := m.MakeShortSentence(140)

	if s == "" || len(s) > 140 {
		t.Fail()
	}
}

func TestSentenceIsFeasible(t *testing.T) {
	s1 := "This 'entence is not good enough"
	s2 := "This sentence is not [either"
	s3 := "or ]this!"
	sS := []string{s1, s2, s3}

	for _, s := range sS {
		if sentenceIsFeasible(s) {
			t.Fail()
		}
	}
}

func TestTestSentence(t *testing.T) {
	text := NewTextModel(text2, DEFAULT_STATE_SIZE)

	if !text.testSentenceSimilarity("This is not a drill!") {
		t.Fail()
	}

	if text.testSentenceSimilarity("") {
		t.Fail()
	}

	if text.testSentenceSimilarity("To take a trivial example, which of us ever undertakes laborious physical exercise, except to obtain some advantage from it?") {
		t.Fail()
	}

}
