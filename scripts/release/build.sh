#!/usr/bin/env bash
# Build forge for each platform a release carries, into dist/, with checksums.
#
# Usage: scripts/release/build.sh <label>
#
# <label> names the archives, as forge_<label>_<os>_<arch>: a version such as
# 0.2.0, or "latest". Run it in a checkout of the commit to build: go build
# stamps the binary with that commit's version, a tag's or a pseudo-version,
# which forge version prints.
set -euo pipefail

label=${1:?usage: build.sh <label>}
targets="linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64"

rm -rf build dist
mkdir -p build dist
for target in $targets; do
  os=${target%/*}
  arch=${target#*/}
  name="forge_${label}_${os}_${arch}"
  bin=forge
  [[ $os == windows ]] && bin=forge.exe
  mkdir -p "build/$name"
  GOOS=$os GOARCH=$arch CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o "build/$name/$bin" ./cmd/forge
  cp -r README.md docs "build/$name/"
  if [[ $os == windows ]]; then
    (cd build && zip -qr "../dist/$name.zip" "$name")
  else
    tar -C build -czf "dist/$name.tar.gz" "$name"
  fi
done
(cd dist && sha256sum -- * >checksums.txt)
echo "Built $(find dist -type f | wc -l) files into dist/"
