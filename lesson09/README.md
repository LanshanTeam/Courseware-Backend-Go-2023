## 前言

俗话说，不会 linux 的后端不是好运维，后端开发人员也需要懂一点运维。

展望目前行业的发展，自动化运维和 DevOps 等理念的诞生和推广，其实也让运维和开发的边界变得更加的模糊，运维人员可能需要掌握一些开发语言来实现编写自动化运维的脚本，或者需要掌握一些开发的语言和工作流程以便更好地协作。开发人员可能也需要了解运维的内容以便更好地开发实际有用的软件以及更好地协作。

## Linux 基础

### 啥是linux?

> Linux是一套免费使用和自由传播的类Unix操作系统，是一个基于POSIX和UNIX的多用户、多任务、支持多线程和多CPU的操作系统。它能运行主要的UNIX工具软件、应用程序和网络协议。它支持32位和64位硬件。Linux继承了Unix以网络为核心的设计思想，是一个性能稳定的多用户网络操作系统。 Linux操作系统诞生于1991 年10 月5 日（这是第一次正式向外公布时间）。Linux存在着许多不同的Linux版本，但它们都使用了Linux内核。Linux可安装在各种计算机硬件设备中，比如手机、平板电脑、路由器、视频游戏控制台、台式计算机、大型机和超级计算机。 严格来讲，Linux这个词本身只表示**Linux内核**，但实际上人们已经习惯了用Linux来形容整个基于Linux内核，并且使用GNU 工程各种工具和数据库的操作系统。

全世界 99% 服务器的操作系统都是 Linux

### linux 的各个发行版

