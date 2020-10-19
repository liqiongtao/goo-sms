package goo_sms

import (
	"fmt"
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
)

func TestGooAliyun_Send(t *testing.T) {
	code, err := NewAliyun(conf).Send("18510381580", "mob-login")
	fmt.Println(code, err)
}
