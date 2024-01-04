#!/bin/bash

set -e

base_dir="$(dirname "${BASH_SOURCE[0]}" | xargs realpath)/.."

declare -a ALL_OS=("windows" "linux")
declare -a ALL_ARCH=("amd64" "arm64")

for os in "${ALL_OS[@]}"; do
    for arch in "${ALL_ARCH[@]}"; do
        GOOS="${os}" GOARCH="${arch}" "${base_dir}/hack/build.sh"
    done
done
