export PORT=8080
export DATABASE_URL=postgres://testusr:testpass123@localhost/testdb
export GIN_MODE=debug
BINARY = NoteletServer

all: install start
install:
	go install
start:
	$(BINARY)
kill:
	pkill $(BINARY)
test:
	go test -v ./...