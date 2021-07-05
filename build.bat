go env -w CGO_ENABLED=0
go build -o turingchain.exe
go build -o turingchain-cli.exe github.com/turingchain2020/plugin/cli
