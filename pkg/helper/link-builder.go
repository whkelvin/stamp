package helper

import (
	"errors"
	"regexp"
)

// http://www.youtube.com/watch?v={id here}&feature=feedrec_grec_index
// http://www.youtube.com/user/IngridMichaelsonVEVO#p/a/u/1/{id_here}
// http://www.youtube.com/v/{id_here}?fs=1&amp;hl=en_US&amp;rel=0
// http://www.youtube.com/watch?v={id here}
// http://www.youtube.com/embed/{id here}?rel=0
// http://www.youtube.com/watch?v={id here}
// http://youtu.be/{id here}
func GetYoutubeEmbedLink(link string) (string, error) {
	regExp := `/^.*(youtu\.be\/|v\/|u\/\w\/|embed\/|watch\?v=|\&v=)([^#\&\?]*).*/`
	exp := regexp.MustCompile(regExp)
	submatches := exp.FindStringSubmatch(link)

	if len(submatches) >= 2 && len(submatches[2]) == 11 {
		id := submatches[2]
		return "https://www.youtube.com/embed/" + id, nil
	}
	return "", errors.New("Invalid youtube link")
}
