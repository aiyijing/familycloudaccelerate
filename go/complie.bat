:: 依赖gox 请使用go get github.com/mitchellh/gox
gox -build-toolchain
:: linux 386 and amd64
gox -os="linux" -arch="386"  -output="./bin/FamilySpeedUp_linux_386"
gox -os="linux" -arch="amd64"  -output="./bin/FamilySpeedUp_linux_amd64"
:: linux_mips
gox -os="linux" -arch="mips"  -output="./bin/FamilySpeedUp_linux_mips"
gox -os="linux" -arch="mips64"  -output="./bin/FamilySpeedUp_linux_mips64"
:: linux_mipsle
gox -os="linux" -arch="mipsle"  -output="./bin/FamilySpeedUp_linux_mipsle"
gox -os="linux" -arch="mips64le"  -output="./bin/FamilySpeedUp_linux_mips64le"
:: linux_arm
gox -os="linux" -arch="mipsle"  -output="./bin/FamilySpeedUp_linux_arm"
gox -os="linux" -arch="mips64le"  -output="./bin/FamilySpeedUp_linux_arm64"
:: windows
gox -os="windows" -arch="386"  -output="./bin/FamilySpeedUp_windows_386"
gox -os="windows" -arch="amd64"  -output="./bin/FamilySpeedUp_windows_amd64"