# ReadME

The User services

## Dependices

Use following command:

```sh
glide up
```

## ProtoBuf Generate

Use following command:

```sh
protoc -I proto/ --go_out=plugins=grpc:proto proto/service.proto
```

## Start Server

```sh
go run main.go
```
