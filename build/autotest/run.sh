#!/usr/bin/env bash

set -e
set -o pipefail
#set -o verbose
#set -o xtrace

# os: ubuntu18.04 x64

sedfix=""
if [ "$(uname)" == "Darwin" ]; then
    sedfix=".bak"
fi

## get turingchain path
TURINGCHAIN_PATH=$(go list -f "{{.Dir}}" github.com/turingchain2020/turingchain)

function build_auto_test() {

    trap "rm -f ../autotest/main.go" INT TERM EXIT
    local AutoTestMain="${TURINGCHAIN_PATH}/cmd/autotest/main.go"
    cp "${AutoTestMain}" ./
    sed -i $sedfix '/^package/a import _ \"github.com\/turingchain2020\/plugin\/plugin\"' main.go
    go build -v -i -o autotest
}

function copyAutoTestConfig() {

    declare -a TuringchainAutoTestDirs=("${TURINGCHAIN_PATH}/system" "../../plugin")
    echo "#copy auto test config to path \"$1\""
    local AutoTestConfigFile="$1/autotest.toml"

    #pre config auto test
    {

        echo 'cliCmd="./turingchain-cli"'
        echo "checkTimeout=60"
    } >"${AutoTestConfigFile}"

    #copy all the dapp test case config file
    for rootDir in "${TuringchainAutoTestDirs[@]}"; do

        if [[ ! -d ${rootDir} ]]; then
            continue
        fi

        testDirArr=$(find "${rootDir}" -type d -name autotest)

        for autotest in ${testDirArr}; do

            dapp=$(basename "$(dirname "${autotest}")")
            dappConfig=${autotest}/${dapp}.toml

            #make sure dapp have auto test config
            if [ -e "${dappConfig}" ]; then

                cp "${dappConfig}" "$1"/

                #add dapp test case config
                {
                    echo "[[TestCaseFile]]"
                    echo "dapp=\"$dapp\""
                    echo "filename=\"$dapp.toml\""
                } >>"${AutoTestConfigFile}"

            fi

        done
    done
}

function copyTuringchain() {

    echo "# copy turingchain bin to path \"$1\", make sure build turingchain"
    cp ../turingchain ../turingchain-cli ../turingchain.toml "$1"
    cp "${TURINGCHAIN_PATH}"/cmd/turingchain/turingchain.test.toml "$1"
}

function copyAll() {

    dir="$1"
    #check dir exist
    if [[ ! -d ${dir} ]]; then
        mkdir "${dir}"
    fi
    cp autotest "${dir}"
    copyAutoTestConfig "${dir}"
    copyTuringchain "${dir}"
    echo "# all copy have done!"
}

function main() {

    if [[ $1 == "build" ]]; then #build autotest
        build_auto_test
    else
        dir="$1"
        echo "$dir"
        rm -rf ../autotest/"$dir" && mkdir "$dir"
        cp -r "$TURINGCHAIN_PATH"/build/autotest/"$dir"/* ./"$dir"/ && copyAll "$dir"
        chmod -R 755 "$dir" && cd "$dir" && ./autotest.sh "${@:2}" && cd ../
    fi
}

main "$@"
