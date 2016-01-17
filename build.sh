#!/usr/bin/env bash

if [ "$xREL_TERMINAL_CLIENT_CONSUMER_KEY" == "" ] || [ "$xREL_TERMINAL_CLIENT_CONSUMER_KEY" == "" ]; then
    echo "You need to set the following env variables: xREL_TERMINAL_CLIENT_CONSUMER_KEY and xREL_TERMINAL_CLIENT_CONSUMER_KEY."
    echo "Get those from http://www.xrel.to/api-apps.html"
fi

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# NOTE: The build process is somewhat similar to https://github.com/syncthing/syncthing, thanks for that!
platforms=(
    darwin-amd64 dragonfly-amd64 freebsd-amd64 linux-amd64 netbsd-amd64 openbsd-amd64 solaris-amd64 windows-amd64
    freebsd-386 linux-386 netbsd-386 openbsd-386 windows-386
    linux-arm linux-arm64 linux-ppc64 linux-ppc64le
)

rm -Rf "$DIR"/bin/
mkdir -p "$DIR"/bin/ 2>/dev/null

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
    -ldflags '-X github.com/hashworks/xRELTerminalClient/oauth.CONSUMER_KEY='"$xREL_TERMINAL_CLIENT_CONSUMER_KEY"' -X github.com/hashworks/xRELTerminalClient/oauth.CONSUMER_SECRET='"$xREL_TERMINAL_CLIENT_CONSUMER_SECRET" \
    -o "$tmpFile" "$DIR"/xREL.go

    if [ "$?" != 0 ]; then
        echo "Failed!"
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
        echo "Failed to pack the binary!"
        exit "$?"
    fi
    echo Done!

    rm "$tmpFile"

    echo
done