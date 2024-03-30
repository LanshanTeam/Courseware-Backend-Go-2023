# 第十六节课

### 负载均衡

---

负载均衡（Load Balance，简称 LB）是高并发、高可用系统必不可少的关键组件，目标是 尽力将网络流量平均分发到多个服务器上，以提高系统整体的响应速度和可用性。

**负载均衡的主要作用如下：**

**高并发：**负载均衡通过算法调整负载，尽力均匀的分配应用集群中各节点的工作量，以此提高应用集群的并发处理能力（吞吐量）。

**伸缩性：**添加或减少服务器数量，然后由负载均衡进行分发控制。这使得应用集群具备伸缩性。

**高可用：**负载均衡器可以监控候选服务器，当服务器不可用时，自动跳过，将请求分发给可用的服务器。这使得应用集群具备高可用的特性。



### 反向代理

---

反向代理（Reverse Proxy）方式是指以 代理服务器 来接受网络请求，然后 将请求转发给内网中的服务器，并将从内网中的服务器上得到的结果返回给网络请求的客户端。

与反向代理相对的就是正向代理。

正向代理：发生在 **客户端**，是由用户主动发起的。vpn软件就是典型的正向代理，客户端通过主动访问代理服务器，让代理服务器获得需要的外网数据，然后转发回客户端。

反向代理：发生在 **服务端**，用户不知道代理的存在。

nginx就是反向代理的一个主流产品。



### 限流 熔断 降级

---

在分布式系统中，**熔断**、**限流**和**服务降级**是关键的容错设计模式，用于保护系统免受突发流量、故障或性能问题的影响。

让我们深入了解这些概念：

1. **限流**：

   - 当系统的处理能力无法应对外部请求的突发流量时，为了避免系统崩溃，必须采取限流措施。
   - 限流指标可以是：
     - **TPS**（每秒事务数）：按事务完成数量来限制流量。
     - **HPS**（每秒请求数）：服务端每秒收到的客户端请求数量。
     - **QPS**（每秒查询请求数）：服务端每秒响应的客户端查询数量。
   - 常用的限流方法包括流量计数器、滑动时间窗口、漏桶算法和令牌桶算法。

2. **熔断**：

   - 当流量过大或下游服务出现问题时，熔断会自动断开与下游服务的交互，以避免影响整个系统。
   - 熔断可以自我恢复，通过自我诊断下游系统是否已修复或上游流量是否减少至正常水平来判断。

3. **服务降级**：

   - 降级是从系统内部的平级服务或业务维度考虑，用于保护其他正常使用的功能。

   - 当服务出现问题或影响核心流程性能时，需要暂时屏蔽掉，待高峰或问题解决后再打开。

     

> 3者的关系这样理解：
> 拿下棋比喻：
> 限流： 相当于尽量避免同时和两三个人同时下
> 熔断：相当于你的一颗卒被围死了，就不要利用其它棋去救它了，弃卒保帅，否则救他的棋也可能被拖死
> 降级：相当于尽量不要走用处不大的棋了，浪费走棋机会（资源），使已经过河的棋有更多的走棋机会（资源）发挥最大作用



