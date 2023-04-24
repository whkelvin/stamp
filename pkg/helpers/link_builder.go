package helpers

import (
	"errors"
	"regexp"
)

func GetYoutubeEmbedLink(link string) (string, error) {
	regExp := `^.*(?:(?:youtu\.be\/|v\/|vi\/|u\/\w\/|embed\/|shorts\/)|(?:(?:watch)?\?v(?:i)?=|\&v(?:i)?=))([^#\&\?]*).*`
	exp := regexp.MustCompile(regExp)
	submatches := exp.FindStringSubmatch(link)

	if len(submatches) >= 1 && len(submatches[1]) == 11 {
		id := submatches[1]
		return "https://www.youtube.com/embed/" + id, nil
	}
	return "", errors.New("Invalid youtube link")
}
