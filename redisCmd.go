package dbconnector

import (
	"errors"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// 从连连接池获取一个连接. 获得一个有效连接后，可以使用封装的常用命令，也可以直接使用 conn.Do原生函数执行任何命令
func (r RedisConnectors) Connect() (*redis.Conn, error) {
	pool, ok := redislist[string(r)]
	if !ok {
		return nil, fmt.Errorf("redis[%s] not existing", string(r))
	}
	redisClient := pool.Get()
	return &redisClient, nil
}

// 断开连接
func Diconnect(conn *redis.Conn) (err error) {
	if conn == nil {
		return errors.New("connection is nil")
	}
	err = (*conn).Close()
	return
}

// 常用命令GET
func GET(conn *redis.Conn, key string) (string, error) {
	if conn == nil {
		return "", errors.New("connection is nil")
	}
	val, err := redis.String((*conn).Do("GET", key))
	return val, err
}

func SET(conn *redis.Conn, key, value string) (err error) {
	if conn == nil {
		return errors.New("connection is nil")
	}
	_, err = (*conn).Do("SET", key, value)
	return err
}

func DEL(conn *redis.Conn, keys ...interface{}) (err error) {
	if conn == nil {
		return errors.New("connection is nil")
	}
	_, err = (*conn).Do("DEL", keys...)
	return err
}

func KEYS(conn *redis.Conn, query string) (keys []string, err error) {
	if conn == nil {
		return nil, errors.New("connection is nil")
	}
	keys, err = redis.Strings((*conn).Do("KEYS", query))
	return keys, err
}

func HMSET(conn *redis.Conn, params ...interface{}) (err error) {
	if conn == nil {
		return errors.New("connection is nil")
	}
	_, err = (*conn).Do("HMSET", params...)
	return err
}

func HMGET(conn *redis.Conn, params ...interface{}) (vals []string, err error) {
	if conn == nil {
		return nil, errors.New("connection is nil")
	}
	vals, err = redis.Strings((*conn).Do("HMGET", params...))
	return vals, err
}

func HGETALL(conn *redis.Conn, key string) (ret map[string]string, err error) {
	if conn == nil {
		return nil, errors.New("connection is nil")
	}
	ret, err = redis.StringMap((*conn).Do("HGETALL", key))
	return ret, err
}

func HDEL(conn *redis.Conn, params ...interface{}) (err error) {
	if conn == nil {
		return errors.New("connection is nil")
	}
	_, err = (*conn).Do("HDEL", params...)
	return err
}

func EXISTS(conn *redis.Conn, key string) (ret bool, err error) {
	if conn == nil {
		return false, errors.New("connection is nil")
	}
	res, err := redis.Int((*conn).Do("EXISTS", key))
	if err != nil {
		return false, err
	}
	return res != 0, nil
}

func BGREWRITEAOF(conn *redis.Conn) {
	if conn == nil {
		return
	}
	(*conn).Do("BGREWRITEAOF")
}
