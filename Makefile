run:
	go run main.go &
	npm run dev --prefix view

deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

install:
	wget https://dl.google.com/go/go1.11.1.linux-amd64.tar.gz &&
  tar -C /usr/local -xzf go11.1.linux-amd64.tar.gz &&
  export PATH=$PATH:/usr/local/go/bin >> ~/.bashrc` &&
	mkdir ~/go &&
	echo "export GOPATH=/$HOME/go" >> ~/.bashrc &&
	echo "export PATH=$PATH:$GOPATH/bin" >> ~/.bashrc &&
	source ~/.bashrc &&
	go version &&
	sudo yum install -y nodejs &&
	sudo yum install -y npm &&
  npm install -g vue-cli &&
	go get -u github.com/tockn/exp-encrypter

