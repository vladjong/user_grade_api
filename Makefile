APP_BIN = build/app
CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
GOVER=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
MOCKGEN=${BINDIR}/mockgen_${GOVER}
SMARTIMPORTS=${BINDIR}/smartimports_${GOVER}
LINTVER=v1.49.0
LINTBIN=${BINDIR}/lint_${GOVER}_${LINTVER}

all: build

build: clean $(APP_BIN)

$(APP_BIN):
	go build -o $(APP_BIN) cmd/main/main.go
	./build/app

clean:
	rm -rf build || true

lint: install-lint
	${LINTBIN} run

bindir:
	mkdir -p ${BINDIR}

install-lint: bindir
	test -f ${LINTBIN} || \
		(GOBIN=${BINDIR} go install github.com/golangci/golangci-lint/cmd/golangci-lint@${LINTVER} && \
		mv ${BINDIR}/golangci-lint ${LINTBIN})
