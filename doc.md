```bash

go version

go run main.go

# build a binary and run it
go build -o main main.go

# run executable
./main

# make a go module
go mode init learn

# install stuff after creating module, inside current module
go get gopl.io/ch1/helloworld

# run the installed file
go run gopl.io/ch1/helloworld

# go code formatting
go fmt

go get golang.org/x/tools/cmd/goimports


# show docs
go doc fmt
go doc fmt.Printf
# my own local package
go doc even

# testing
# go test -> this just works for working directory

# specific package
go test ./even/

# all packages
go test ./...

```
