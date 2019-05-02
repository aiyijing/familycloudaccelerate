#!/bin/bash
# windows
env CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags '-s -w' -o ./bin/FamilySpeedUp_windows_386.exe
env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags '-s -w' -o ./bin/FamilySpeedUp_windows_amd64.exe
# macos
env CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build -ldflags '-s -w' -o ./bin/FamilySpeedUp_darwin_386
env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w' -o ./bin/FamilySpeedUp_darwin_amd64
# linux 386 and amd64
env CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags '-s -w' -o ./bin/FamilySpeedUp_linux_386
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o ./bin/FamilySpeedUp_linux_amd64
# 嵌入式平台
# linux_mips
env CGO_ENABLED=0 GOOS=linux GOARCH=mips GOMIPS=softfloat go build -ldflags '-s -w' -o ./bin/FamilySpeedUp_linux_mips
env CGO_ENABLED=0 GOOS=linux GOARCH=mips64 go build -ldflags '-s -w' -o ./bin/FamilySpeedUp_linux_mips64
# linux_mipsle
env CGO_ENABLED=0 GOOS=linux GOARCH=mipsle GOMIPS=softfloat go build -ldflags '-s -w' -o ./bin/FamilySpeedUp_linux_mipsle
env CGO_ENABLED=0 GOOS=linux GOARCH=mips64le go build -ldflags '-s -w' -o ./bin/FamilySpeedUp_linux_mipsle64
# linux_arm
env CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags '-s -w' -o ./bin/FamilySpeedUp_linux_arm
env CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags '-s -w' -o ./bin/FamilySpeedUp_linux_arm64
