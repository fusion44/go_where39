package where39

import (
	"errors"
)

func FromWords(words []string) (LatLng, error) {
	if len(words) < 1 || len(words) > 5 {
		return LatLng{}, errors.New("Word list must contain between 1 and 5 words")
	}

	var la float64
	var lo float64

	for i, word := range words {
		p, err := getIndexOfWord(GetTileSeeds()[i], word)

		if err != nil {
			return LatLng{}, err
		}

		la += p.Lat * TileSizes[i]
		lo += p.Lng * TileSizes[i]
	}

	la -= 90
	lo -= 180

	return LatLng{Lat: la, Lng: lo}, nil
}

func getIndexOfWord(words [][]string, word string) (LatLng, error) {
	lngIndex := -1
	for latIndex := 0; latIndex < len(words); latIndex++ {
		for j, value := range words[latIndex] {
			if value == word {
				lngIndex = j
				break
			}
		}

		if lngIndex > 0 {
			return LatLng{Lat: float64(latIndex), Lng: float64(lngIndex)}, nil
		}
	}

	return LatLng{}, errors.New("Word not found")
}
