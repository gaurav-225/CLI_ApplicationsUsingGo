package tags

import (
	"audiofile/models"
	"fmt"
	"os"

	"github.com/dhowden/tag"
)

// go get github.com/dhowden/tag
func Extract(m *models.Audio) error {
	// Extract the tags from the audio file

	f, err := os.Open(m.Path)

	if err != nil {		
		return err
	}

	defer f.Close()
	fmt.Println("filepath: ", m.Path)

	tagMetadata, err := tag.ReadFrom(f)

	if err != nil {
		return err
	}

	fmt.Println("tagMetadata: ", tagMetadata)

	m.Metadata.Tags = models.Tags{
		Title: tagMetadata.Title(),
		Artist: tagMetadata.Artist(),
		Album: tagMetadata.Album(),
		AlbumArtist: tagMetadata.AlbumArtist(),
		Composer: tagMetadata.Composer(),
		Genre: tagMetadata.Genre(),
		Year: tagMetadata.Year(),
		Comment: tagMetadata.Comment(),
		Lyrics: tagMetadata.Lyrics(),
	}

	fmt.Println("metadata: ", m.Metadata.Tags)

	return nil
}