package transcript

import (
	"audiofile/models"
	"context"
	"fmt"
	"os"

	aai "github.com/AssemblyAI/assemblyai-go-sdk"
	"github.com/joho/godotenv"
)

/*
we have to get API fromAssemblyAI,
a third-party API, and requires an API key, which can be set locally to **ASSEMBLY_API_KEY**.

go get github.com/AssemblyAI/assemblyai-go-sdk

https://www.assemblyai.com/blog/golang-speech-recognition/
link to get free API key -> create account  and you will get api key

*/
// -----------------------------------------------------------------------------------------------------------
/*

go get github.com/joho/godotenv
 I want to store API key in .env file and will have to use godotenv package to read the .env file.
 as in python also put .env in gitignore file and for sameple I will create sameple_env file.
*/

// -----------------------------------------------------------------------------------------------------------

// https://www.assemblyai.com/docs/getting-started/transcribe-an-audio-file #helpful docs
func Extract(m *models.Audio) error {

	err := godotenv.Load(".env")	
	
	if err != nil {
		fmt.Println("Error loading .env file", err)
		return nil
	}

	// get the API key from the environment variable
	apiKey := os.Getenv("ASSEMBLY_API_KEY")
	// fmt.Println("My api key---------------",apiKey)
	// load the audio file

	fmt.Println("Sending the audio file to AssemblyAI for transcription")
	data, err := os.Open(m.Path)

	if err != nil {
		return err
	}	


	ctx := context.Background()

    client := aai.NewClient(apiKey )

    params := &aai.TranscriptOptionalParams{
        SpeakerLabels: aai.Bool(true),
    }


    transcript, err := client.Transcripts.TranscribeFromReader(ctx, data, params)



	if err != nil {
        fmt.Println("Something bad happened:", err)
        return err
    }


	m.Metadata.Transcript = *transcript.Text

	fmt.Println("transcript: ", m.Metadata.Transcript)

	return nil

}