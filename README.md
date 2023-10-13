# k8s-tftp
[![build-ghcr](https://github.com/marcel2012/k8s-tftp/actions/workflows/docker-image.yml/badge.svg?branch=master)](https://github.com/marcel2012/k8s-tftp/actions/workflows/docker-image.yml)
[![build-dockerhub](https://github.com/marcel2012/k8s-tftp/actions/workflows/dockerhub-image.yml/badge.svg?branch=master)](https://github.com/marcel2012/k8s-tftp/actions/workflows/dockerhub-image.yml)
[![CodeQL](https://github.com/marcel2012/k8s-tftp/actions/workflows/github-code-scanning/codeql/badge.svg?branch=master)](https://github.com/marcel2012/k8s-tftp/actions/workflows/github-code-scanning/codeql)
### With GET and PUT support
Multiarch TFTP server which can run on Kubernetes.

Repository: https://hub.docker.com/r/marcel2012/k8s-tftp

Github: https://github.com/marcel2012/k8s-tftp

## How to use

### Docker 

```shell
docker run -p 69:69/udp -v ./volume:/tftpboot marcel2012/k8s-tftp:latest
```

Test

```shell
% tftp          
tftp> connect 127.0.0.1
tftp> put file.txt
tftp> get file.txt
```

### Docker Compose

Create `docker-compose.yaml` file

```yaml
version: '3.8'

services:
  app:
    image: marcel2012/k8s-tftp:latest
    ports:
      - "69:69/udp"
    volumes:
      - storage:/tftpboot

volumes:
  storage:
```

Run and test

```shell
% docker compose up -d
% tftp          
tftp> connect 127.0.0.1
tftp> put file.txt
tftp> get file.txt
```

### Kubernetes

Create a `tftp.yaml` like this

``` yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  namespace: pxe
  name: pxe-deployment
spec:
  selector:
    matchLabels:
      app: pxe
  template:
    metadata:
      labels:
        app: pxe
    spec:
      containers:
      - name: pxe
        image: marcel2012/k8s-tftp
        ports:
        - containerPort: 69
        volumeMounts:
          - name: nfs
            mountPath: /tftpboot
      volumes:
      - name: nfs
        persistentVolumeClaim:
          claimName: YOUR TFTP FILES PVC!!!
---
apiVersion: v1
kind: Service
metadata:
  name: pxe-deployment
  namespace: pxe
spec:
  externalIPs:
  - IP YOU WANT
  ports:
  - port: 69
    protocol: UDP
    targetPort: 69
  selector:
    app: pxe
  sessionAffinity: None
  type: LoadBalancer
```
