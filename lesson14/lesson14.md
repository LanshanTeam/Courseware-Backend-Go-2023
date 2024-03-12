# Go-zero 快速搭建微服务demo

## 前言

我认为框架的学习，快速搭建微服务demo，看官方文档自己摸索足矣，主要讲一些踩坑经验和进阶的学习

## 官方文档（必看）！！！

直接看快速开始，跟着做就可以了

https://go-zero.dev/

## 社区群（里面大佬多，不懂的可以去问一下）

https://go-zero.dev/docs/reference/about-us

## 项目推荐

如果快速开始过了之后，能搭建一个小demo，那么就可以看大佬们的项目深入学习啦，这里我推荐两个，一个是looklook，可以学习一些规范和很多技术栈，包括链路追踪用的jeager。日志收集然后可视化呈现的elk，云原生监控prometheus，k8s..... 再一个就是电商项目lebron，这个对业务的学习是很有帮助的，尤其是秒杀那一块，能学到很多高并发处理手段。

### lookook(规范&&技术栈)

> 项目地址：https://github.com/Mikaelemmmm/go-zero-looklook
>
> LookLook文档地址：https://github.com/Mikaelemmmm/go-zero-looklook/tree/main/doc/chinese

### lebron（业务处理&&高并发设计）

> 项目地址：https://github.com/zhoushuguang/lebron
>
> 系列博客：https://learnku.com/articles/68472

## 一些视频教学

looklook作者讲解的入门教学视频，可以看一下

[go-zero入门视频](https://www.bilibili.com/video/BV1LS4y1U72n/?spm_id_from=333.788&vd_source=23d8be6eab81d2d14531293e81d6a5dc)

go-zero分布式缓存设计，这个万总讲的很好，**必看**

[go-zero分布式缓存设计](https://www.bilibili.com/video/BV1Rv411L72P/?spm_id_from=333.337.search-card.all.click&vd_source=23d8be6eab81d2d14531293e81d6a5dc)

## 一些tips

### 并发神器mr包

可以用mr.Finish()去并发rpc请求，相当于封装好的goroutine

还有MapReduce，也值得学习一些，下面有讲解文章可以看看

https://juejin.cn/post/7031483035820556295

视频推荐：https://www.bilibili.com/video/BV1eQ4y1W7PV?p=4&vd_source=23d8be6eab81d2d14531293e81d6a5dc

这个佬的视频都可以看，**很硬核很低层**

### go-zero实现的分布式锁

[官方文档demo](https://go-zero.dev/docs/tutorials/redis/redis-lock)

博客解读源码实现

https://juejin.cn/post/7026544325030838303

视频推荐：https://www.bilibili.com/video/BV1Pm4y1b76u?p=1&vd_source=23d8be6eab81d2d14531293e81d6a5dc

这里有介绍etcd实现的，很值得看看

## 作业

* Lv1 用go-zero随便写个服务demo
* Lv2 把looklook和lebron过一遍
* LvX 自己深入go-zero的一些底层设计实现源码、按照自己的兴趣广泛深入发展，重点是**深入**......
