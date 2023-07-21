# A simple GRPC service

# GoFrame Sites

### GoFrame Repo

* [https://github.com/lj1570693659/gfcq_organize.git](https://github.com/lj1570693659/gfcq_organize.git)

### GoFrame Home

* [https://github.com/lj1570693659](https://github.com/lj1570693659) 


### step1: 創建ETCD容器
* docker run -d -p 23791:2379 -p 23801:2380 appcelerator/etcd --listen-client-urls http://0.0.0.0:2379 --advertise-client-urls http://0.0.0.0:2379 

### step2: 啟動gfcq_organize容器
* docker build -t gfcq_organize_images -f .\manifest\docker\Dockerfile .
* docker run -itd -p 9090:9090  --name gfcq_organize  --net host  gfcq_organize_images
* docker run -itd -p 9090:9090  --name gfcq_organize  gfcq_organize_images



[//]: # (docker run -d -p 2379:2379 -p 2381:2381 --name etcd_local quay.io/coreos/etcd:v3.5.7 .\manifest\docker\etcd --data-dir=/etcd-data --name node1 --listen-client-urls http://10.24.12.84:2379 --advertise-client-urls http://10.24.12.84:2379 --listen-peer-urls http://10.24.12.84:2381 --initial-advertise-peer-urls http://10.24.12.84:2381 --initial-cluster node1=http://10.24.12.84:2381 --log-level info --logger zap --log-outputs stderr)
