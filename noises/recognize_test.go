package noises

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestRecognize(t *testing.T) {
	data, err := ioutil.ReadFile("audio_test.mp3")
	if err != nil {
		t.Fatal(err)
	}

	b := bytes.NewReader(data)
	cB, err := Convert(b)
	if err != nil {
		t.Fatal(err)
	}

	// result channel
	c := make(chan string)
	// error channel
	e := make(chan error)

	go Recognize(cB, c, e)

	select {
	case err := <-e:
		close(c)
		close(e)
		t.Fatal(err.Error())
	case text := <-c:
		close(c)
		close(e)
		fmt.Println("Resp:", text)
		if text != "Tavşan ile kuşun macerası" {
			t.Fatal("Recognition is wrong.")
		}
	}
}
