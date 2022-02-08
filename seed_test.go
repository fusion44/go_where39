package go_where39

import (
	"reflect"
	"testing"
)

func TestShuffleWordList(t *testing.T) {
	wordList, err := shuffledWordList(shuffleValue)

	if err != nil {
		t.Log("Error running shuffledWordList() ", err)
		t.Fail()

	}

	if !reflect.DeepEqual(wordList, fullWordListShuffled1337) {
		t.Log("Shuffled word list doesn't match")
		t.Fail()
	}
}
