# test-ip

测试 Kubernetes 中真实 IP 是否可以获取。

## 功能

获取客户端真实 IP 和请求头信息，用于验证网关代理（如 Istio IngressGateway）是否正确传递客户端请求信息。

## API

### GET /test/ip

获取客户端真实 IP 地址。

```bash
curl -H "Host: ok.com" http://<ingress-gateway>/test/ip
```

响应示例：
```json
{
  "ip": "192.168.1.100"
}
```

### GET /test/header

获取请求头信息。

```bash
curl -H "Host: ok.com" http://<ingress-gateway>/test/header
```

响应示例：
```json
{
  "Accept": "*/*",
  "Host": "ok.com",
  "X-Real-IP": "192.168.1.100",
  ...
}
```

## 部署

### 1. 构建镜像

```bash
docker build -t test-ip:latest .
```

### 2. 部署应用

```bash
kubectl apply -f deployment.yaml
```

这会创建 Deployment 和 Service 资源。

### 3. 部署 Istio 网关配置

```bash
kubectl apply -f gateway.yaml
```

这会创建 Gateway 和 VirtualService 资源，将 `ok.com` 域名的流量路由到你的 test-ip 服务。

### 4. 访问服务

通过 Istio IngressGateway 域名 `ok.com` 访问服务。

## Istio 配置

- **Gateway**: `test-gateway` - 监听 80 端口，接收 `ok.com` 域名流量
- **VirtualService**: `test-vs` - 将流量路由到 `test-ip` 服务的 8080 端口
