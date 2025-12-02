package cache

import (
	"testing"
	"time"
)

func TestCachOp(t *testing.T){
	testData := []struct{
		key string
		val any
		expire time.Duration
	}{
		{"a",678,2*time.Second},
		{"b",false,2*time.Second},
		{"c",map[string]any{"a":1,"b":"b"},2*time.Second},
		{"d","string",2*time.Second},
		{"e","中文字符串",2*time.Second},
		{"f",678,2*time.Second},
	}

	c:= NewMemCache()
	c.SetMaxMemory("10MB")
	for _,item:=range testData{
		c.Set(item.key,item.val,item.expire)
		val,ok:=c.Get(item.key)
		if !ok{
			t.Error("缓存取值失败")
		}
		if item.key!="c"&&val!=item.val{
			t.Error("缓存取值数据与预期不一致")
		}
		_,ok = val.(map[string]any)
		if item.key=="c"&&!ok{
			t.Error("缓存取值数据与预期不一致")
		}
	}

	if int64(len(testData))!=c.Keys(){
		t.Error("缓存数量不一致")
	}
	time.Sleep(6*time.Second)
	if c.Keys()!=0{
		t.Error("缓存清除失败")
	}
}