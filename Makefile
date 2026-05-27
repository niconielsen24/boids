MOD := cmd
TEST_FLAGS := $(if $(v),-v,)

.PHONY: run build test

run:
	go -C $(MOD) run .

build:
	go -C $(MOD) build -o ../bin/boids .

test:
	go -C $(MOD) test $(TEST_FLAGS) ./...
