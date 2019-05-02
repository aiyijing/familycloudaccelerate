#!/bin/bash

rm -fr ./bin
mkdir ./bin
bash ./complie.sh

if [ $? -ne 0 ]; then
    echo "complie error"
    exit 1
fi

all_os='linux windows darwin'
all_arch='386 amd64 arm arm64 mips mips64 mipsle mips64le'

rm -fr ./packages
mkdir ./packages

for os in $all_os; do
    for arch in $all_arch; do
        file_name="FamilySpeedUp_${os}_${arch}"
        file_path="./bin/${file_name}"
        
        if [ "${os}" == "windows" ]; then
            file_path="${file_path}.exe"
        fi
        if [ ! -f "${file_path}" ];then
            continue
        fi
        echo ${file_path}
        mkdir ${file_name}
        mv ${file_path} ${file_name}
        cp -r config.json ${file_name}

        if [ "${os}" == "windows" ]; then
            zip -rq "./packages/${file_name}.zip" ${file_name}
        else
            tar -zcf "./packages/${file_name}.tar.gz" ${file_name}
        fi
        rm -fr ${file_name}
    done
done
