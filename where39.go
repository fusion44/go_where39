package where39

import (
	"errors"
	"math"
)

func ToWords(coords LatLng) ([]string, error) {
	lat := coords.Lat
	lng := math.Remainder((coords.Lng-180.0), 360.0) + 180.0

	lat += 90
	lng += 180

	finalWords := make([]string, 5)
	var latw = lat
	var lngw = lng
	for i := 0; i < 5; i++ {
		var tilesize = TileSizes[i]
		var seeds = GetTileSeeds()[i]
		var clatw = math.Floor(latw / tilesize)
		var clngw = math.Floor(lngw / tilesize)

		latw -= tilesize * clatw
		lngw -= tilesize * clngw
		finalWords[i] = seeds[int(clatw)][int(clngw)]
	}

	return finalWords, nil
}

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
