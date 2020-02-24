default: proto

build: .make.build
.make.build: Dockerfile docker-compose.yml
	docker-compose build protoc
	touch .make.build

proto: .make.proto
.make.proto: .make.build
	docker-compose run -v "$(PWD)":/protobuf --rm protoc -I=/protobuf -I=/go/src -I=/go/src/github.com/gogo/protobuf/protobuf --gogo_out=:/protobuf /protobuf/decimal/v1.proto
