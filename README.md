# Mô tả
Demo tính năng thêm và tìm kiếm dữ liệu trên PostgreSQL. Sử dụng GORM, gRPC và Elasticsearch

Project gồm 2 server:
- elastic (port 8081): 
  + Đánh index theo 3 trường name, age, id
  + Search theo trường name, age
- gin (port 8080): handle api, thêm dữ liệu vào database (đánh index khi thêm)

# Need
elasticsearch 7.7.0 <br>
github.com/olivere/elastic/v7 v7.0.15 <br>
google.golang.org/grpc v1.29.1 <br>
google.golang.org/protobuf v1.24.0 <br>

# Chạy server
`cd ./server && go run main.go`
<br>
`cd ./elastic && go run .`

