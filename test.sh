#!/bin/sh

set -e

mkdir -p /secret
printf "%s" "$GOOGLE_AUTH_JSON" > /secret/google-auth.json

echo "" > coverage.txt
for d in $(go list ./... | grep -v vendor); do
  go test -race -coverprofile=profile.out -covermode=atomic $d
  if [ -f profile.out ]; then
    cat profile.out >> coverage.txt
    rm profile.out
  fi
done
if [ "$CI" = "true" ]; then
  curl -o upload.sh https://codecov.io/bash
  bash ./upload.sh
fi
rm -rf coverage.txt
