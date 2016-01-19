#!/usr/bin/env bash

if [ "$xREL_TERMINAL_CLIENT_CONSUMER_KEY" == "" ] || [ "$xREL_TERMINAL_CLIENT_CONSUMER_SECRET" == "" ]; then
    echo You need to set the following env variables: xREL_TERMINAL_CLIENT_CONSUMER_KEY and xREL_TERMINAL_CLIENT_CONSUMER_SECRET. >&2
    echo Get those from http://www.xrel.to/api-apps.html >&2
    exit 1
fi

checkCommand() {
    which "$1" >/dev/null 2>&1
    if [ "$?" != "0" ]; then
        echo Please make sure the following command is available: "$1" >&2
        exit "$?"
    fi
}

checkCommand go
checkCommand tar
checkCommand zip

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# NOTE: The build process is somewhat similar to https://github.com/syncthing/syncthing, thanks for that!
platforms=(
    linux-amd64 windows-amd64 darwin-amd64 dragonfly-amd64 freebsd-amd64 netbsd-amd64 openbsd-amd64 solaris-amd64
    freebsd-386 linux-386 netbsd-386 openbsd-386 windows-386
    linux-arm linux-arm64 linux-ppc64 linux-ppc64le
)

cd "$DIR"
commit="$(git rev-parse --short HEAD 2>/dev/null)"

if [ "$1" == "" ]; then
    echo You didn\'t provide a version string as the first parameter, setting version to \"unknown\".
    version="unknown"
else
    version="$1"
fi

if [ "$commit" != "" ]; then
    version="$version"-"$commit"
fi

rm -Rf bin/
mkdir -p bin/ 2>/dev/null

for plat in "${platforms[@]}"; do
    echo Building "$plat" ...

    GOOS="${plat%-*}"
    GOARCH="${plat#*-}"

    if [ "$GOOS" != "windows" ]; then
        tmpFile="/tmp/xRELTerminalClient/bin/xREL"
    else
        tmpFile="/tmp/xRELTerminalClient/bin/xREL.exe"
    fi

    GOOS="${plat%-*}" GOARCH="${plat#*-}" go build \
    -ldflags '-X main.VERSION='"$version"' -X github.com/hashworks/xRELTerminalClient/oauth.CONSUMER_KEY='"$xREL_TERMINAL_CLIENT_CONSUMER_KEY"' -X github.com/hashworks/xRELTerminalClient/oauth.CONSUMER_SECRET='"$xREL_TERMINAL_CLIENT_CONSUMER_SECRET" \
    -o "$tmpFile" "$DIR"/xREL.go

    if [ "$?" != 0 ]; then
        echo Build failed! >&2
        exit "$?"
    fi

    if [ "$GOOS" != "windows" ]; then
        tarPath="$DIR"/bin/xREL-"$plat".tar.gz
        echo Build succeeded, creating "$tarPath" ...
        tar -czf "$tarPath" -C "${tmpFile%/*}" xREL
    else
        zipPath="$DIR"/bin/xREL-"$plat".zip
        echo Build succeeded, creating "$zipPath" ...
        zip -j "$zipPath" "$tmpFile"
    fi

    if [ "$?" != 0 ]; then
        echo Failed to pack the binary! >&2
        exit "$?"
    fi
    echo Done!

    rm "$tmpFile"

    echo
done