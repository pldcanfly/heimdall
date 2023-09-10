build: clean copy
	go build -o bin/heimdall
	

clean:
	rm -rf ./bin/*
	
copy: 
	mkdir ./bin/templates
	cp -r ./templates/* ./bin/templates/

run: build
	./bin/heimdall
