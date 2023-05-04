#!/bin/bash

ca=`cat ./ca.crt | base64 | sed ':label;N;s/\n//g;b label'`

clientKey=`cat ./client.key | base64 | sed ':label;N;s/\n//g;b label'`
clientPem=`cat ./client.pem | base64 | sed ':label;N;s/\n//g;b label'`

cat << EOF > ./secrets/xds-client.yml
apiVersion: v1
kind: Secret
metadata:
  name: xds-client
  namespace: xxx
data:
  ca.crt: ${ca}
  client.pem: ${clientPem}
  client.key: ${clientKey}
EOF


serverKey=`cat ./server.key | base64 | sed ':label;N;s/\n//g;b label'`
serverPem=`cat ./server.pem | base64 | sed ':label;N;s/\n//g;b label'`
cat << EOF > ./secrets/xds-server.yml
apiVersion: v1
kind: Secret
metadata:
  name: xds-server
  namespace: xxx
data:
  ca.crt: ${ca}
  server.pem: ${serverPem}
  server.key: ${serverKey}
EOF

publicDomainKey=`cat ./public-domain.com.key | base64 | sed ':label;N;s/\n//g;b label'`
publicDomainCrt=`cat ./public-domain.com.crt | base64 | sed ':label;N;s/\n//g;b label'`
cat << EOF > ./secrets/public-server.yml
apiVersion: v1
kind: Secret
metadata:
  name: public-server
  namespace: xxx
data:
  public-domain.com.crt: ${publicDomainCrt}
  public-domain.com.key: ${publicDomainKey}
EOF