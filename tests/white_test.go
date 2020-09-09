package tests

import (
	"testing"

	"github.com/zzsds/zhongbang/pkg/zhongbang"
	proto_zhongbang "github.com/zzsds/zhongbang/pkg/zhongbang/proto"
)

var (
	zb = zhongbang.New(func(o *zhongbang.Zhongbang) {
		o.Host = "http://settle.kwfzhifu.com"
		o.Merchno = ""
	})
)

func TestOpenAcct(t *testing.T) {
	var msg = proto_zhongbang.OpenAcctMessage{
		TrueName: "jayden",
		Certno:   "500101199108221111",
		Mobile:   "18584565111",
		Cardno:   "622225754411255464",
		BankName: "众邦银行",
		Bankno:   "323521012066",
	}
	result, err := zb.OpenAcct(&msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result.Message)
}
