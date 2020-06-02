export PORT = 8080
export GIN_MODE=debug
BINARY = NoteletServer

all: install start
install:
	go install
start:
	$(BINARY)
kill:
	pkill $(BINARY)
