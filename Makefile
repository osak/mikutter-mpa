GO_SRCS=$(wildcard src/**/*.go)

all: mpa

mpa: $(GO_SRCS) web-build
	GOPATH=`pwd` go install mpa

web-build: web
	rsync -a --progress --exclude='*.jsx' --exclude='*.*~' --exclude='js' web/ web-build/

webpack-watch:
	NODE_ENV=development ./node_modules/.bin/webpack --progress --colors --watch
