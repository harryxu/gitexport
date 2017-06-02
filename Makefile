# https://gist.github.com/turtlemonvh/38bd3d73e61769767c35931d8c70ccb4

BINARY = gitexport
GOARCH = amd64

# Symlink into GOPATH
CURRENT_DIR=$(shell pwd)
BUILD_DIR=${CURRENT_DIR}/bin

LDFLAGS = -ldflags "-s -w"

# Build the project
all: clean linux darwin windows

linux: 
	GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BUILD_DIR}/${BINARY}-linux-${GOARCH} . ; 

darwin:
	GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BUILD_DIR}/${BINARY}-darwin-${GOARCH} . ; 

windows:
	GOOS=windows GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BUILD_DIR}/${BINARY}-windows-${GOARCH}.exe . ; 

clean:
	-rm -f ${BUILD_DIR}/${BINARY}-*

.PHONY: link linux darwin windows clean