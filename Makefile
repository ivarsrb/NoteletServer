BINARY = NoteletServer

all: install start
install:
	go install
start:
	$(BINARY) -addr=":8080" -static="web"
kill:
	pkill $(BINARY)
