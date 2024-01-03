#!/bin/bash

set -e

base_dir="$(dirname "${BASH_SOURCE[0]}" | xargs realpath)/.."

bin_dir="${base_dir}/bin"

GOOS="${GOOS:-$(go env GOOS)}"
GOARCH="${GOARCH:-$(go env GOARCH)}"

GO_BUILD_FLAGS="${GO_BUILD_FLAGS:--ldflags="-w -s"}"

output_name="${bin_dir}/infraspace-savegame-editor_${GOOS}_${GOARCH}"

if [ "${GOOS}" == "windows" ]; then
    output_name="${output_name}.exe"
fi

pushd "${base_dir}" >/dev/null

echo "Building $(basename "${output_name}")"
GOOS="${GOOS}" GOARCH="${GOARCH}" go build "${GO_BUILD_FLAGS}" -o "${output_name}" ./cmd/save-editor/...
