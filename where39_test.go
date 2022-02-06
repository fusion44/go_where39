package where39

import (
	"fmt"
	"strings"
	"testing"
)

const shuffleValue = 1337

// (ex)Room77, Berlin
var okCoords = LatLng{Lat: 52.49303704, Lng: 13.41792593}

var emptyWordList = []string{}
var tooLongWordList = []string{"slush", "extend", "until", "layer", "arch", "small"}
var wordsStandard = []string{"slush", "extend", "until", "layer", "arch"}
var wordsShuffled1337 = []string{"small", "extra", "unusual", "lazy", "accident"}

func TestFromWordsArgs(t *testing.T) {
	// test an empty word list
	_, err := FromWords(emptyWordList)
	if !strings.HasPrefix(err.Error(), "Word list must contain between") {
		t.Log("Should return error when passing an empty array")
		t.Fail()
	}

	// test a to long word list
	_, err = FromWords(tooLongWordList)
	if !strings.HasPrefix(err.Error(), "Word list must contain between") {
		t.Log("Should return error when passing an empty array")
		t.Fail()
	}
}

func TestFromWordsConversion(t *testing.T) {
	coords, err := FromWords(wordsStandard)
	if err != nil {
		t.Log("Should not return an error: ", err)
		t.Fail()
	}

	if fmt.Sprintf("%.8f", coords.Lat) != fmt.Sprintf("%.8f", okCoords.Lat) || fmt.Sprintf("%.8f", coords.Lng) != fmt.Sprintf("%.8f", okCoords.Lng) {
		t.Log("Conversion results should be ", okCoords)
		t.Log("Conversion results should are ", coords)
		t.Fail()
	}
}
