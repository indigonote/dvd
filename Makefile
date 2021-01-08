.PHONY: test build

BUILD := ./build

install:
	go get github.com/mitchellh/gox

test:
	go test -count 1 -cover ./...

clean:
	mkdir -p $(BUILD)
	rm -Rf $(BUILD)/*

build: clean
	gox -osarch=linux/amd64 -osarch=linux/arm64 -osarch=darwin/amd64 -output="$(BUILD)/{{.Dir}}_{{.OS}}_{{.Arch}}"
	cd $(BUILD) && find . -type f ! -name '*.exe' | xargs -I % sh -c "mv % dvd && zip %.zip dvd && rm -f dvd"
