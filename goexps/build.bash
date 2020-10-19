echo "[build] removing previously built dist dir (if exists)"

exe (){
    set -x
    "$@"
    set +x
}

exe rm -rf dist/

echo "[build] go-bindata embedding assets"
exe go generate

PKG_ROOT="github.com/phanirithvij/experiments/goexps"
CONFIG_FILE="$PKG_ROOT/experiments/config"
GIT_COMMIT=$(git rev-list -1 HEAD)
# Must use -R to convert it inside go as a time.Time
DATE=$(date -R)
VERSION="0.0.1"

DIST_DIR="dist"

PLATFORMS=("windows/amd64" "windows/386" "darwin/amd64" "linux/amd64" "linux/386")

package_split=(${PKG_ROOT//\// })
package_name=${package_split[-1]}

for platform in "${PLATFORMS[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    
    BIN=$package_name'_'$GOOS'_'$GOARCH
    DESTDIR=$DIST_DIR/$VERSION
    BIN="$DESTDIR/$BIN"
    
    mkdir -p $DESTDIR
    
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
    
    VERSIONFLAG="-X '$CONFIG_FILE.Version=$VERSION'"
    COMMITFLAG="-X '$CONFIG_FILE.CommitID=$GIT_COMMIT'"
    BUILDFLAG="-X '$CONFIG_FILE.BuildTime=$DATE'"
    ARCHFLAG="-X '$CONFIG_FILE.Architecture=$GOARCH'"
    OSFLAG="-X '$CONFIG_FILE.Platform=$GOOS'"
    
    LDFLAGS="$VERSIONFLAG $BUILDFLAG $COMMITFLAG $ARCHFLAG $OSFLAG"
    exe env GOOS=$GOOS GOARCH=$GOARCH go build -ldflags "$LDFLAGS" -o $BIN $PKG_ROOT
    
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
    
    echo "[build] UPX compressing $BIN"
    exe upx --lzma -v $BIN
    # upx --brute -v $BIN
    echo "[build] UPX testing $BIN"
    exe upx -t -v $BIN
    echo -e "[build] file info\n"
    exe file $BIN
done
