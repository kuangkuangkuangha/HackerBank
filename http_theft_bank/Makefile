GOFILES=`find . -name "*.go" -type f`
PACKAGES=`go list ./...`
VETPACKAGES=`go list ./... | grep -v /examples/`
PROJECTNAME=`go list .`

default:tidy fmt
	@go build -o main -v .

clean:
	@rm -f main
	@find . -name "[._]*.s[a-w][a-z]" | xargs -i rm -f {}

fmt:
	@goimports -w -local ${PROJECTNAME} ${GOFILES} # fmt package
	@gofmt -s -w ${GOFILES} # fmt file

fmt-check:
	@diff=$$(gofmt -s -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

tidy:
	@go mod tidy

test:
	@go test -v -count=1  ./...

docker:
	@docker build -t muxih4ck/Go-Web-Application-Template:latest .

githook:
	@sudo npm install yarn
	@yarn global add husky
	@npm --registry https://registry.npm.taobao.org install --save-dev @commitlint/{config-conventional,cli}
	@husky install
	@npx husky add .husky/commit-msg 'npx --no -- commitlint --edit "$1"'
	@curl https://gist.githubusercontent.com/kocoler/517fbd7a806d69b4b7a80fef29c14414/raw/cccdfb4f35c4c8fbc17f4c9455ddcc2944d21ee6/commitlint.config.js > commitlint.config.js

cleangithook:
	@git config --unset core.hooksPath
	@rm -rf .husky/ >/dev/null 2>&1
	@rm commitlint.config.js >/dev/null 2>&1

ca: openssl req -new -nodes -x509 -out conf/server.crt -keyout conf/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=127.0.0.1/emailAddress=xxxxx@qq.com"

help:
	@echo "make - compile the source code and tidy fmt"
	@echo "make fmt - format the source code"
	@echo "make fmt-check - check the source code"
	@echo "make tidy - go mod tidy"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make docker - build a docker image"
	@echo "make ca - generate ca files"
	@echo "make githook - add git hook(with husky)"
	@echo "make cleangithook - clean git hook"

.PHONY: default clean fmt fmt-check tidy test docker ca help githook cleangithook