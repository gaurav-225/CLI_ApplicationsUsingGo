package metadata

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)


func (m *MetadataService) getByIdHandler(res http.ResponseWriter, req *http.Request) {
	value, ok := req.URL.Query()["id"]

	if !ok || len(value[0]) < 1 {
		fmt.Println("Url Param 'id' is missing")
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	id := value[0]

	fmt.Println("Url Param 'id' is: " + string(id))

	audio, err := m.Storage.GetByID(id)

	if err != nil {
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "no such file or directory") {
			res.WriteHeader(http.StatusNotFound)
			return
		}
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	audoString, err := audio.JSON()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	io.WriteString(res, audoString)
}