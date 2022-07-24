cd cmd/mgh
go build -o /usr/local/bin/mgh
cd ../../
go test -v ./e2e_tests
rm /usr/local/bin/mgh