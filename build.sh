#!/bin/bash

platforms=(
    "windows/amd64"
    "linux/amd64"
    "linux/arm64"
    "darwin/amd64"
    "darwin/arm64"
)

for platform in "${platforms[@]}"
do
    GOOS=${platform%/*}
    GOARCH=${platform#*/}
    output_name="mantela-adder-${GOOS}-${GOARCH}"
    if [ "$GOOS" = "windows" ]; then
        output_name+=".exe"
    fi

    echo "Building for $GOOS/$GOARCH..."
    env GOOS=$GOOS GOARCH=$GOARCH go build -o "$output_name" main.go
    if [ $? -ne 0 ]; then
        echo "An error has occurred! Aborting the script execution..."
        exit 1
    fi

    # 作業ディレクトリ作成
    workdir="dist-${GOOS}-${GOARCH}"
    mkdir -p "$workdir/static"
    cp "$output_name" "$workdir/"
    cp config.sample.toml override.json "$workdir/"
    cp -r static/* "$workdir/static/"

    # ZIP圧縮
    (cd "$workdir" && zip -r "../${workdir}.zip" .)
    rm -r "$workdir"
    rm "$output_name"
done
