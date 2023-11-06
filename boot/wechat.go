package boot

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client
var DepartRedis *redis.Client
var Cache *gcache.Cache

func init() {
	address, _ := g.Config("config.yaml").Get(context.Background(), "redis.default.address")
	defaultDB, _ := g.Config("config.yaml").Get(context.Background(), "redis.default.db")
	departDb, _ := g.Config("config.yaml").Get(context.Background(), "redis.default.departDb")
	pass, _ := g.Config("config.yaml").Get(context.Background(), "redis.default.pass")

	Redis = redis.NewClient(&redis.Options{
		Addr:     address.String(),
		Password: pass.String(),   // no password set
		DB:       defaultDB.Int(), // use default DB
	})

	DepartRedis = redis.NewClient(&redis.Options{
		Addr:     address.String(),
		Password: pass.String(),  // no password set
		DB:       departDb.Int(), // use default DB
	})

	//docker run -p 6379:6379 --name gfcq_redis -v D:\code\product\golang\src\gfcq_organize\manifest\redis\redis.conf:/etc/redis/redis.conf --restart=always --network gfcq -d redis redis-server /etc/redis/redis.conf
}
