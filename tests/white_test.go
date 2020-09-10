package tests

import (
	"strconv"
	"testing"
	"time"

	"github.com/zzsds/bank-dock/pkg/zhongbang"
	proto_zhongbang "github.com/zzsds/bank-dock/pkg/zhongbang/proto"
)

var (
	zb = zhongbang.New(func(o *zhongbang.Zhongbang) {
		// o.Host = "http://settle.kwfzhifu.com"
		o.Host = "http://settle.test.kwfzhifu.com"
		// 商户号
		o.Merchno = "886331041120001"
		// 商家虚拟账户号（白名单接口无用）
		o.Cardno = "886033082400000078"
		// 交易密钥
		o.Secret = "9C473E2C81982E05C0AAAC7B436ABC69"
		// 代付密钥
		o.AgentSecret = "E857F4B8751309F258D331F31C784C78"
	})
)

// 新增白名单
func TestOpenAcct(t *testing.T) {
	var msg = proto_zhongbang.OpenAcctRequest{
		TrueName: "jayden",
		Certno:   "500101195108221116",
		Mobile:   "18584565116",
		Cardno:   "500101195108221116",
		BankName: "中国银行",
		Bankno:   "104100000004",
	}
	result, err := zb.OpenAcct(&msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result.Message)
}

// 修改白名单
func TestOpenUpdate(t *testing.T) {
	var msg = proto_zhongbang.OpenAcctRequest{
		TrueName: "frank",
		Certno:   "500101195108221116",
		Mobile:   "18584565116",
		Cardno:   "500101195108221117",
		BankName: "招商银行",
		Bankno:   "323521012066",
	}
	result, err := zb.OpenUpdate(&msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result.Message)
}

// 查询白名单
func TestOpenQuery(t *testing.T) {
	var msg = proto_zhongbang.OpenQueryRequest{
		Certno: "500101195108221116",
	}
	result, err := zb.OpenQuery(&msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result.GetMessage(), result.GetTrueName())
}

// 实时代付接口
func TestVirtPay(t *testing.T) {
	var msg = proto_zhongbang.VirtPayRequest{
		// 数据来源 2-接口提现（默认）
		PayType: "2",
		// 资金来源  1-余额提现（默认） 2-垫资提现
		PayMode: "1",
		// 交易方式 1-单笔
		SubType: "1",
		// 通知地址
		NotifyUrl: "http://xxx/zb/notify",
		// 商户订单号 商家自定义的流水号，不能重复
		Traceno: strconv.Itoa(int(time.Now().UnixNano())),
		// 代付金额 以元为单位，带2位小数
		Amount: "1.00",
		// 结算账号
		Accountno: "500101195108221116",
		// 结算户名
		AccountName: "frank",
		// 手机号码 可以填固定值“13800000000”
		Mobile: "18584565116",
		// 身份证号 可以填固定值18个0
		Certno: "500101195108221116",
		// 联行号 人行规定的12位的数字
		Bankno: "104626180805",
		// 支行名称 需要精确到分行或者支行
		BankName: "中国银行广西百色分行",
		// 银行类型 例如：中国工商银行
		BankType: "中国银行",
		// 备注信息 默认填写“实时提现”
		Remark: "实时提现",
	}
	result, err := zb.VirtPay(&msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result.GetMessage())
}
