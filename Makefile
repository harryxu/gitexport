# https://gist.github.com/turtlemonvh/38bd3d73e61769767c35931d8c70ccb4

BINARY = gitexport
GOARCH = amd64

# Symlink into GOPATH
CURRENT_DIR=$(shell pwd)
BUILD_DIR=${CURRENT_DIR}/bin

LDFLAGS = -ldflags "-s -w"

# Build the project
all: clean linux darwinx64 darwinarm64 windows

linux: 
	GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BUILD_DIR}/${BINARY}-linux-${GOARCH} . ; 

darwinx64:
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ${BUILD_DIR}/${BINARY}-darwinx64 . ; 

darwinarm64:
	GOOS=darwin GOARCH=arm64 go build ${LDFLAGS} -o ${BUILD_DIR}/${BINARY}-darwinarm64 . ;

windows:
	GOOS=windows GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BUILD_DIR}/${BINARY}-windows-${GOARCH}.exe . ; 

clean:
	-rm -f ${BUILD_DIR}/${BINARY}-*

.PHONY: link linux darwin windows clean