package go_where39

import (
	"fmt"
	"math"
)

func ToWords(coords LatLng, shuffleValue int) ([]string, error) {
	if shuffleValue < 1 || shuffleValue > 9999999 {
		return nil, fmt.Errorf("shuffleValue must be >= 1 and <= 9999999")
	}

	lat := coords.Lat
	lng := math.Remainder((coords.Lng-180.0), 360.0) + 180.0

	lat += 90
	lng += 180

	finalWords := make([]string, 5)
	latw := lat
	lngw := lng

	seeds, err := GetTileSeeds(shuffleValue)
	if err != nil {
		return nil, fmt.Errorf("error getting tile seeds %s", err)
	}

	for i := 0; i < 5; i++ {
		tilesize := TileSizes[i]
		clatw := math.Floor(latw / tilesize)
		clngw := math.Floor(lngw / tilesize)
		latw -= tilesize * clatw
		lngw -= tilesize * clngw
		finalWords[i] = seeds[i][int(clatw)][int(clngw)]
	}

	return finalWords, nil
}

func FromWords(words []string, shuffleValue int) (LatLng, error) {
	if len(words) < 1 || len(words) > 5 {
		return LatLng{}, fmt.Errorf("word list must contain between 1 and 5 words")
	}
	if shuffleValue < 1 || shuffleValue > 9999999 {
		return LatLng{}, fmt.Errorf("shuffleValue must be >= 1 and <= 9999999")
	}

	var la float64
	var lo float64

	tileSeeds, err := GetTileSeeds(shuffleValue)

	if err != nil {
		return LatLng{}, fmt.Errorf("unable to get tile seeds with shuffleValue %v", shuffleValue)
	}

	for i, word := range words {
		p, err := getIndexOfWord(tileSeeds[i], word)

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

	return LatLng{}, fmt.Errorf("word not found: %s", word)
}
