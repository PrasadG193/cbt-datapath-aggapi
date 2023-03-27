IMAGE_REPO_CLIENT ?= prasadg193/cbt-client
IMAGE_REPO_AGGAPI ?= prasadg193/cbt-datapath
IMAGE_REPO_SAMPLE_DRIVER ?= prasadg193/sample-driver
IMAGE_TAG_CLIENT ?= latest
IMAGE_TAG_AGGAPI ?= latest
IMAGE_TAG_SAMPLE_DRIVER ?= latest

API_GROUP ?= cbt
API_VERSION ?= v1alpha1
API_KIND ?= VolumeSnapshotDeltaToken

GOOS ?= linux
GOARCH ?= amd64

image:
	docker build -t $(IMAGE_REPO_AGGAPI):$(IMAGE_TAG_AGGAPI) -f Dockerfile .
	docker build -t $(IMAGE_REPO_SAMPLE_DRIVER):$(IMAGE_TAG_SAMPLE_DRIVER) -f Dockerfile-grpc .
	#docker build -t $(IMAGE_REPO_CLIENT):$(IMAGE_TAG_CLIENT) -f Dockerfile-client .

push:
	docker push $(IMAGE_REPO_AGGAPI):$(IMAGE_TAG_AGGAPI)
	docker push $(IMAGE_REPO_SAMPLE_DRIVER):$(IMAGE_TAG_SAMPLE_DRIVER)
	#docker push $(IMAGE_REPO_CLIENT):$(IMAGE_TAG_CLIENT)

build:
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -a -o cbt-client ./cmd/mock/client/main.go
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -a -o sample-driver ./cmd/mock/sample-driver/main.go
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -a -o apiserver ./cmd/apiserver/main.go

codegen:
	./hack/update-codegen.sh

codegen-verify:
	./hack/verify-codegen.sh

driver:
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -a -o sample-driver ./cmd/mock/sample-driver/main.go

init_repo:
	apiserver-boot init repo --domain storage.k8s.io

create_group:
	apiserver-boot create group version resource --group $(API_GROUP) --version $(API_VERSION) --kind $(API_KIND)

.PHONY: yaml
yaml:
	rm -rf yaml-generated
	apiserver-boot build config --name cbt-datapath --namespace cbt-svc --image $(IMAGE_REPO_AGGAPI):$(IMAGE_TAG_AGGAPI) --output yaml-generated

.PHONY: proto
proto:
	protoc -I=proto \
		--go_out=pkg/grpc --go_opt=paths=source_relative \
   	--go-grpc_out=pkg/grpc --go-grpc_opt=paths=source_relative \
		proto/cbt.proto
