prepare: clean copy buildcss
	go build -o bin/heimdall

clean:
	rm -rf ./bin/*
	
copy: 
	scripts/copy.sh

buildcss: 
	scripts/buildcss.sh

watchcss:
	scripts/buildcss.sh -w


run: prepare
	cd ./bin && ./heimdall
