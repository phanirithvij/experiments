echo "[build] go-bindata embedding assets"
go generate

PKG_ROOT="github.com/phanirithvij/experiments/go-exps"
GIT_COMMIT=$(git rev-list -1 HEAD)
DATE=$(date -R)
VERSION="0.0.1"
NAME="go-exps"

BINDIR="dist"
mkdir -p $BINDIR

PLATFORMS=("windows/amd64" "windows/386" "darwin/amd64" "linux/amd64")

package_split=(${PKG_ROOT//\// })
package_name=${package_split[-1]}


for platform in "${PLATFORMS[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}

    BIN=$package_name'_'$GOOS'_'$GOARCH'_'$VERSION
    BIN="$BINDIR/$BIN"

    if [[ $GOOS == "windows" ]]
    then
        BIN+='.exe'
    fi

    echo "[build] Version $VERSION"
    echo "[build] Commit $GIT_COMMIT"
    echo "[build] BuildTime $DATE"
    echo "[build] Building to $BIN"
    echo "[build] OS: $GOOS"
    echo "[build] ARCH: $GOARCH"

    LDFLAGS="-X '$PKG_ROOT/config.Version=$VERSION' -X '$PKG_ROOT/config.BuildTime=$DATE' -X '$PKG_ROOT/config.CommitID=$GIT_COMMIT'"
    env GOOS=$GOOS GOARCH=$GOARCH go build -ldflags "$LDFLAGS" -o $BIN $PKG_ROOT

    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi

    echo "[build] UPX compressing $BIN"
    upx --brute -v $BIN
    echo "[build] UPX testing $BIN"
    upx -t -v $BIN
    echo -e "[build] file info\n\n"
    file $BIN
done
