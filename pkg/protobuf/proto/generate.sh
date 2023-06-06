#!/bin/sh
BASEDIR=$(dirname "$0")
echo ""
echo "Start generate proto files..."
for FILE in $BASEDIR/*.proto;
  do
    echo "Generating ${FILE##*/}"
	protoc --proto_path=$BASEDIR --micro_out=$BASEDIR/../ --go_out=$BASEDIR/../ ${FILE##*/}
done