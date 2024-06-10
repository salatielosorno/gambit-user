export GOOS=linux
export GOARCH=amd64

#https://stackoverflow.com/questions/78143353/aws-lambda-with-go-runtime-invalidentrypoint-error-for-post-request-to-api-ga
go build -o bootstrap main.go
rm main.zip
zip main.zip bootstrap