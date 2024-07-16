.PHONY = build

meta_path = github.com/atomicptr/tmplr/pkg/meta
git_commit := $(shell git rev-list -1 HEAD)
git_version := $(shell git --no-pager tag --points-at HEAD | head -n 1)

build:
	go build -o bin/tmplr \
		-ldflags "\
			-X $(meta_path).Version=$(git_version) \
			-X $(meta_path).GitCommit=$(git_commit)" \
		cmd/tmplr/main.go
test:
	go test -v ./...
