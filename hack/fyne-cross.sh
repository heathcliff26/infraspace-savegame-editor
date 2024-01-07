#!/bin/bash

set -e

base_dir="$(dirname "${BASH_SOURCE[0]}" | xargs realpath)/.."

os="${1}"
arches="${2:-$(go env GOARCH)}"

fyne_cross="$(go env GOPATH)/bin/fyne-cross"

if [ ! -f "${fyne_cross}" ]; then
    go install github.com/fyne-io/fyne-cross@latest
fi

pushd "${base_dir}" >/dev/null

${fyne_cross} "${os}" -arch="${arches}" ./cmd/save-editor/

IFS=',' read -ra arch_array <<<"${arches}"

for arch in "${arch_array[@]}"; do
    if [ "${os}" == "linux" ]; then
        mv "fyne-cross/bin/linux-${arch}/save-editor" "fyne-cross/bin/linux-${arch}/infraspace-savegame-editor"
        rm -rf "fyne-cross/dist/linux-${arch}"
        tar -C "fyne-cross/bin/linux-${arch}" -czf "fyne-cross/dist/infraspace-savegame-editor_linux-${arch}.tar.gz" infraspace-savegame-editor
    elif [ "${os}" == "windows" ]; then
        mv "fyne-cross/dist/windows-${arch}/InfraSpace Savegame Editor.exe.zip" "fyne-cross/dist/infraspace-savegame-editor_windows-${arch}.zip"
        rm -rf "fyne-cross/dist/windows-${arch}"
    fi
done

popd >/dev/null
