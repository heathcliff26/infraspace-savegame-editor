#!/bin/bash

base_dir="$(dirname "${BASH_SOURCE[0]}" | xargs realpath)"

APP_ID="io.github.heathcliff26.infraspace-savegame-editor"
BINARY="infraspace-savegame-editor"

bin_dir="$HOME/.local/bin"
if [ "$(whoami)" == "root" ]; then
    bin_dir="/usr/local/bin"
fi

help() {
    echo "Integrate InfraSpace Savegame Editor with common desktop environments."
    echo
    echo "Usage: -i | --install    -- install desktop file"
    echo "       -u | --uninstall  -- uninstall desktop file"
    echo "       -h | --help       -- show usage"
}

install() {
    echo "Installing binary to ${bin_dir}/${BINARY}"
    cp "${base_dir}/${BINARY}" "${bin_dir}/${BINARY}"

    echo "Installing desktop file"
    xdg-desktop-menu install "${base_dir}/${APP_ID}.desktop"

    echo "Installing icon"
    xdg-icon-resource install --size 512 "${base_dir}/${APP_ID}.png"

    xdg-desktop-menu forceupdate
}

uninstall() {
    echo "Removing binary"
    rm "${bin_dir}/${BINARY}"

    echo "Removing desktop file and icon"
    xdg-desktop-menu uninstall "${APP_ID}.desktop"
    xdg-icon-resource uninstall --size 512 "${APP_ID}.png"
}

while [[ "$#" -gt 0 ]]; do
    case $1 in
    -i | --install)
        install
        exit 0
        ;;
    -u | --uninstall)
        uninstall
        exit 0
        ;;
    -h | --help)
        help
        exit 0
        ;;
    *)
        echo "Unknown argument: $1"
        help
        exit 1
        ;;
    esac
    shift
done

help
