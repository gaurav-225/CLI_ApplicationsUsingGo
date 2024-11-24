# CLI_ApplicationsUsingGo



```bash
# start api server
go run cmd/api/main.go -port 8000
```



```bash
# build cli and use to communicate with api server.
cd audiofile
go build -o audiofile-cli cmd/cli/main.go
./audiofile-cli list
```

```bash
# get specific request

# ./audiofile-cli get -id=<id>
./audiofile-cli get -id=feca3585-6cbb-4bfd-847c-1eb00e3ec495
```


```bash
# ./audiofile-cli upload -filename <audio file> 
./audiofile-cli upload -filename ~/Downloads/beatdoctor.mp3 
```