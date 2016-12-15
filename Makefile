GO_SRCS=$(wildcard src/**/*.go)
.PHONY: web

all: mpa

mpa: $(GO_SRCS) web
	GOPATH=`pwd` go install mpa

web:
	mkdir -p web-build/static
	rsync -a --progress --exclude='*.jsx' --exclude='*.*~' --exclude='js' --exclude='index.html' web/ web-build/static/
	rsync -a --progress web/index.html web-build/

webpack-watch:
	NODE_ENV=development ./node_modules/.bin/webpack --progress --colors --watch
