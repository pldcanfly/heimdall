#! /bin/bash

mkdir -p ./bin/web/static/css
if [ "$1" ]; then
    ./build/tailwindcss -i web/css/style.css -o web/static/css/style.css -c build/tailwind.config.js --watch
else 
    ./build/tailwindcss -i web/css/style.css -o web/static/css/style.css -c build/tailwind.config.js 
fi