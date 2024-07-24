package dbconnector

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	redislist    map[string]*redis.Pool
	redis_struct map[string]*Redis_t
)

func init() {
	redislist = make(map[string]*redis.Pool)
	redis_struct = make(map[string]*Redis_t)
}

// 通过json字符串原型增加一个Redis连接器
func addRedisByJsonString(jsonstr string) error {
	p := &Redis_t{
		Key:           "",
		Server:        "",
		Port:          0,
		Pwd:           "",
		DB:            0,
		PoolMaxActive: 0,
		MaxIdle:       0,
		IdleTimeout:   0,
		MaxActive:     0,
	}
	err := json.Unmarshal([]byte(jsonstr), p)
	if err != nil {
		return fmt.Errorf("Redis连接数据JSON格式分析错误:%s", err.Error())
	}
	if p.Key == "" {
		return fmt.Errorf("Redis连接数据JSON格式分析错误:Key不能为空")
	}
	addr := fmt.Sprintf("%s:%d", p.Server, p.Port)
	redisPool := &redis.Pool{
		MaxIdle:     p.MaxIdle,                                  //2,
		IdleTimeout: time.Duration(p.IdleTimeout) * time.Second, //240 * time.Second,
		MaxActive:   p.PoolMaxActive,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr, redis.DialDatabase(p.DB))
			if err != nil {
				return nil, err
			}
			if p.Pwd != "" {
				if _, err := c.Do("AUTH", p.Pwd); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	redis_struct[p.Key] = p
	redislist[p.Key] = redisPool
	// err = sqldb.Ping()
	// if err != nil {
	// 	return err
	// }
	return err
}

func addRedisByStruct(r *Redis_t) error {
	addr := fmt.Sprintf("%s:%d", r.Server, r.Port)
	redisPool := &redis.Pool{
		MaxIdle:     2,
		IdleTimeout: 240 * time.Second,
		MaxActive:   1000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr, redis.DialDatabase(r.DB))
			if err != nil {
				return nil, err
			}
			if r.Pwd != "" {
				if _, err := c.Do("AUTH", r.Pwd); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	redis_struct[r.Key] = r
	redislist[r.Key] = redisPool
	// err = sqldb.Ping()
	// if err != nil {
	// 	return err
	// }
	return nil
}

// 清除所有的连接器
func CleanRedis() {
	for k := range redislist {
		delete(mariadbs, k)
	}
	for k := range redis_struct {
		delete(redis_struct, k)
	}
}
