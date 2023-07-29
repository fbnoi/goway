# goway

make:
protoc -I=. --go_out=./protobuf ./pb/*.proto