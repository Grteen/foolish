cd /root/foolish/idl

protoc --go_out=../grpc/userdemo/ user.proto

protoc --go-grpc_out=../grpc/userdemo/ user.proto

protoc --go_out=../grpc/articaldemo/ artical.proto

protoc --go-grpc_out=../grpc/articaldemo/ artical.proto

protoc --go_out=../grpc/searchdemo/ search.proto

protoc --go-grpc_out=../grpc/searchdemo/ search.proto

protoc --go_out=../grpc/notifydemo/ notify.proto

protoc --go-grpc_out=../grpc/notifydemo/ notify.proto

protoc --go_out=../grpc/actiondemo/ action.proto

protoc --go-grpc_out=../grpc/actiondemo/ action.proto

protoc --go_out=../grpc/commentdemo/ comment.proto

protoc --go-grpc_out=../grpc/commentdemo/ comment.proto

protoc --go_out=../grpc/msmtpdemo/ smtp.proto

protoc --go-grpc_out=../grpc/msmtpdemo/ smtp.proto

cd /root/foolish/script/bin
./remove_tag