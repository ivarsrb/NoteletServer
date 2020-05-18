BINARY = NoteletServer

install:
	go install
start:
	$(BINARY)
kill:
	pkill $(BINARY)
all: install start
