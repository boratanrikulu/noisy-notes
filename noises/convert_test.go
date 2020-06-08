package noises

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestConvert(t *testing.T) {
	data, err := ioutil.ReadFile("audio_test.mp3")
	if err != nil {
		t.Fatal(err)
	}

	b := bytes.NewReader(data)
	_, err = Convert(b)
	if err != nil {
		t.Fatal(err)
	}
}
