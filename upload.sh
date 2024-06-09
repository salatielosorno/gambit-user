export GOOS=linux
export GOARCH=amd64
go build main.go
rm main.zip
zip -r main.zip main