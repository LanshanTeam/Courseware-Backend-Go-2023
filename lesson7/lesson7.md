

# MySQL

![img](https://camo.githubusercontent.com/0d321bc853789843d8659f9d694de2f1d5149d5cba6b792a114824ec98e2ebf5/68747470733a2f2f73312e617831782e636f6d2f323032322f31312f31372f7a6d6e5556732e706e67)

## 概念篇

### 什么是数据库？

数据库（Database）是指存储数据的容器，也被称为数据存储库（Data Store）。数据库能够存储大量结构化和非结构化的数据，包括文本、数字、图像、音频等各种类型的数据。

根据数据存储方式和结构的不同，数据库可以分为关系型数据库(如MySQL、Oracle、SQL Server、PostgreSQL)和非关系型数据库(如MongoDB、Cassandra、Redis、Elasticsearch)。

### MySQL介绍

> MySQL是一个**关系型数据库**，一种开源关系数据库管理系统（RDBMS），它使用最常用的数据库管理语言-结构化查询语言（SQL）进行数据库管理

关系型数据库的数据都以**数据表**的形式进行存储和管理。

表是一种二维行列表格，关系型数据库中的数据通常以表的形式存储，如图所示

![img](https://camo.githubusercontent.com/e04f85a224ede31a3e08285692578a8f99b5b5b80d550fbf9e0fa2b6eaadebfe/68747470733a2f2f73312e617831782e636f6d2f323032322f31312f31392f7a4b39394a672e706e67)

- `列（column）` - 表中的一个字段 ，如图中`patient_id` 。所有表都是由一个或多个列组成的

- `行（row）` - 表中的一个记录，比如第一行，这条记录，记录下了小红的信息(病人号、姓名、年龄、身高体重等等)

- `主键（primary key）` - 一列（或一组列），其值能够唯一标识表中每一行。 

  > 比如图2中每个patient_id，都标识了一个**独一无二**的病人
  >
  > 又比如b站的uid，也标识了一个**独一无二**的用户

#### SQL

> SQL 是一种**操作数据库**的语言，包括创建数据库、删除数据库、查询记录、修改记录、添加字段等。
>
> SQL 是关系型数据库的标准语言，所有的关系型数据库管理系统（RDBMS）都将 SQL 作为其标准处理语言。

SQL有以下用途：

- 允许用户访问关系型数据库系统中的数据；(**查数据**)

- 允许用户描述数据；(**自定义要查哪些数据**)

- 允许用户定义数据库中的数据，并处理该数据；(**建表**)

- 允许将 SQL 模块、库或者预处理器嵌入到其它编程语言中；(**在Go中使用SQL**)

- 允许用户创建和删除数据库、表、数据项（记录）；

- 允许用户在数据库中创建视图、存储过程、函数；

- 允许用户设置对表、存储过程和视图的权限；

  **SQL命令分为三类**

  1.DDL - Data Definition Language，数据定义语言

  > 对数据的结构和形式进行定义，一般用于数据库和表的创建、删除、修改等。

  | 命令   | 说明                                               |
  | ------ | -------------------------------------------------- |
  | CREATE | 用于在数据库中创建一个新表、一个视图或者其它对象。 |
  | ALTER  | 用于修改现有的数据库，比如表、记录。               |
  | DROP   | 用于删除整个表、视图或者数据库中的其它对象         |

  2.DML - Data Manipulation Language，数据处理语言

  > 对数据库中的数据进行处理，一般用于数据项（记录）的插入、删除、修改和查询。

  | 命令   | 说明                                 |
  | ------ | ------------------------------------ |
  | SELECT | 用于从一个或者多个表中检索某些记录。 |
  | INSERT | 插入一条记录。                       |
  | UPDATE | 修改记录。                           |
  | DELETE | 删除记录。                           |

  3.DCL - Data Control Language，数据控制语言

  | 命令            | 说明                                       |
  | --------------- | ------------------------------------------ |
  | COMMIT          | 提交                                       |
  | SAVEPOINT       | 保存点                                     |
  | ROLLBACK        | 回滚                                       |
  | SET TRANSACTION | 设置当前事务的特性，它对后面的事务没有影响 |

#### 数据类型

MySQL 支持多种数据类型，大致可以分为三类：数值、日期/时间和字符串(字符)类型。

##### 数值类型

| 类型            | 字节大小                                |
| --------------- | --------------------------------------- |
| TINYINT         | 1                                       |
| SMALLINT        | 2                                       |
| MEDIUMINT       | 3                                       |
| INT(INTEGER)    | 4                                       |
| BIGINT          | 8                                       |
| FLOAT           | 4                                       |
| DOUBLE          | 8                                       |
| DECIMAL（小数） | 对DECIMAL(M,D)中 M>D 则为M+2，否则为D+2 |

##### 日期和时间类型

| **类型**  | **字节大小** | **格式**            | 用途                     |
| --------- | ------------ | ------------------- | ------------------------ |
| DATE      | 3            | YYYY-MM-DD          | 日期值                   |
| TIME      | 3            | HH:MM:SS            | 时间值或持续时间         |
| YEAR      | 1            | YYYY                | 年份值                   |
| DATETIME  | 8            | YYYY-MM-DD hh:mm:ss | 混合日期和时间值         |
| TIMESTAMP | 4            | YYYY-MM-DD hh:mm:ss | 混合日期和时间值，时间戳 |

##### 字符串类型

| 类型        | 字节大小     | 用途                         |
| ----------- | ------------ | ---------------------------- |
| **CHAR**    | 0-255        | 定长字符串                   |
| **VARCHAR** | 0-65535      | 变长字符串                   |
| TINYBLOB    | 0-255        | <=255个字符的二进制字符串    |
| TINYTEXT    | 0-255        | 短文本字符串                 |
| BLOB        | 0-65535      | 二进制形式的长文本数据       |
| TEXT        | 0-65535      | 长文本数据                   |
| MEDIUMBLOB  | 0-16777215   | 二进制形式的中等长度文本数据 |
| MEDIUMTEXT  | 0-16777215   | 中等长度文本数据             |
| LONGBLOB    | 0-4294967295 | 二进制形式的极大文本数据     |
| LONGTEXT    | 0-4294967295 | 极大文本数据                 |

## 实践篇

**在开始实践前 请先安装好MySQL和配置好环境变量。**

### SQL语法

使用MySQL前当然要学会使用SQL语句。

我们先来了解一下SQL的基本规则

1. **SQL 语句要以分号`;`结尾**，在 RDBMS （关系型数据库）当中，SQL 语句是逐条执行的，一条 SQL 语句代表着数据库的一个操作
2. **SQL 语句不区分大小写**，例如，不管写成 SELECT 还是 select，解释都是一样的。表名和列名也是如此，但是，**插入到表中的数据是区分大小写的**
3. **SQL 语句的单词之间必须使用半角空格（英文空格）或换行符来进行分隔**

**登录MySQL**

windows系统打开cmd或windows powershell；mac和linux打开终端

输入

```
mysql -u root -p
```

root是你的用户名

然后根据提示输入密码来启动MySQL

下面来一些 SQL 语句的示例

```sql
creat database student;
//创建数据库

drop database school;
//删除数据库

show databases;
//查看所有数据库

show databases like 'school';
//查看名为 "school" 的数据库

show databases like '%sch%';
//查看名字中包含sch的数据库(%表示任意个字符)

show databases like 'sch%';
//查看以sch开头的数据库

use school
//进入数据库"school"

CREATE TABLE `student` (
		`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
		`name` VARCHAR(20) DEFAULT '0',
		`sex` VARCHAR(8) DEFAULT '',
		`age` INT(11) ,
		PRIMARY KEY(`id`)
	)ENGINE=InnoDB AUTO_INCREMENT=1 CHARSET=utf8mb4;
//创建表"student",ENGINE=InnoDB:指定 MySQL引擎为InnoDB,AUTO_INCREMENT=1:自增的字段每次自增1,CHARSET=utf8mb4：指定编码为utf8base64以支持中文

drop table student;
//删除表"student"
```

接下来我们以创建好的 student 表为例，对数据进行操作

**插入 Insert**

> 由于id是自增，所以添加的时候无需插入id

```
insert into 表名(列1，列2，列3，列4，...)  values(值，值，值)
```

```sql
insert into student(name,sex,age) values('小明','男',18);
```

当字段允许空时或者设置了默认值时，也可以不插入这个字段，字段会默认为空或者默认值

```sql
insert into student(name,sex) values('小红','女');
```

**查询 Select**

```
select 列1,列2 from 表名 where ...;
```

查询所有学生

```sql
select *from student;
```

查询所有学生的姓名与年龄

```sql
select name,age from student;
```

查询男学生

```sql
select name,age from student where sex='男';
```

如果想改变输出时的字段名 可以使用**AS**

```sql
select name as '姓名' from student;
```

**子查询**

顾名思义，子查询嵌套在另一个查询语句里

```sql
select name,age from student where name in(
	select name from student where sex='男'
)order by id desc;
```

我们将这条语句拆解一下

最外层：

```sql
select name,age from student where name in(..)order by id desc;
```

里层:

```sql
select name from student where sex='男';
```

外层在里层结果的基础上再进行查询

最后的:

```sql
order by id desc;
```

意思是结果按id排序，降序(desc)

升序是`asc`

**更新 Update**

```
update <表名> set 列1=值1，列2=值2，.... where...
```

```sql
update student set age=19,sex='女' where id=1;
```

这里用`where`来限定要更新的记录,因为update会**更新所有**符合限定条件的记录，如果不限定，会更新所有记录

条件之间可以用`and`、`or`连接

```sql
update student set age=19,sex='女' where id>=1 and name='小明';
```

**删除 Delete**

```sql
delete from student where id=1;
```

**修改数据表**

修改字段

```
alter table <表名> change <旧字段名> <新字段名> <新数据类型>；
```

```
alter table student change name 姓名 varchar(32);
```

删除字段

```
alter table student drop 姓名;
```

添加字段

> 需要注意的是，添加的字段会默认在表最后一列

```
alter table <表名> add <新字段名><数据类型>[约束条件];
```

```
alter table student add name varchar(12) default '李华' not null;
```

在开头添加字段

> 末尾加上first即可

```
alter table student add name varchar(12) default '李华' not null first;
```

在中间添加字段

> 末尾加after接字段名即可

```
alter table student add name varchar(12) default '李华' not null after id;
```

### Go操作MySQL

上述内容是直接使用SQL语句操作MySQL，下面我们学习使用Golang来操作数据库。

Go官方为我们提供了 **database/sql** 包保证SQL或类SQL数据库的泛用接口，**sqlx** 是基于标准库的扩展，在开发中 **sqlx** 包使用更加广泛，下面我们以使用 **sqlx** 包为例。

首先下载 MySQL 驱动

```go
go get -u github.com/go-sql-driver/mysqlsqlx 
```

然后安装 sqlx

```go
go get github.com/jmoiron/sqlx
```

#### 连接数据库

```go
package main

import (
    _ "github.com/go-sql-driver/mysql" //要导入相应驱动包，否则会报错
	"github.com/jmoiron/sqlxl"					   
	"fmt"
)

// 定义一个全局对象db
var db *sqlx.DB

func initDB() {
	var err error 
    
	dsn := "username:password@tcp(127.0.0.1:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
	
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	
    fmt.Println("connecting to MySQL...")
	return
}

func main() {
	//初始化连接
	initDB()
}
```

#### CRUD

##### 查询

查询单行数据

```go
func queryRowDemo() {
	sqlStr := "select id, name, age from user where id=?"
	var u user
	err := db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
}
```

查询多行数据

```go
func queryMultiRowDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	var users []user
	err := db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}
