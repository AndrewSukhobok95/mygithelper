go build -o /usr/local/bin/mgh
go test -v ./e2e_tests
rm /usr/local/bin/mgh