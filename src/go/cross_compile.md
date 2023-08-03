cross compile



x64

env GOOS=windows GOARCH=amd64 go build package-import-path
x32

env GOOS=windows GOARCH=386 go build package-import-path
