package helpers

import (
	"errors"
	"fmt"
	"regexp"
)

func GetYoutubeEmbedLink(link string) (string, error) {
	regExp := `^.*(?:(?:youtu\.be\/|v\/|vi\/|u\/\w\/|embed\/)|(?:(?:watch)?\?v(?:i)?=|\&v(?:i)?=))([^#\&\?]*).*`
	exp := regexp.MustCompile(regExp)
	submatches := exp.FindStringSubmatch(link)

	if len(submatches) >= 1 && len(submatches[1]) == 11 {
		id := submatches[1]
		return "https://www.youtube.com/embed/" + id, nil
	}
	return "", errors.New("Invalid youtube link")
}

func ValidateGithubLink(link string) error {
	regExp := `^.*github\.com(?:\/[^\/]+){2}$`
	exp := regexp.MustCompile(regExp)
	submatches := exp.FindString(link)
	fmt.Printf("%v", submatches)
	if len(submatches) == 0 {
		return errors.New("Invalid github link")
	}
	return nil
}
