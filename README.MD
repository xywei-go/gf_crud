# GoFrame Template For SingleRepo

Project Makefile Commands: 
- `make cli`: Install or Update to the latest GoFrame CLI tool.
- `make dao`: Generate go files for `Entity/DAO/DO` according to the configuration file from `hack` folder.
- `make service`: Parse `logic` folder to generate interface go files into `service` folder.
- `make image TAG=xxx`: Run `docker build` to build image according `manifest/docker`.
- `make image.push TAG=xxx`: Run `docker build` and `docker push` to build and push image according `manifest/docker`.
- `make deploy TAG=xxx`: Run `kustomize build` to build and deploy deployment to kubernetes server group according `manifest/deploy`.

# 概述
实现员工-部门之间的增删改查操作

# 需求
1. 一个部门有多个员工，一个员工属于一个部门，通过部门带出员工
2. 通过部门查询所有员工
3. 对部门进行增删查改
4. 对员工进行增删查改
5. 通过员工查询也带出部门
6. 数据库使用MySQL
7. 部门员工统计使用Redis缓存
8. 新增员工，新增部门，修改员工部门后，使用RabbitMQ或者Kafka进行缓存更新
9. 

