package notes

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestRecognize(t *testing.T) {
	data, err := ioutil.ReadFile("audio_test.mp3")
	if err != nil {
		t.Fatal(err)
	}

	resp, err := Recognize(data)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Resp:", resp)
	if resp != "Tavşan ile kuşun macerası" {
		t.Fatal("Ses eşleşmedi.")
	}
}
