package noises

import (
	"context"
	"fmt"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

// Recognize returns text result from speech using speech-to-text api.
func Recognize(data []byte) (string, error) {
	ctx := context.Background()

	client, err := speech.NewClient(ctx)
	if err != nil {
		return "", err
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
		return "", fmt.Errorf("We could not take text from the speech.")
	}

	// Return the first result.
	transcript := resp.Results[0].Alternatives[0].Transcript
	return transcript, nil
}
