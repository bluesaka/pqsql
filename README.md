## PostgreSQL

### 创建数据库

```
create database test;
```

### 创建带参数的数据库

```
CREATE DATABASE database_name
WITH
   [OWNER =  role_name]
   [TEMPLATE = template]
   [ENCODING = encoding]
   [LC_COLLATE = collate]
   [LC_CTYPE = ctype]
   [TABLESPACE = tablespace_name]
   [ALLOW_CONNECTIONS = true | false]
   [CONNECTION LIMIT = max_concurrent_connection]
   [IS_TEMPLATE = true | false ]

如：
create database test
WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1;
```

### 创建表

```
CREATE TABLE "user"
(
    id bigint primary key,
    name varchar(20),
    created_at timestamp without time zone NOT NULL
)
```

### 表字段类型

### 表相关操作命令

```
alter table "user" alter column name type varchar(20);
alter table "user" rename column name to name2;
alter table "user" drop name;
alter table "user" rename to "user2";
create table user2 (like "user" including all);
drop table if exists user2;
insert into "user" (id,name,created_at) values(1,'a1','2020-01-01 00:00:01'),(2,'a2','2020-01-01 00:00:01');
insert into "user2" select * from "user";
delete from "user" where id = 1;
```

### 查询

```
select * from "user";
select id,name as 姓名 from "user" order by id desc, name asc;
select distinct id,name from "user";
select * from "user" where id > 1 and id in (2,3) and name like 'a%' and name like 'ab_';
select count(*) from "user";
select avg(xxx)、sum(xxx)、max(xxx)... from "user";
```

### 更新

```
update "user" set name='aa1' where id=1;
```

### PostgreSQL特有索引

- 索引类型：B-tree，Hash，GiST，SP-GiST，GIN
  
> 一般普通表使用 `B-tree`
  
- 全文检索的用 GIN，有两种：
  ```
  tsvector：表示一个被优化的，可以基于搜索的文档
  tsquery：表示一个文本查询，支持布尔操作 & | !
  两者搭配使用：select * from tableName where to_tsvector('zhparser', fieldName) @@ to_tsquery('搜索内容1|其他内容2');
  GIN 索引支持操作符：@>， ?，?&， ?|
  ```
  
### json 和 jsonb 类型

- PostgreSQL 的 JSON 存储有两种字段类型：json、jsonb
- jsonb 是二进制存储，json 是文本存储
- json 写入快，读取慢；jsonb 写入慢，读取快。
- jsonb 支持 GIN 索引类型，实际使用推荐 jsonb

```
例如字段 info 的json为 {"id":111, "name": "aaa"}
创建jsonb索引: CREATE INDEX idx_name ON tableName USING GIN(info);
查询：select* from tableName where jsonFieldName @> '{"name":"aaa"}'
```

### explain分析
```
explain (analyze,verbose,costs,buffers,timing)select * from sys_user where to_tsvector('zhparser', real_name) @@ to_tsquery('李四|admin');
```

### 中文分词、全文检索

```
PostgreSQL 全文索引的实现要靠 gin 索引(通用倒排索引)。
默认内置了英文、西班牙文等分词，但是没有中文。

可使用插件：zhparser
```

### 常用插件

- PostGis
  ```
  支持空间对象、空间索引、空间操作函数和空间操作符
  常见业务场景：
      圈地：以当前或者指定中心点，找出周边指定范围内的数据
      围栏：检测指定点落在哪个地理围栏中，比如共享单车是否停在禁停区
  ```

- TimescaleDB
  ```
  时序数据的自动分片、高效写入、检索、准实时聚合等。
  ```

### PostgreSQL与MySQL的语法区别

- 符号区别
  
> PostgreSQL 不支持有 ` 号包裹表名、字段名等情况。
  
- 命名区别
  
  > 创建索引的时候，MySQL只需同表中索引名称唯一即可，PostgreSQL需要整个库唯一。


### PostgreSQL和MySQL的适用场景

**PostgreSQL适用场景**

- 复杂查询和事务处理
  
> PostgreSQL在处理复杂查询和大量事务时较为出色，适用于需要高度数据完整性和一致性的应用程序，如金融系统或企业级应用。
  
- 地理信息系统（GIS）
  
> PostGIS是一个用于地理空间对象的开源空间数据库扩展，使得PostgreSQL可以很好地处理地理信息。
  
- 大数据和数据仓库
  
  > PostgreSQL的分区表、并行查询和流复制等功能，使其成为处理大数据和数据仓库方面的强力工具。

**MySQL适用场景**

- Web应用程序
  
> 由于其高性能和易用性，MySQL是许多Web应用程序的首选数据库，尤其是在有大量查询操作的场景。
  
- 小型应用程序
  
> 针对小型的应用程序，MySQL提供了一个稳定、经济高效的解决方案。
  
- 快速开发和迭代
  
  > MySQL具有较低的学习曲线，适用于快速开发和迭代的项目。

### PostgreSQL和MySQL性能对比

**PostgreSQL**
- 复杂查询和事务
  
> PostgreSQL在处理复杂查询和高并发事务时表现出色，这得益于其MVCC（多版本并发控制）实现。
  
- 扩展性
  
  > 具有良好的扩展性，可以通过添加额外的节点来提高性能。

**MySQL**
- 读取性能
  
> MySQL在处理大量读操作时表现出色，尤其是在使用索引时。
  
- 写入性能
  
  > MySQL的写入性能也很好，但在某些情况下会受到锁的影响。

#### PostgreSQL和MySQL的数据模型和功能

**PostgreSQL**
- 复杂数据类型
  
> 支持诸如数组、JSON、XML等复杂数据类型，使其适用于处理半结构化数据。
  
- 触发器和存储过程
  
  > 提供了强大的触发器和存储过程功能，使得可以在数据库层面实现复杂的业务逻辑。

**MySQL**
- 简单且高效的索引
  
> MySQL提供了简单且高效的索引机制，可以很好地支持大量的读操作。
  
- 存储引擎
  
  > MySQL支持多种存储引擎，例如InnoDB、MyISAM等。

### 