![img](https://pic4.zhimg.com/80/v2-c79ecff3534b9e135aac51673172a7cf_720w.webp)

Linux发行版主要有三个分支：Debian、Slackware、Redhat。
(1)Debian:（以社区的方式运作）

1. Ubuntu:基于Debian开发的开源Linux操作系统，主要针对桌面和服务器；
2. Linux Mint:基于Debian和Ubuntu的Linux发行版，致力于桌面系统对个人用户每天的工作更易用，更高效，且目标是提供一种更完整的即刻可用体验。

(2)slackware

1. suse：基于Slackware二次开发的一款Linux,主要用于商业桌面、服务器。
2. SLES(SUSE Linux Enterprise Server(SLES):企业服务器操作系统，是唯一与微软系统兼容的Linux操作系统。
3. OpenSUSE:由suse发展而来，旨在推进linux的广泛使用，主要用于桌面环境，用户界面非常华丽，而且性能良好。

(3)Redhat

1. rhel(red hat enterprise Linux):Red Hat公司发布的面向企业用户的Linux操作系统。早起版本主要用于桌面环境，免费：
2. Fedora:基于Red Hat Linux终止发行后，红帽公司计划以Fedora来取代Red Hat Linux在个人领域的应用，而另外发行的Red Hat Enterprise Linux取代Red Hat Linux在商业应用的领域。Fedora的功能对于用户而言，它是一套功能完备、更新快速的免费操作系统，而对赞助者Red Hat公司而言，它是许多新技术的测试平台，被认为可用的技术最终会加入到Red Hat Enterprise Linux中。Fedora大约每六个月发布新版本。
3. Centos:基于Red hat Linux提供的可自由使用源代码的企业级Linux发行版本。每个版本的Centos都会获得十年的支持（通过安全更新的方式）。新版本的Centos大约每两年发行一次，而每个版本的Centos会定期（大概6个月）更新一次，以支持新的硬件。这样，建立一个安全、低维护、稳定、高预测性、高重复性的Linux环境。

(4)其他发行版本：

1. Gentoo:基于linux的自由操作系统，基于Linux的自由操作系统，它能为几乎任何应用程序或需求自动地作出优化和定制。追求极限的配置、性能，以及顶尖的用户和开发者社区，都是Gentoo体验的标志特点， Gentoo的哲学是自由和选择。得益于一种称为Portage的技术，Gentoo能成为理想的安全服务器、开发工作站、专业桌面、游戏系统、嵌入式解决方案或者别的东西--你想让它成为什么，它就可以成为什么。由于它近乎无限的适应性，可把Gentoo称作元发行版。
2. Aech Linux(或称Arch):以轻量简洁为设计理念的Linux发行版。其开发团队秉承简洁、优雅和代码最小化的设计宗旨。

### linux 目录

根目录 `/` 是整个文件系统的起始点，但实际上，Linux 文件系统采用了一种层次结构来组织文件和目录。

根目录 `/` 下有许多子目录，这些子目录用于组织不同类型的文件和系统资源。以下是一些常见的子目录及其用途：

- `/bin`：包含可执行文件，用于存放基本的系统命令。
- `/boot`：包含启动所需的文件，例如内核和引导加载程序。
- `/etc`：包含系统配置文件。
- `/home`：用户的主目录，每个用户都有一个对应的子目录。
- `/lib`：包含共享库文件，用于支持系统和应用程序。
- `/opt`：用于安装可选软件包的目录。
- `/root`：超级用户（root）的主目录。
- `/sbin`：包含系统管理员使用的系统命令。
- `/usr`：包含用户安装的应用程序和文件，类似于程序文件夹。
- `/var`：包含可变的数据文件，例如日志文件和缓存文件。

### linux常用命令

##### 安装

###### **Ubuntu（apt）**

- 安装软件包

```text
sudo apt-get update  # 更新软件包列表
sudo apt-get install packageName  # 安装软件包
```

- 删除软件包

```text
sudo apt-get remove packageName  # 删除软件包（保留配置文件）
sudo apt-get purge packageName   # 删除软件包及其配置文件
```

- 搜索软件包

```text
sudo apt search packageName  # 搜索软件包
```

- 更新软件包列表

```text
sudo apt-get update  # 更新软件包列表
```

###### **CentOS（yum）**

- 安装软件包

```text
sudo yum update  # 更新软件包列表（yum）
sudo yum install packageName  # 安装软件包（yum）
```

- 删除软件包

```text
sudo yum remove packageName  # 删除软件包
```

- 搜索软件包

```text
sudo yum search packageName  # 搜索软件包（yum）
```

- 更新软件包列表

```text
sudo yum update  # 更新软件包列表（yum）
```

###### **Alpine（apk）**

- 安装软件包

```text
apk add packageName  # 安装软件包
```

- 删除软件包

```text
apk del packageName  # 删除软件包
```

- 搜索软件包

```text
apk search packageName  # 搜索软件包
```

- 更新软件包列表

```text
apk update  # 更新软件包列表
```



##### 说明书

```
man xxx
```



##### 查看当前目录下所有文件

```
ls
## 列出所有文件（包括隐藏文件）
ls -a
```



##### 切换目录

```
#绝对路径
cd /home/
#相对路径
cd GoProJ
```



##### 查看端口号占用情况

```
lsof -i:8080
```



##### 查看所有用户所有进程

```
ps -a
```



##### 强制杀死指定进程

课前问题：kill 命令是什么呢？

```
kill -9 12345
```



##### 查看文本文件内容

```
cat 1.txt
cat main.go
```



##### 修改文件权限

```
chmod 733 main
chmod +x main

#千万不要输入
chmod -Rf 777 /
#千万不要输入

#威力仅次于 rm -r -f /
```

[权限相关拓展](https://zhuanlan.zhihu.com/p/165091887)

##### 创建文件夹

```
mkdir goProJ
```



##### 创建文件

```
#touch命令
touch main.go
#vim工具
vim main.go
```



[vim入门到精通]:https://zhuanlan.zhihu.com/p/68111471 .

##### 删除文件

```
#有文件夹会一起删除
rm -rf main.go
```



##### 管道 ｜

竖线符号 `|` 是管道符号（Pipe Symbol）。该符号用于将一个命令的输出传递给另一个命令作为输入，以便二者之间进行数据传输和处理。

```
command1 | command2

ls | grep "pattern"
```



##### 常用快捷键

```
 #停止进程
 ctrl + c
 #清屏
 ctrl + l
```



[进阶]:https://www.lanqiao.cn/courses/1 .

### 部署

go 的程序部署非常简单（相比 Java 而言）

简单来说仅需两步：1，编译；2，运行

而如果我们想把可执行文件放到服务器上运行，还需要额外几个步骤

大致分为两种

1.编译为二进制文件，上传到服务器，服务器直接运行二进制文件

2.将源代码上传到服务器，服务器进行编译(前提是服务器配置好go环境)

#### 上传二进制文件到服务器

首先，在goland终端更改环境变量

```
go env -w GOOS=linux
```

然后，来到项目根目录，输入命令编译源代码

```
go build main.go
```

等待几秒，你会发现项目中多了一个main文件(注意不是main.exe)

这就是编译好的二进制文件，可以在任何linux上运行

我们可以通过以下命令，将此文件上传到服务器

```
#将当前目录下main文件拷贝到127.0.0.1的/home/路径下
scp main root@127.0.0.1:/home/
```

如果想要免密登录 配置SSH即可

1.在主机A创建密钥对

```
ssh-keygen #创建证书
```

然后均回车（选择默认）

2、将文件上传至免登录主机B的authorized_keys

```
/root/.ssh/authorized_keys
```

#### 上传源代码到服务器

通常会先将源码上传到github再拉下来，或者使用goland一键上传

[服务器go安装](https://blog.csdn.net/qq_43098070/article/details/126075629)

然后进入项目目录直接编译即可

```
go build main.go
```

#### 运行程序

现在服务器上有编译好的二进制文件了，那我们如何运行它？

只需要

```
./main
```

非常的简单

但是我们会发现无法运行，因为文件没有权限

这时候只需要

```
chmod +x main
chmod 733 main
```

给它高权限再次运行即可

这时候，使用云服务器的同学们需要去控制台找到安全组的配置规则，在入方向添加相应端口

##### 后台运行

如果按以上方式运行，当我们关闭窗口时会发现程序被终止了

那有什么办法可以让程序在后台运行呢？

我们可以使用 tmux，systemctl 或者 nohup+& 命令

###### nohup+&

```
nohup ./main &
```

###### tmux

[tmux](https://www.ruanyifeng.com/blog/2019/10/tmux.html)

###### systemctl

[systemctl设置自己的systemd.service服务设置守护进程](https://blog.csdn.net/lonnng2004/article/details/88964763)

[systemctl 详解](https://www.ruanyifeng.com/blog/2016/03/systemd-tutorial-commands.html)

### 作业

发送到 [wanziyu@lanshan.email](mailto:wanziyu@lanshan.email)

#### lv0

熟悉linux指令

#### lv1

部署课件中样例到服务器，浏览器访问页面并截图

没有服务器的同学在类 unix 操作系统下编译运行并在浏览器中访问页面并截图 