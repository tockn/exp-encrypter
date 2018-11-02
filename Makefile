run:
	go run main.go &
	npm run dev --prefix view

deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

build:
	npm run build --prefix view
	GOOS=linux GOARCH=amd64 go build -o linux.out
	GOOS=darwin GOARCH=amd64 go build -o mac.out
	GOOS=windows GOARCH=amd64 go build -o windows.exe

