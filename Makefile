.PHONY: test
test:
	which gotestsum || (pushd /tmp && go install gotest.tools/gotestsum@latest && popd)
	gotestsum -- --mod=vendor -bench=^$$ -race ./...

.PHONY: cyclo
cyclo:
	which gocyclo || (pushd /tmp && go install github.com/fzipp/gocyclo/cmd/gocyclo@latest && popd)
	gocyclo -over 10 -ignore 'vendor/' .

.PHONY: vet
vet:
	go vet ./...

.PHONY: quality
quality: cyclo vet

.PHONY: pr-ready
pr-ready: test quality