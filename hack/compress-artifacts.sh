#!/bin/bash

set -e

base_dir="$(dirname "${BASH_SOURCE[0]}" | xargs realpath)/.."

bin_dir="${base_dir}/bin"

pushd "${bin_dir}"

for f in infraspace-savegame-editor*; do
    if [[ "${f}" == *.zip ]] || [[ "${f}" == *.tar.gz ]]; then
        continue
    fi
    echo "Compressing ${f}"
    if [[ "${f}" == *.exe ]]; then
        zip "${f%.exe}.zip" "${f}"
    else
        tar -czf "${f}.tar.gz" "${f}"
    fi
done

popd >/dev/null
