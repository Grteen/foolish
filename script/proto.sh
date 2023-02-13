cd /root/BE/idl

protoc --go_out=../grpc/userdemo/ user.proto

protoc --go-grpc_out=../grpc/userdemo/ user.proto

protoc --go_out=../grpc/articaldemo/ artical.proto

protoc --go-grpc_out=../grpc/articaldemo/ artical.proto

protoc --go_out=../grpc/searchdemo/ search.proto

protoc --go-grpc_out=../grpc/searchdemo/ search.proto

cd /root/BE/script/bin
./remove_tag