package helpers

import (
	"github.com/stretchr/testify/assert"
	. "github.com/whkelvin/stamp/pkg/helpers"
	"testing"
)

func TestGetYoutubeEmbedLinkShouldReturnYoutubeEmbedLink(t *testing.T) {
	videoId := "mZWsyUKwTbg"
	link1 := "http://www.youtube.com/watch?v=" + videoId + "&feature=feedrec_grec_index"
	link2 := "http://www.youtube.com/user/IngridMichaelsonVEVO#p/a/u/1/" + videoId
	link3 := "http://www.youtube.com/v/" + videoId + "?fs=1&amp;hl=en_US&amp;rel=0"
	link4 := "http://www.youtube.com/watch?v=" + videoId
	link5 := "http://www.youtube.com/embed/" + videoId + "?rel=0"
	link6 := "http://www.youtube.com/watch?v=" + videoId
	link7 := "http://youtu.be/" + videoId
	expected := "https://www.youtube.com/embed/" + videoId

	actual, err := GetYoutubeEmbedLink(link1)
	assert.Equal(t, err, nil)
	assert.Equal(t, expected, actual)

	actual, err = GetYoutubeEmbedLink(link2)
	assert.Equal(t, err, nil)
	assert.Equal(t, expected, actual)

	actual, err = GetYoutubeEmbedLink(link3)
	assert.Equal(t, err, nil)
	assert.Equal(t, expected, actual)

	actual, err = GetYoutubeEmbedLink(link4)
	assert.Equal(t, err, nil)
	assert.Equal(t, expected, actual)

	actual, err = GetYoutubeEmbedLink(link5)
	assert.Equal(t, err, nil)
	assert.Equal(t, expected, actual)

	actual, err = GetYoutubeEmbedLink(link6)
	assert.Equal(t, err, nil)
	assert.Equal(t, expected, actual)

	actual, err = GetYoutubeEmbedLink(link7)
	assert.Equal(t, err, nil)
	assert.Equal(t, expected, actual)
}

func TestGetYoutubeEmbedLinkShouldNotReturnYoutubeEmbedLink(t *testing.T) {
	link1 := "https://www.youtube.com/shorts/i5lM2QboT4U"
	link2 := "https://www.youtube.com/@sysprog"
	link3 := ""
	invalidVideoId := "mZWsyUKwTbg123455afldi"
	link4 := "http://www.youtube.com/watch?v=" + invalidVideoId + "&feature=feedrec_grec_index"
	invalidVideoId = "123"
	link5 := "http://www.youtube.com/watch?v=" + invalidVideoId + "&feature=feedrec_grec_index"

	_, err := GetYoutubeEmbedLink(link1)
	assert.NotEqual(t, err, nil)

	_, err = GetYoutubeEmbedLink(link2)
	assert.NotEqual(t, err, nil)

	_, err = GetYoutubeEmbedLink(link3)
	assert.NotEqual(t, err, nil)

	_, err = GetYoutubeEmbedLink(link4)
	assert.NotEqual(t, err, nil)

	_, err = GetYoutubeEmbedLink(link5)
	assert.NotEqual(t, err, nil)
}

func TestValidateGithubLinkShouldNotReturnError(t *testing.T) {
	link1 := "https://www.github.com/abcde/eolisj-eefj"
	link2 := "http://www.github.com/aoijf/seof-2te"

	err := ValidateGithubLink(link1)
	assert.Equal(t, err, nil)

	err = ValidateGithubLink(link2)
	assert.Equal(t, err, nil)

}

func TestValidateGithubLinkShouldReturnError(t *testing.T) {
	link1 := "https://www.github.com/abcde/eolisj-e/"
	link2 := "http://www.github.com/aoijf/"
	link3 := "http://www.github.com/"
	link4 := "http://www.github.com/helo/hello/hello"

	err := ValidateGithubLink(link1)
	assert.NotEqual(t, err, nil)

	err = ValidateGithubLink(link2)
	assert.NotEqual(t, err, nil)

	err = ValidateGithubLink(link3)
	assert.NotEqual(t, err, nil)

	err = ValidateGithubLink(link4)
	assert.NotEqual(t, err, nil)
}
