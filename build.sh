#!/usr/bin/env bash

if [ "$xREL_TERMINAL_CLIENT_CONSUMER_KEY" == "" ] || [ "$xREL_TERMINAL_CLIENT_CONSUMER_SECRET" == "" ]; then
    echo You need to set the following env variables: xREL_TERMINAL_CLIENT_CONSUMER_KEY and xREL_TERMINAL_CLIENT_CONSUMER_SECRET. >&2
    echo Get those from http://www.xrel.to/api-apps.html >&2
    exit 1
fi

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

cd "$DIR"
commit="$(git rev-parse --short HEAD 2>/dev/null)"

if [ "$commit" != "" ]; then
    version=dev-"$commit"
fi

rm -Rf ./bin/
mkdir ./bin/

go build -ldflags '-X main.VERSION='"$version"'
                     -X main.OAUTH_CONSUMER_KEY='"$xREL_TERMINAL_CLIENT_CONSUMER_KEY"'
                     -X main.OAUTH_CONSUMER_SECRET='"$xREL_TERMINAL_CLIENT_CONSUMER_SECRET" \
    -o "./bin/xREL" "$DIR"/src/*.go