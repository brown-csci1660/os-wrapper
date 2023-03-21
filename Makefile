all: build # Default target

# Compile the code
build:
	go build -o oswrapper.a *.go

# Remove any build files
clean:
	rm -fv oswrapper.a
