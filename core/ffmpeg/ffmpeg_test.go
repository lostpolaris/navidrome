package ffmpeg

import (
	"testing"

	"github.com/navidrome/navidrome/log"
	"github.com/navidrome/navidrome/tests"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFFmpeg(t *testing.T) {
	tests.Init(t, false)
	log.SetLevel(log.LevelFatal)
	RegisterFailHandler(Fail)
	RunSpecs(t, "FFmpeg Suite")
}

var _ = Describe("ffmpeg", func() {
	BeforeEach(func() {
		_, _ = ffmpegCmd()
		ffmpegPath = "ffmpeg"
		ffmpegErr = nil
	})
	Describe("createFFmpegCommand", func() {
		It("creates a valid command line", func() {
			args := createFFmpegCommand("ffmpeg -i %s -b:a %bk mp3 -", "/music library/file.mp3", 123)
			Expect(args).To(Equal([]string{"ffmpeg", "-i", "/music library/file.mp3", "-b:a", "123k", "mp3", "-"}))
		})
	})

	Describe("createProbeCommand", func() {
		It("creates a valid command line", func() {
			args := createProbeCommand(probeCmd, []string{"/music library/one.mp3", "/music library/two.mp3"})
			Expect(args).To(Equal([]string{"ffmpeg", "-i", "/music library/one.mp3", "-i", "/music library/two.mp3", "-f", "ffmetadata"}))
		})
	})
})
