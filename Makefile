deps:
	go get -v && cd front && npm i

build:	
	make deps
	cd front && npx vue-tsc && npx vite build
	go build -o server.bin .
