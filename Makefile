prepare: clean copy buildcss
	go build -o bin/heimdall

clean:
	rm -rf ./bin/*
	
copy: 
	scripts/copy.sh

buildcss: 
	scripts/tailwind.sh

watchcss:
	scripts/tailwind.sh -w


run: prepare
	./bin/heimdall
