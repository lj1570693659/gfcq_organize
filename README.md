# gfcq_product
重庆赣锋项目管理部后端代码仓库

## 生成PB文件
*  protoc -I .\manifest\protobuf --go_out=.\manifest\pb\ .\manifest\protobuf\user\v1\user.proto 