.PHONY: build clean deploy gomodgen schemagen install_dependencies


install_dependencies:
	go get -u github.com/go-bindata/go-bindata/...

schemagen:
	go generate ./...

build: schemagen
	export GO111MODULE=on
	#env GOOS=linux go build -ldflags="-s -w" -o bin/graphql ./graphql
	go build -ldflags="-s -w" -o bin/graphql ./graphql

clean:
	rm -rf ./bin ./vendor Gopkg.lock

run: build
	SQS_QUEUE_NAME=go-graphql-sqs-example bin/graphql

test:
	go test ./...

deploy: clean test build
	go install

docker_build:
	docker build -t go-graphql-sqs-example .

docker_run: docker_build
	docker run -p 3000:3000 --env-file=.env go-graphql-sqs-example


show_aws_config:
	tail -n 8 ~/.aws/config

run_as_iam_user:
	aws-vault exec go-sqs-example -- make run

run_as_iam_role:
	aws-vault exec go-sqs-example-role -- make run