kitex中使用 [负载均衡](https://www.cloudwego.io/zh/docs/kitex/tutorials/service-governance/loadbalance/) [熔断](https://www.cloudwego.io/zh/docs/kitex/tutorials/service-governance/circuitbreaker/) [限流](https://www.cloudwego.io/zh/docs/kitex/tutorials/service-governance/limiting/)



**Sentinel** 是一款面向分布式服务架构的轻量级流量控制产品 

Sentinel 的主要特点：

1. **流量控制**：Sentinel 可以限制外部请求的流量，防止系统因突发流量而崩溃。
2. **实时监控**：它提供实时监控功能，您可以在控制台中查看接入应用的运行情况。
3. **开源生态**：Sentinel 可以与其他开源框架/库（如 Spring Cloud、Dubbo、gRPC）整合。

[sentinel文档](https://sentinelguard.io/zh-cn/docs/golang/quick-start.html)



### Nginx

---

Nginx是一个高性能、开源的HTTP服务器和反向代理服务器。它可以作为一个独立的Web服务器，也可以作为其他Web服务器的反向代理。Nginx以其出色的性能和高度的可扩展性而闻名，是许多高流量网站的首选服务器。

Nginx的特点包括：

- 高性能：Nginx采用事件驱动的方式处理请求，可以处理大量的并发请求。
- 可扩展性：Nginx可以通过添加模块来扩展其功能，支持各种第三方模块。
- 反向代理：Nginx可以作为反向代理服务器，可以隐藏真实服务器的地址，从而提高网络安全性。
- 负载均衡：Nginx可以实现负载均衡，可以将请求分配到多个服务器上，从而提高网站的性能和可靠性。

Nginx是一个功能强大而灵活的服务器，适合用于构建高性能、可扩展的Web应用程序。

在Nginx中实现反向代理非常简单。只需要在Nginx配置文件中添加一个`proxy_pass`指令即可。例如，要将所有`/api`请求代理到`http://localhost:3000`，可以使用以下配置：

```
location /api {
    proxy_pass <http://localhost:3000>;
}
```



这将使Nginx将所有`/api`请求代理到`http://localhost:3000`。在这个例子中，Nginx充当了客户端和服务器之间的中转站，将所有请求代理到真实的服务器上，然后将响应返回给客户端。

除了`proxy_pass`指令之外，Nginx还提供了许多其他的反向代理指令，包括：

- `proxy_set_header`：设置HTTP请求头。
- `proxy_redirect`：重定向代理请求。
- `proxy_cache`：缓存代理请求的响应。
- `proxy_connect_timeout`：设置代理连接超时时间。

通过使用这些指令，可以轻松地配置Nginx作为反向代理服务器，并提供高性能、可扩展的Web应用程序。



### Traefik

---

Traefik是一款现代的、基于容器的反向代理和负载均衡器。它可以自动发现容器、自动配置路由和HTTPS证书，并提供了许多其他的高级功能，如动态配置、健康检查和访问日志记录等。

Traefik的特点包括：

- 自动化：Traefik可以自动发现容器，并自动配置路由和HTTPS证书。它还可以自动调整流量，将请求分配到最少的容器上，并在容器出现问题时自动将流量转移到健康的容器上。
- 动态配置：Traefik支持动态配置，可以通过REST API或文件变更自动配置路由和负载均衡策略。它还支持多种后端服务，如Docker、Kubernetes、Mesos、Consul等。
- 高可用性：Traefik可以实现高可用性，可以将请求分配到多个Traefik实例上，从而实现负载均衡和故障转移。
- 访问日志记录：Traefik可以记录访问日志，并将其发送到各种日志记录服务中。

Traefik的配置文件非常简单，可以使用以下格式：

```
# traefik.toml
[entryPoints]
  [entryPoints.http]
  address = ":80"

[api]
  dashboard = true

[docker]
  endpoint = "unix:///var/run/docker.sock"
  domain = "example.com"
```



这将启用Traefik的API和仪表板，并配置Traefik使用Docker作为后端服务。

通过使用Traefik，可以轻松地实现反向代理和负载均衡，并提供高可用性、动态配置和访问日志记录等高级功能。

[Traefik文档](https://www.traefik.tech/)



### Kong

---

Kong是一个高性能的API网关，它可以管理微服务架构中的API通信。

- **核心组件**：Kong的主要组件包括Kong Server（基于nginx的服务器，用来接收API请求）、数据库（Apache Cassandra或PostgreSQL，用来存储操作数据）以及Kong dashboard（官方推荐的UI管理工具）。
- **插件系统**：Kong通过插件来扩展其功能，这些插件可以在API请求响应循环的生命周期中被执行。插件使用Lua编写，目前已有多个基础功能插件，如HTTP基本认证、密钥认证、CORS、TCP、UDP、文件日志、API请求限流、请求转发以及Nginx监控。
- **可扩展性**：Kong可以通过简单地添加更多的服务器来进行横向扩展，这意味着它可以在较低负载的情况下处理任何请求。
- **模块化**：Kong的模块化设计允许通过添加新的插件进行扩展，这些插件可以通过RESTful Admin API轻松配置。
- **跨基础架构运行**：Kong可以在云或内部网络环境中部署，包括单个或多个数据中心设置，以及public、private或invite-only APIs。
- **集群和高可用性**：Kong支持集群部署，集群中的节点通过gossip协议自动发现其他节点。当通过一个Kong节点的管理API进行变更时，也会通知其他节点。

[Kong文档](https://docs.konghq.com/)



### Nginx.conf

---

删除注释长这样

```
worker_processes  1;


events {
    worker_connections  1024;
}


http {
    include       mime.types;
    default_type  application/octet-stream;

    sendfile        on;
    
    keepalive_timeout  65;

    server {
        listen       80;
        server_name  localhost;

        location / {
            root   html;
            index  index.html index.htm;
        }


        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   html;
        }
    }
}
```

结构

```
...              #全局块

events {         #events块
   ...
}

http      #http块
{
    ...   #http全局块
    server        #server块
    { 
        ...       #server全局块
        location [PATTERN]   #location块
        {
            ...
        }
        location [PATTERN] 
        {
            ...
        }
    }
    server
    {
      ...
    }
    ...     #http全局块
}
```

1、**全局块**：配置影响nginx全局的指令。一般有运行nginx服务器的用户组，nginx进程pid存放路径，日志存放路径，配置文件引入，允许生成worker process数等。

2、**events块**：配置影响nginx服务器或与用户的网络连接。有每个进程的最大连接数，选取哪种事件驱动模型处理连接请求，是否允许同时接受多个网路连接，开启多个网络连接序列化等。

3、**http块**：可以嵌套多个server，配置代理，缓存，日志定义等绝大多数功能和第三方模块的配置。如文件引入，mime-type定义，日志自定义，是否使用sendfile传输文件，连接超时时间，单连接请求数等。

4、**server块**：配置虚拟主机的相关参数，一个http中可以有多个server。

5、**location块**：配置请求的路由，以及各种页面的处理情况。



### 作业

lv1:尝试在自己的项目中实现负载均衡，限流或者熔断

lv2:在docker中用nginx模拟实现负载均衡，涉及到网络部分可自行学习docker network

