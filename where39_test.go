package where39

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

const shuffleValue = 1337

// Room77, Berlin
var room77Coords = LatLng{Lat: 52.49303704, Lng: 13.41792593}
var emptyWordList = []string{}
var tooLongWordList = []string{"slush", "extend", "until", "layer", "arch", "small"}
var room77Words = []string{"slush", "extend", "until", "layer", "arch"}
var room77Shuffled1337 = []string{"small", "extra", "unusual", "lazy", "accident"}

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
	coords, err := FromWords(room77Words)
	if err != nil {
		t.Log("Should not return an error: ", err)
		t.Fail()
	}

	if fmt.Sprintf("%.8f", coords.Lat) != fmt.Sprintf("%.8f", room77Coords.Lat) || fmt.Sprintf("%.8f", coords.Lng) != fmt.Sprintf("%.8f", room77Coords.Lng) {
		t.Log("Conversion results should be ", room77Coords)
		t.Log("Conversion results should are ", coords)
		t.Fail()
	}
}

func TestToWordsConversion(t *testing.T) {
	words, err := ToWords(room77Coords)

	if err != nil {
		t.Log("Error running ToWords() ", err)
		t.Fail()

	}

	if !reflect.DeepEqual(words, room77Words) {
		t.Log("Conversion results should be ", room77Coords)
		t.Log("Conversion results should are ", words)
		t.Fail()
	}

}
