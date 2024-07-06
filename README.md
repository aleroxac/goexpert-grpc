# goexpert-grpc


## Conceitos
### gRPC
- É um protocolo de remote procedure call de alta performance e open source desenvolvido pelo Google.

### Protocol Buffers
- Formato de dados usado para serializar estrutura de dados; os dados trafegados via gRPC são nesse formato

### HTTP/2
- Versão mais recente do protocolo HTTP, que usa multiplexação(as requisições e respostas são paralelas e assíncronas), server push(estáticos, por exemplo), headers comprimidos e formatados em HPACK

### Tipos de chamadas
```
- API unary:                    client -> request(1) -> data  <- response(1) <- server
- API server streaming:         client -> request(1) -> data  <- response(N) <- server
- API client streaming:         client -> request(N) -> data  <- response(1) <- server
- API bi directional streaming: client -> request(N) -> data  <- response(N) <- server
```

### REST vs gRPC
| Característica | REST           | gRPC                        |
| -------------- | -------------- | --------------------------- |
| Formato        | Texto / JSON   | Protocol Buffers            |
| Relacionamento | Uniderecional  | Bi-direcional e Assíncrono  |
| Latência       | Alta latência  | Baixa latência              |
| Contrato       | Sem contrato   | Tem contrato                |
| Streaming      | Sem suporte    | Tem suporte                 |
| Design         | Pré-definido   | Livre                       |
| Bibliotecas    | Terceiros      | Nativa                      |


## Setup
``` shell
wget -O /tmp/protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v26.1/protoc-26.1-linux-x86_64.zip
unzip /tmp/protoc.zip -d ~/.local/
sudo ln -s ~/.local/bin/protoc /usr/local/bin/protoc

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

wget -O /tmp/evans.tar.gz https://github.com/ktr0731/evans/releases/download/v0.10.11/evans_linux_amd64.tar.gz
tar -xzvf /tmp/evans.tar.gz
sudo mv evans /usr/local/bin
```

## Como trabalhar com gRPC no go
``` shell
### 1. Crie uma [pasta com os arquivos proto para cada uma de suas entidades](proto/course_category.proto)

### 2. Gere os binários via protoc
protoc --go_out=. --go-grpc_out=. proto/course_category.proto

### 3. Implemente os [services](internal/service/category.go) com base nas [interfaces auto-geradas pelo protoc](internal/pb/course_category_grpc.pb.go) 'not implemented'

### 5. Suba o server
go run cmd/grpcServer/main.go

### 6. Use evans para chamar os services do app
evans --proto proto/course_category.proto --host localhost --port 50051
=> call CreateCategory
=> name1
=> desc1
```



## Como rodar o projeto
``` shell
### 1. Crie a tabela
sqlite3 db.sqlite 'create table categories (id string, name string, description string);'

### 5. Suba o server
go run cmd/grpcServer/main.go

### 6. Use evans para chamar os services do app
evans --proto proto/course_category.proto --host localhost --port 50051
=> call CreateCategory
=> name1
=> desc1
```
