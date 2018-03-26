#!/bin/bash

export USER=$(id -u):$(id -g)
export ROOT_DIR=$(cd $(dirname $0);cd ..;pwd)

if ! type "convert" > /dev/null; then
    echo 'このスクリプトを利用するにはImageMagickが必要です.'
    echo 'Ubuntuの場合は'
    echo 'sudo apt install imagemagick'
    echo 'を実行してください.'
    exit 1
fi

rm -rf fixtures
docker-compose down
docker-compose build

docker-compose \
  run \
  -u $USER \
  resizer \
  sh -c "GOCACHE=off go run cmd/generate_fixtures/main.go -output ./fixtures"

cp $ROOT_DIR/fixtures/f.jpg $ROOT_DIR/fixtures/f-orientation.jpg
$ROOT_DIR/bin/tools/orient.sh $ROOT_DIR/fixtures/f-orientation.jpg
rm $ROOT_DIR/fixtures/f-orientation.jpg

convert -quality 9 fixtures/f.png  png8:fixtures/f-png8.png
convert -quality 9 fixtures/f.png  png24:fixtures/f-png24.png
