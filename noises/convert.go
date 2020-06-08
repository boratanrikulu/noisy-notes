package noises

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

// Convert noises to wav format by using ffmpeg.
func Convert(file io.Reader) ([]byte, error) {
	output := &bytes.Buffer{}

	ffmpeg := exec.Command("ffmpeg",
		"-i", "/dev/stdin", // input from stdin
		"-ac", "1", // mono channel
		"-ar", "16000", // sample rate hertz
		"-f", "wav", // format
		"-") // output to stdout
	ffmpeg.Stdin = file
	ffmpeg.Stdout = output

	err := ffmpeg.Run()
	if err != nil {
		return nil, fmt.Errorf("There is an error with the noise file.\n We could not convert it to .wav")
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(output)

	return buf.Bytes(), nil
}
