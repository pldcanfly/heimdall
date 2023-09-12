build: clean copy buildcss
	go build -o bin/heimdall

	

clean:
	rm -rf ./bin/*
	
copy: 
	mkdir -p ./bin/frontpage/components
	cp -r ./frontpage/components/* ./bin/frontpage/components
	mkdir -p ./bin/frontpage/static
	cp -r ./frontpage/static/* ./bin/frontpage/static
	mkdir -p ./bin/frontpage/templates
	cp -r ./frontpage/templates/* ./bin/frontpage/templates

buildcss: 
	mkdir -p ./bin/frontpage/static/css
	./tailwindcss -i frontpage/css/style.css -o frontpage/static/css/style.css

watchcss:
		mkdir -p ./bin/frontpage/static/css
	./tailwindcss -i frontpage/css/style.css -o frontpage/static/css/style.css --watch


run: build
	./bin/heimdall
