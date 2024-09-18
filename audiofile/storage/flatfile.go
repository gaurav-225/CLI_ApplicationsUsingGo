package storage

import (
	"audiofile/models"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	
)

type FlatFile struct {
	Name string
}

// 1. Listing audio metadata in storage
// 2. Searching audio metadata in storage
// 3. Deleting audio from storage


func (f FlatFile) Delete(id string) error {
	fmt.Println("Deleting")
	return nil
}

func (f FlatFile) GetByID(id string) (*models.Audio, error) {
	dirname, err := os.UserHomeDir()

	if err != nil {
		return nil, err
	}

	mettadataFilePath := filepath.Join(dirname, "audiofile", id, "metadata.json")




	file, err := os.ReadFile(mettadataFilePath)

	if err != nil {
		return nil, err
	}

	data := models.Audio{}

	err = json.Unmarshal([]byte(file), &data)
	
	return &data, err
}

func (f FlatFile) List() ([]*models.Audio, error) {
	dirname, err := os.UserHomeDir()

	if err != nil {
		return nil, err
	}	

	metadataDirPath := filepath.Join(dirname, "audiofile")

	dirs, err := os.ReadDir(metadataDirPath)

	if err != nil {
		return nil, err	
		
	}

	audioFiles := []*models.Audio{}

	for _, dir1 := range dirs {
		if dir1.IsDir() {
			name, err := f.GetByID(dir1.Name())
			if err != nil {	
				return nil, err
			}
			audioFiles = append(audioFiles, name)
		}
	}

	return audioFiles, nil

}

func (f FlatFile) SaveMetadata(audio *models.Audio) error {
	dirname, err := os.UserHomeDir()

	if err != nil {	
		return err
	}	

	audioDirPath := filepath.Join(dirname, "audiofile", audio.Id)

	metadataFilePath := filepath.Join(audioDirPath, "metadata.json")

	file, err := os.Create(metadataFilePath)

	if err != nil {	
		return err	
	}

	defer file.Close()

	data, err := audio.JSON()

	if err != nil {
		
		fmt.Println("Error in JSON", err)
		return err
	}

	_, err = io.WriteString(file, data)

	if err != nil {	
		return err	
	}

	// Sync commits the current contents of the file to stable storage
	return file.Sync()

}

// we will return id, audiofilepath, error
func (f FlatFile) Upload(bytes []byte, filename string) (string, string, error) {
	id := uuid.New()

	dirname, err := os.UserHomeDir()

	if err != nil {	
		return id.String(), "", err
	}

	audioDirPath := filepath.Join(dirname, "audiofile", id.String())

	if err := os.MkdirAll(audioDirPath, os.ModePerm); err != nil {
		return id.String(), "", err
	}

	// dirname/audiofile/id/filename 
	audioFilePath := filepath.Join(audioDirPath, filename)

	err = os.WriteFile(audioFilePath, bytes, 0664)

	if err != nil {
		return id.String(), "", err
	}

	return id.String(), audioFilePath, nil



}

