package command

import (
	"audiofile/internal/interfaces"
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
)

type ListCommand struct  {
	fs *flag.FlagSet
	client interfaces.Client
}


func NewListCommand(client interfaces.Client) *ListCommand {
	lc := &ListCommand{
		fs:     flag.NewFlagSet("list", flag.ContinueOnError),
		client: client,
	}
	return lc
}


func (cmd *ListCommand) Name() string {
	return cmd.fs.Name()
}

func (cmd *ListCommand) ParseFlags(flags []string) error {
	return cmd.fs.Parse(flags)
}

func (cmd *ListCommand) Run() error {
	//path := "http://localhost/list"
	path := "http://127.0.0.1:8000/list"
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
		return err
	}
	fmt.Println(string(b))
	return nil
}