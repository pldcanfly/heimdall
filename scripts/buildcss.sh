#! /bin/bash
mkdir -p ./bin/web/static/css
if [ "$1" ]; then
    sass --style=compressed --watch ./web/scss:./bin/web/static/css
else 
     sass --style=compressed ./web/scss:./bin/web/static/css
fi
