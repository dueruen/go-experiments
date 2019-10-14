svc_name:=user_service

compile: clean protoc

protoc:
	protoc -I .. ../proto/User.proto --go_out=plugins=grpc:gen

include ../base.mk