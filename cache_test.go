package Figo

import (
	"log"
	"testing"
	"time"
)

func TestNewTimerCache(t *testing.T) {
	tc := NewTimerCache(10, func(key, value interface{}) {
		log.Println("expire @key:", key, "  @value:", value)
	})
	tc.Put("foo", "bar")
	tc.Put("hello", "world")
	log.Println(tc.Get("foo"))
	log.Println(tc.Get("hello"))
	time.Sleep(time.Duration(11) * time.Second)
}

func TestNewRedisCache(t *testing.T) {
	rc := NewRedisCache(RedisPool("localhost:6379", ""))
	rc.Put("foo", "bar")
	rc.Put("hello", "world")
	log.Println(TpString(rc.Get("foo")))
	log.Println(TpString(rc.Get("hello")))
}

func TestNewRedisTimerCache(t *testing.T) {
	rc := NewRedisTimerCache(RedisPool("localhost:6379", ""), 10)
	rc.Put("foo", "bar")
	rc.Put("hello", "world")
	log.Println(TpString(rc.Get("foo")))
	log.Println(TpString(rc.Get("hello")))
	time.Sleep(time.Duration(11) * time.Second)
	log.Println(TpString(rc.Get("foo")))
	log.Println(TpString(rc.Get("hello")))
}

func TestNewAsCache(t *testing.T) {
	ac := AsUtee.AsConnect("127.0.0.1:3000")
	type Val struct {
		Offset int64
	}
	cache := NewAsCache(ac, AsUtee.NewSetInfo("push", "test"), TypeOf(Val{}))
	cache.put("foo2016", &Val{88888})
	v := cache.get("foo2016")
	log.Println(v.(*Val))
}