```

##### 插入、更新和删除

插入、更新和删除均使用 **Exec** 方法

```go
// 插入数据
func insertRowDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, "沙河小王子", 19)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
func updateRowDemo() {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 39, 6)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 6)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

```

##### 事务

事务详细介绍：[一文读懂数据库事务_数据库事务是什么-CSDN博客](https://blog.csdn.net/K346K346/article/details/114085663)

简单来说，事务（Transaction）指一个操作，由多个步骤组成，要么全部成功，要么全部失败。

下面的代码演示了一个简单的事务操作，该事物操作能够确保两次更新操作要么同时成功要么同时失败，不会存在中间状态。

```go
func transactionDemo2()(err error) {
	tx, err := db.Beginx() // 开启事务
	if err != nil {
		fmt.Printf("begin trans failed, err:%v\n", err)
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			fmt.Println("rollback")
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
			fmt.Println("commit")
		}
	}()

	sqlStr1 := "Update user set age=20 where id=?"

	rs, err := tx.Exec(sqlStr1, 1)
	if err!= nil{
		return err
	}
	n, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}
	sqlStr2 := "Update user set age=50 where i=?"
	rs, err = tx.Exec(sqlStr2, 5)
	if err!=nil{
		return err
	}
	n, err = rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}
	return err
}
```

以上是 sqlx 包最基础的用法介绍，它更强大的功能就留给大家自行探索。

这里贴一下李文周老师的 sqlx 教程：[sqlx库使用指南 | 李文周的博客 (liwenzhou.com)](https://www.liwenzhou.com/posts/Go/sqlx/)

李文周老师介绍了 sqlx 的更多用法，上述示例代码也有一部分是直接从里面copy的 :)

#### Gorm介绍

- Gorm的使用并不是我们这节课的重点，所以这里就不介绍Gorm的具体用法。

##### 什么是Gorm

GORM 是 Go 语言中的一个强大的对象关系映射（ORM）库，它提供了一种简单且优雅的方式来进行数据库操作。ORM 是一种编程技术，它将数据库表与程序中的对象进行映射，使得开发者可以使用面向对象的方式进行数据库的增删改查操作，而无需直接编写 SQL 语句。（关于什么是ORM，大家可以在课下阅读相关文章了解）

##### 为什么我们需要Gorm

ORM的作用是在关系型数据库和对象之间作一个映射，这样，我们在具体的操作数据库的时候，就不需要再去和复杂的SQL语句打交道，只要像平时操作对象一样操作它就可以了 。

使用 ORM带来的好处有

- **减少我们重复的写sql语句的实践，提高我们的工作效率，降低开发成本**
- **访问数据的时候，可以使用抽象的方式，用起来非常简洁**

以使用sqlx和Gorm进行查询进行对比：

```go
func queryRowDemo() {	
	var u user
    err := db.First(&user, 10)
    // SELECT * FROM user WHERE id = 10;
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
}
```

```go
func queryRowDemo() {
	sqlStr := "select * from user where id=?"
	var u user
	err := db.Get(&u, sqlStr,10)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
}
```

这里看起来我们写出的代码并差不了几行，但在一个正式的项目开发中 我们如果使用 sqlx 进行开发 我们会大量编写sql 语句 此时的时间成本就会远高于使用 Gorm 进行开发。

##### Gorm学习

Gorm官方有详尽的中文文档:[GORM 指南 | GORM - The fantastic ORM library for Golang, aims to be developer friendly.](https://gorm.io/zh_CN/docs/)

阅读官方文档是我们提高能力的有效方法🥰

## 作业

发送到邮箱nanach0@qq.com

提交格式：第七次作业-2022214012-罗一尧-Lv？

**截止时间**：下一次上课之前

### Lv0

在MySQL中创建一个database，在里面创建一张数据表student，然后使用go语言操作MySQL，向里面插入十条记录，然后全部读出并打印

### Lv1

在上次作业的基础上，使用MySQL进行数据持久化(不要使用gorm，sqlx跟databse/sql均可)

### Lv2

学习MySQL与SQL理论知识（无需提交）

### Lv3

如果你学有余力，使用gin框架+MySQL实现QQ的好友功能

- 登录与创建账号
- 加好友
- 删好友
- 查看所有好友
- 好友分组
- 好友搜索

你可能需要知道的知识点:

- 模糊查询



