package helpers

import (
	. "github.com/franela/goblin"
	. "github.com/whkelvin/stamp/pkg/helpers"
	"testing"
)

func TestHelpers(t *testing.T) {

	g := Goblin(t)

	g.Describe("GetYoutubeEmbedLink", func() {

		videoId := "mZWsyUKwTbg"
		link1 := "http://www.youtube.com/watch?v=" + videoId + "&feature=feedrec_grec_index"
		link2 := "http://www.youtube.com/user/IngridMichaelsonVEVO#p/a/u/1/" + videoId
		link3 := "http://www.youtube.com/v/" + videoId + "?fs=1&amp;hl=en_US&amp;rel=0"
		link4 := "http://www.youtube.com/watch?v=" + videoId
		link5 := "http://www.youtube.com/embed/" + videoId + "?rel=0"
		link6 := "http://www.youtube.com/watch?v=" + videoId
		link7 := "http://youtu.be/" + videoId
		expected := "https://www.youtube.com/embed/" + videoId

		GetYoutubeEmbedLink(link1)
		g.It("should return youtube embed link", func() {
			actual, err := GetYoutubeEmbedLink(link1)
			g.Assert(err).Equal(nil)
			g.Assert(actual).Equal(expected)

			actual, err = GetYoutubeEmbedLink(link2)
			g.Assert(err).Equal(nil)
			g.Assert(actual).Equal(expected)

			actual, err = GetYoutubeEmbedLink(link3)
			g.Assert(err).Equal(nil)
			g.Assert(actual).Equal(expected)

			actual, err = GetYoutubeEmbedLink(link4)
			g.Assert(err).Equal(nil)
			g.Assert(actual).Equal(expected)

			actual, err = GetYoutubeEmbedLink(link5)
			g.Assert(err).Equal(nil)
			g.Assert(actual).Equal(expected)

			actual, err = GetYoutubeEmbedLink(link6)
			g.Assert(err).Equal(nil)
			g.Assert(actual).Equal(expected)

			actual, err = GetYoutubeEmbedLink(link7)
			g.Assert(err).Equal(nil)
			g.Assert(actual).Equal(expected)
		})
	})
}
