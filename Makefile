# Borrowed from: 
# https://github.com/silven/go-example/blob/master/Makefile
# https://vic.demuzere.be/articles/golang-makefile-crosscompile/

BINARY = gitexport
GOARCH = amd64

# Symlink into GOPATH
CURRENT_DIR=$(shell pwd)
BUILD_DIR=${CURRENT_DIR}/bin

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