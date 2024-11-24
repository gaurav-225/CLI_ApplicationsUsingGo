package command

import (
	"audiofile/internal/interfaces"
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
)


type GetCommand struct {
	fs *flag.FlagSet
	client interfaces.Client
	id string
}

func NewGetCommand(client interfaces.Client)  *GetCommand {
	gc  := &GetCommand{fs: flag.NewFlagSet("get", flag.ContinueOnError), 
		client: client,
	}

	gc.fs.StringVar(&gc.id, "id", "", "id of requested file")

	return gc
}

func (cmd *GetCommand) Name() string {
	return cmd.fs.Name()
}

func (cmd *GetCommand) ParseFlags(flags []string) error {
	if len(flags) == 0 {
		fmt.Println("usage: ./audiofile-cli get -id <id>")
		return fmt.Errorf("missing flags")
	}

	return cmd.fs.Parse(flags)
}

func (cmd *GetCommand) Run() error {
	if cmd.id == "" {
		return fmt.Errorf("missing id")
	}

	params := "id=" + url.QueryEscape(cmd.id)
	//fmt.Println(params)
	//path := fmt.Sprintf("http://localhost/request?%s", params)
	path := fmt.Sprintf("http://127.0.0.1:8000/get?%s", params)
	//path := fmt.Sprintf("http://127.0.0.1:8000/get?id=feca3585-6cbb-4bfd-847c-1eb00e3ec495")
	
	payload := &bytes.Buffer{}
	client := cmd.client

	req, err := http.NewRequest(http.MethodGet, path, payload)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)


	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response: ", err.Error())
		return err
	}
	fmt.Println(string(b))
	return nil

}