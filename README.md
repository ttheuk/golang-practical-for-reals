# Mô tả
Project gồm các tính năng: 
- Thêm & tìm kiếm dữ liệu trên PostgreSQL (sử dụng GORM, gRPC, Elasticsearch). 
- Export excel file (handle bằng RabbitMQ).

Project có 4 server:
- elastic (port 8081): 
  + Đánh index theo 3 trường name, age, id
  + Search theo trường name, age
- gin (port 8080): handle api, thêm dữ liệu vào database (đánh index khi thêm)
- rpc (port 5000): vận chuyển data từ gin đến server excel
- rabbitmq (port 5672): handle task export excel file

# Need
elasticsearch 7.7.0 <br>
github.com/olivere/elastic/v7 v7.0.15 <br>
rabbitmq-server-v3.8.4 <br>
google.golang.org/grpc v1.29.1 <br>
google.golang.org/protobuf v1.24.0 <br>

# Chạy server
`cd ./_app/cmd/server && go run .`
<br>
`cd ./_app/rpc_server && go run .`
<br>
`cd ./_elastic_/ && go run .`
<br>
`cd ./_excel/ && go run .`

# Test API
**Search** <br>
`GET http://localhost:8080/students?keyword=some_keyword`

**Create** <br>
`POST http://localhost:8080/students`
JSON
```
{
	"name": "Lữ Gia",
	"age": 21
}
```

**Export excel** <br>
`GET http://localhost:8080/students/excel?path=some_path&file-name=some_name`

