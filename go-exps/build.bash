echo "[build] go-bindata embedding assets"
go generate

PKG_ROOT="github.com/phanirithvij/experiments/go-exps"
GIT_COMMIT=$(git rev-list -1 HEAD)
DATE=$(date -R)
VERSION="0.0.1"
NAME="go-exps"

BINDIR="dist"
mkdir -p $BINDIR

PLATFORMS=("windows/amd64" "windows/386" "darwin/amd64", "linux/amd64")

for platform in "${PLATFORMS[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}

    BIN=$PKG_ROOT'-'$GOOS'-'$GOARCH'-'$VERSION
    BIN="$BINDIR/$BIN"

    if [[ $GOOS == "windows" ]]
    then
        BIN+='.exe'
    fi

    echo "[build] Version $VERSION"
    echo "[build] Commit $GIT_COMMIT"
    echo "[build] BuildTime $DATE"
    echo "[build] Building to $BIN"

    go build -ldflags "-X '$PKG_ROOT/config.Version=$VERSION' -X '$PKG_ROOT/config.BuildTime=$DATE' -X '$PKG_ROOT/config.CommitID=$GIT_COMMIT'" -o $BIN
    file $BIN
    upx --brute -v $BIN
done
