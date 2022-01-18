BIN_DIR = bin
export GOPATH ?= $(shell go env GOPATH)
export GO111MODULE ?= on
export PROJECT_SERUM_IMAGE ?= projectserum/build:v0.20.0

install:
	go mod download
	go get github.com/onsi/ginkgo/v2/ginkgo/generators@v2.0.0-rc2
	go get github.com/onsi/ginkgo/v2/ginkgo/internal@v2.0.0-rc2
	go get github.com/onsi/ginkgo/v2/ginkgo/labels@v2.0.0-rc2
	go get github.com/smartcontractkit/chainlink-relay/ops@v0.0.0-20211215192527-583f627029d9
	go install github.com/onsi/ginkgo/v2/ginkgo

copy_keys:
	./scripts/programs-keys-cp.sh

start_anchor:
	docker run --rm -it -v $(shell pwd):/workdir --entrypoint bash ${PROJECT_SERUM_IMAGE}

start_test_env:
	echo "TODO"

start_pulumi_env:
	echo "TODO"

:PHONY
build_js:
	cd gauntlet && yarn install --frozen-lockfile && yarn bundle

build:
	cd gauntlet && yarn install --frozen-lockfile && yarn bundle
	docker run --rm -it -v $(shell pwd):/workdir ${PROJECT_SERUM_IMAGE} /bin/bash ./scripts/anchor-build.sh

test_smoke:
	./scripts/programs-keys-cp.sh
	SELECTED_NETWORKS=solana NETWORK_SETTINGS=$(shell pwd)/tests/e2e/networks.yaml ginkgo tests/e2e/smoke

test_ocr:
	./scripts/programs-keys-cp.sh
	SELECTED_NETWORKS=solana NETWORK_SETTINGS=$(shell pwd)/tests/e2e/networks.yaml ginkgo --focus=@ocr tests/e2e/smoke

test_gauntlet:
	./scripts/programs-keys-cp.sh
	SELECTED_NETWORKS=solana NETWORK_SETTINGS=$(shell pwd)/tests/e2e/networks.yaml ginkgo --focus=@gauntlet tests/e2e/smoke