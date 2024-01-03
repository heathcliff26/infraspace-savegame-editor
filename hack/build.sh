#!/bin/bash

set -e

base_dir="$(dirname "${BASH_SOURCE[0]}" | xargs realpath)/.."

bin_dir="${base_dir}/bin"

GOOS="${GOOS:-$(go env GOOS)}"
GOARCH="${GOARCH:-$(go env GOARCH)}"

GO_LD_FLAGS="${GO_LD_FLAGS:-"-w -s"} -X github.com/heathcliff26/infraspace-savegame-editor/pkg/cli.GitCommit=$(git rev-parse HEAD)"

if [ ! -z "${RELEASE_VERSION}" ]; then
    GO_LD_FLAGS="${GO_LD_FLAGS} -X github.com/heathcliff26/infraspace-savegame-editor/pkg/cli.Version=${RELEASE_VERSION}"
fi

output_name="${bin_dir}/infraspace-savegame-editor_${GOOS}_${GOARCH}"

if [ "${GOOS}" == "windows" ]; then
    output_name="${output_name}.exe"
fi

pushd "${base_dir}" >/dev/null

echo "Building $(basename "${output_name}")"
GOOS="${GOOS}" GOARCH="${GOARCH}" go build -ldflags="${GO_LD_FLAGS}" -o "${output_name}" ./cmd/save-editor/...
