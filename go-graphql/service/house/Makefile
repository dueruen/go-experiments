svc_name:=house_service

compile: clean protoc

protoc:
	protoc -I ../.. ../../proto/House.proto --go_out=plugins=grpc:gen

include ../../base.mk