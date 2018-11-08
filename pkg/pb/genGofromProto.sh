#! /bin/bash
#1.nf.pb.go

cd  /root/notification/pkg/pb
rm notification.pb.go notification.pb.gw.go notification.swagger.json
#nf.swagger.go

cd  /root/notification/pkg/
protoc -I/usr/local/include -I. \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
--go_out=plugins=grpc:. \
pb/notification.proto

#2.nf.pb.gw.go
protoc -I/usr/local/include -I. \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
--grpc-gateway_out=logtostderr=true:. \
pb/notification.proto


#3.nf.swagger.json
protoc -I/usr/local/include -I. \
-I$GOPATH/src  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
--swagger_out=logtostderr=true:. \
pb/notification.proto



#将step3生成的swagger.json文件的值替换到swagger.go文件中的const值中去















