all: compile

build-brew-packages:
	./generate-package-list.sh > data/brew-dependencies.txt
	go-bindata -o data.go data/...

compile:
	go build -v -o do-package-tree  *.go && go test -v && go vet


dependencies:
	go get -u github.com/jteeuwen/go-bindata/...
