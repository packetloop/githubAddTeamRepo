PROJECT_NAME := githubAddTeamRepo
package = github.com/packetloop/$(PROJECT_NAME)


.PHONY: vendor
vendor: dep
	dep ensure

.PHONY: dep
dep:
	go get github.com/tcnksm/ghr
	go get github.com/mitchellh/gox
	go get github.com/golang/dep/cmd/dep
	go get github.com/goreleaser/goreleaser

.PHONY: env
env:
ifndef GITHUB_TOKEN
	$(error GITHUB_TOKEN is not set)
endif

.PHONY: build
build: dep
	gox -output="./release/{{.Dir}}_{{.OS}}_{{.Arch}}" -os="linux windows darwin" -arch="amd64" .

.PHONY: build-local
build-local: dep
	go build -o examples/githubAddTeamRepo

.PHONY: create-tag
create-tag: next-tag
	 git fetch --tags packetloop
	 git tag -a v$(TAG) -m "v$(TAG)"
	 git push packetloop v$(TAG)

.PHONY: release
release: dep
	goreleaser

.PHONY: next-tag
next-tag:
ifndef TAG
	$(error TAG is not set)
endif