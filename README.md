# http2mq 

提供HTTP的文件上传服务，然后转发到Kafka。文件内容的格式采用行协议，每一行代表一条消息。

# Features
  
文件支持gzip压缩

# Example

```
get latest message:
curl -H 'Authorization: Basic dXNlcjE6cHdkMQ==' 'http://127.0.0.1:8080/kafka/http2mq'

send message to http2mq
curl -H 'Authorization: Basic dXNlcjE6cHdkMQ==' -d '2' 'http://127.0.0.1:8080/kafka/http2mq'
```

# Install

```
go get -u github.com/WALL-E/http2mq

go run main.go

```

# Configure

```
web:
  port: 8080  #http listen port


kafka:
  brokers: 192.168.1.182:9092
  topic: http2mq
  consumer_user: http2mq_unit_test
  zk: 192.168.1.169:2181,192.168.1.179:2181,192.168.1.180:2181
  zk_root: /qianbao/kafka/logtest

users:
    - user1:pwd1 # username:password,  basic auth
    - user2:pwd2
    - user3:pwd3

topics:
    - ipos
    - http2mq
```
