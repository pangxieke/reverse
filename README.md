# Api网关
简易Api 开发平台网关，用来记录接口调用日志。

## 目标

- 记录接口调用次数
- 访问调用、响应日志
- 鉴权
- 统一网关入口
- 代理分流

## 功能
### 日志
记录请求日志，响应日志。使用json格式，可以结合elasticsearch等工具，快速查询响应

### 访问量统计
通过Header中app_key等，记录请求日志到Mongo中，方便统计调用日志。用于计费等。
