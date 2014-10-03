include ../includes.mk

COMPONENT = cache
IMAGE = $(IMAGE_PREFIX)$(COMPONENT):$(BUILD_TAG)
DEV_IMAGE = $(DEV_REGISTRY)/$(IMAGE)

check-create:
	@test -s ./rootfs.tar || { echo "rootfs.tar does not exist! Exiting...";  exit 1; }
	@sudo chown 1000:1000 rootfs*

create: check-create
	@sudo docker build -t tarmaker:${COMPONENT} tarmaker
	@sudo docker run --name builder-${COMPONENT} tarmaker:${COMPONENT}
	@sudo docker cp builder-${COMPONENT}:/tmp/redis-dist.tar.gz ./
	@test ! -s ./redis-dist.tar.gz || tar -xzf ./redis-dist.tar.gz
	@sudo docker build -t ${IMAGE} .
	@test ! -s ./redis-server || rm ./redis-s*
	@test ! -s ./redis-cli || rm ./redis-c*
	@sudo docker rm -f builder-${COMPONENT}

build: check-docker
	docker build -t $(IMAGE) .

clean: check-docker check-registry
	docker rmi $(IMAGE)

full-clean: check-docker check-registry
	docker images -q $(IMAGE_PREFIX)$(COMPONENT) | xargs docker rmi -f

install: check-deisctl
	deisctl scale $(COMPONENT)=1

uninstall: check-deisctl
	deisctl scale $(COMPONENT)=0

start: check-deisctl
	deisctl start $(COMPONENT)

stop: check-deisctl
	deisctl stop $(COMPONENT)

restart: stop start

run: install start

dev-release: check-registry check-deisctl
	docker tag $(IMAGE) $(DEV_IMAGE)
	docker push $(DEV_IMAGE)
	deisctl config $(COMPONENT) set image=$(DEV_IMAGE)

release:
	docker push $(IMAGE)

deploy: build dev-release restart

test: test-unit test-functional

test-unit:
	@echo no unit tests

test-functional:
	@docker history deis/test-etcd >/dev/null 2>&1 || docker pull deis/test-etcd:latest
	GOPATH=$(CURDIR)/../tests/_vendor:$(GOPATH) go test -v ./tests/...
