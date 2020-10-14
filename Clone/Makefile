all:
	go build -v -o github.com/aferrercrafter/operacion-fuego-quazar/cmd/topsecret

build:
	GOOS=linux go build -o github.com/aferrercrafter/operacion-fuego-quazar/cmd/topsecret

zip: build
	zip deployment.zip github.com/aferrercrafter/operacion-fuego-quazar/cmd/topsecret