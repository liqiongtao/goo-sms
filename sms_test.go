package goo_sms

import (
	"fmt"
	"github.com/liqiongtao/goo"
	"testing"
)

var (
	conf = AliyunConfig{
		Region:       "beijing",
		Appid:        "",
		Secret:       "",
		SignName:     "",
		TemplateCode: "",
	}

	redisConfig = goo.RedisConfig{
		Addr:     "127.0.0.1:6379",
		Password: "test",
		DB:       9,
		Prefix:   "t:",
	}
)

func TestGooAliyun_Send(t *testing.T) {
	goo.RedisInit(redisConfig)
	InitCache(goo.Redis())

	code, err := NewAliyun(conf).Send("18510381580", "mob-login")
	fmt.Println(code, err)

	err = NewAliyun(conf).Verify("18510381580", "mob-login", code)
	fmt.Println(err)
}
