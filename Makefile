GOFLAGS:=-mod=mod
# Build the app
#.PHONY: build
build:
	$(GOFLAGS) go build -o bin/gopher-cli main.go

# Run the app
run:
	$(GOFLAGS) go run main.go

# Remove all retrieved *.png files
clean:
	rm -rf img && rm -rf bin

## demo
demo:
	$(GOFLAGS) go build -o bin/gopher-cli main.go
	bin/gopher-cli get friends
