package metadata

import (
	"audiofile/extractors/tags"
	"audiofile/extractors/transcript"
	"audiofile/models"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func (m *MetadataService) uploadHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("req:", req)

	file, handler, err := req.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer file.Close()

	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("Error Opening the File", err)
		res.WriteHeader(http.StatusInternalServerError)
		return

	}

	defer func(){
		os.Remove(handler.Filename)

		if err != nil {
			fmt.Println("Error Deleting the File", err)
			res.WriteHeader(http.StatusInternalServerError)
			
		}
		f.Close()
	}()

	buf := bytes.NewBuffer(nil)

	if io.Copy(buf, file); err != nil {
		fmt.Println("Error Copying the File to buffer", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	id, audioFilePath, err := m.Storage.Upload(buf.Bytes(), handler.Filename)

	audio := &models.Audio{
		Id: id,
		Path: audioFilePath,
	}

	err = m.Storage.SaveMetadata(audio)
	
	if err != nil {
		fmt.Println("Error Saving Metadata", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	audio.Status = "Initialized"

	io.WriteString(res, audio.Id)

	go func(){
		var errors []string

		audio.Status = "In-progress"

		// Extract the tags from the audio file

		err := tags.Extract(audio)
		fmt.Println("Extracting Tags I had send the data to the function")

		if err != nil {
			errors = append(errors, err.Error())
			fmt.Println("Error Extracting Tags", err)

		}
		// save the metadata
		err = m.Storage.SaveMetadata(audio)

		if err != nil {
			errors = append(errors, err.Error())
			fmt.Println("Error Saving Metadata", err)
		}

		// extract the transcript from the audio file

		err = transcript.Extract(audio)
		fmt.Println("Extracting Transcript I had send the data to the function")
		if err != nil {
			errors = append(errors, err.Error())
			fmt.Println("Error Extracting Transcript", err)
		}

		audio.Error = errors
		audio.Status = "Complete"
		err = m.Storage.SaveMetadata(audio)
		if err != nil {
			fmt.Println("error saving metadata: ", err)
			errors = append(errors, err.Error())
		}

		if len(errors) > 0 {
			fmt.Println("errors occurred extracting metadata: ")
			for i := 0; i < len(errors); i++ {
				fmt.Printf("\terror[%d]: %s\n", i, errors[i])
			}
		} else {
			fmt.Println("successfully extracted and saved audio metadata: ", audio)

			err = m.Storage.PushToMongoDB(audio)
			if err != nil {
				fmt.Println("Error pushing audio to MongoDB Atlas, saved locally:", err)
				res.WriteHeader(http.StatusInternalServerError)
				return
			}

			fmt.Println("pushing audio metadata to MongoDB Atlas", err)
		}



	}()



}