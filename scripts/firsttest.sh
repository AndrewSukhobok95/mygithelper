mkdir tmp
mkdir -p ./tmp/repo-1
mkdir -p ./tmp/repo-2

cd ./tmp/repo-1
git init
touch file.txt
echo hi >> file.txt
cd ./../../

cd ./tmp/repo-2
git init
touch file.txt
echo hi >> file.txt
cd ./../../

cd ./cmd/mgh
go build
mv mgh ./../../tmp/
cd ./../../tmp

./mgh -m addall
./mgh -m commit




