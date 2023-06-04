sh ./script/protoc.sh
sh ./script/wire.sh
sh ./script/mockery.sh
golangci-lint run
sh ./script/coverage.sh