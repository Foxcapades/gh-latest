VERSION=$(shell git describe --tags 2> /dev/null || echo "snapshot")

build:
	env CGO_ENABLED=0 GOOS=linux go build -o bin/gh-latest -ldflags "-X main.version=${VERSION}" cmd/main.go

travis:
	env CGO_ENABLED=0 GOOS=linux go build -o bin/gh-latest -ldflags "-X main.version=${VERSION}" cmd/main.go
	cd bin && tar -czf gh-latest-linux.${TRAVIS_TAG}.tar.gz ./gh-latest && rm gh-latest

	env CGO_ENABLED=0 GOOS=darwin go build -o bin/gh-latest -ldflags "-X main.version=${VERSION}" cmd/main.go
	cd bin && tar -czf gh-latest-darwin.${TRAVIS_TAG}.tar.gz ./gh-latest && rm gh-latest

	env CGO_ENABLED=0 GOOS=windows go build -o bin/gh-latest.exe -ldflags "-X main.version=${VERSION}" cmd/main.go
	cd bin && zip -9 gh-latest-windows.${TRAVIS_TAG}.zip ./gh-latest.exe && rm gh-latest.exe