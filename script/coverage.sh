go test -coverprofile='coverage.out' ./...
go tool cover -html='coverage.out' -o 'coverage.html'
if [[ "$OSTYPE" == "msys" ]]; then
  explorer coverage.html
else
	open coverage.html
fi
sleep 1
rm -rf coverage.out coverage.html
