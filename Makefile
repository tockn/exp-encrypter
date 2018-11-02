run:
	go run main.go &
	npm run dev --prefix view

deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure
