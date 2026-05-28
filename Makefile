MOD := cmd
TEST_FLAGS := $(if $(v),-v,)

.PHONY: run run-4x5 run-QHD build test

run:
	go -C $(MOD) run .

run-4x5:
	go -C $(MOD) run . -w 864 -h 1080 -c 2000

run-QHD:
	go -C $(MOD) run . -w 2560 -h 1440 -c 10000

build:
	go -C $(MOD) build -o ../bin/boids .

test:
	go -C $(MOD) test $(TEST_FLAGS) ./...
