package noises

import (
	"context"
	"fmt"
	"os"

	speech "cloud.google.com/go/speech/apiv1"
	"google.golang.org/api/option"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

// Recognize returns text result from speech using speech-to-text api.
// Send the result to the channel.
// c : result channel
// e : error channel
func Recognize(data []byte, c chan<- string, e chan<- error) {
	ctx := context.Background()

	client, err := speech.NewClient(ctx, option.WithCredentialsJSON([]byte(
		os.Getenv("GOOGLE_APPLICATION_CREDENTIALS_JSON"))))
	if err != nil {
		e <- err
		return
	}

	// Send the contents of the audio file with the encoding and
	// and sample rate information to be transcripted.
	resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:                   speechpb.RecognitionConfig_ENCODING_UNSPECIFIED,
			SampleRateHertz:            16000,
			LanguageCode:               "tr-Tr",
			EnableAutomaticPunctuation: true,
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
		},
	})

	// Return an error message if the result is nil.
	if resp == nil || len(resp.Results) == 0 || len(resp.Results[0].Alternatives) == 0 {
		e <- fmt.Errorf("We could not take text from the speech.")
		return
	}

	transcript := ""
	// Return the first result.
	for _, result := range resp.Results {
		transcript += result.Alternatives[0].Transcript
	}

	c <- transcript
}
