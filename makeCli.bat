env GOOS=darwin  GOARCH=amd64 go build -o target/osx64/cli
env GOOS=linux   GOARCH=amd64 go build -o target/linux64/cli
env GOOS=windows GOARCH=amd64 go build -o target/win64/cli.exe