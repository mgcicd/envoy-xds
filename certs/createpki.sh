#!/bin/bash

openssl genrsa -out ca.key 2048
openssl req -new -x509 -key ca.key -out ca.crt -days 36500

openssl genrsa -out server.key 2048
openssl req -new -nodes -key server.key -out server.csr -days 36500 -subj "/C=CN/OU=CompanyName/O=departmentName/CN=*.private-domain.com.cn" -config ./openssl.cnf -extensions v3_req
openssl x509 -req -days 36500 -in server.csr -out server.pem -CA ca.crt -CAkey ca.key -CAcreateserial -extfile ./openssl.cnf -extensions v3_req

openssl genrsa -out client.key 2048
openssl req -new -nodes -key client.key -out client.csr -days 36500 -subj "/C=CN/OU=CompanyName/O=departmentName/CN=*.private-domain.com.cn" -config ./openssl.cnf -extensions v3_req
openssl x509 -req -days 36500 -in client.csr -out client.pem -CA ca.crt -CAkey ca.key -CAcreateserial -extfile ./openssl.cnf -extensions v3_req

:<<EOF
openssl.cnf需修改: 
[ CA_default ]
...
copy_extensions = copy # 取消注释
...
[ req ]
...
req_extensions = v3_req # 取消注释
...
[v3_req]
...
subjectAltName = @alt_names # 新增配置
[ alt_names ] # 新增配置
DNS.1 = *.private-domain.com.cn # 新增配置
EOF