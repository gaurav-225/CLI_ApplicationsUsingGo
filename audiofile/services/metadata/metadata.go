package metadata

import (
	"audiofile/internal/interfaces"
	"audiofile/storage"
	"fmt"
	"net/http"
)

type MetadataService struct {
	Server *http.Server
	Storage interfaces.Storage
}


func CreateMetadaaSevice(port int, storage interfaces.Storage) *MetadataService {

	// it is a router that will handle the incoming requests
	mux := http.NewServeMux()

	metadataService := &MetadataService{
		Server: &http.Server{
			Addr: fmt.Sprintf(":%v", port),
			Handler: mux,
		},
		Storage: storage,
	}

	mux.HandleFunc("/upload", metadataService.uploadHandler)
	mux.HandleFunc("/list", metadataService.listHandler)
	mux.HandleFunc("/get", metadataService.getByIdHandler)

	return metadataService
}

func Run(port int) *http.Server {
	flatFileStorage := storage.FlatFile{}

	service := CreateMetadaaSevice(port, flatFileStorage)

	err := service.Server.ListenAndServe()

	if err != nil {	
		fmt.Println("api not started", err)
	}

	return service.Server


}


