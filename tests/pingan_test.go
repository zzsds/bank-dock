package tests

import (
	"log"
	"testing"

	"github.com/zzsds/bank-dock/pkg/pingan"
	proto_pingan "github.com/zzsds/bank-dock/pkg/pingan/proto"
)

// QYdm 交易所代码
// SupAcctiD 资金主账号ID
// TranwebName 交易所名称
// SFTP 银行SFTP 地址
// PORT  SFTP 端口
// User SFTP 用户账号
// Password SFTP 密码
// ReconPath 上传清算文件SFTP 路径

var (
	pa = pingan.New(func(o *pingan.Options) {
		o.Host = "47.114.118.106:7072"
		o.QYdm = "9293"
		o.SupAcctiD = "15000139168587"
		o.TranwebName = "贵州汇通盛世商品交易中心（提配）"
		o.Sftp = pingan.Sftp{
			Sftp:      "183.62.75.68",
			Port:      22,
			User:      "pab2biuser",
			Password:  "0_1@32&*s+udd~QhfsVfEW4*5<)",
			ReconPath: "/afe_ftp/ZJJG/9293/",
		}
	})
)

// TestContract ...
func TestContract(t *testing.T) {
	req := proto_pingan.ContractRequest1303{}
	rsp, err := pa.Contract(req)
	if err != nil {
		log.Fatal(err)
	}
	t.Log(rsp)
}

// TestContract ...
func TestCashOut(t *testing.T) {
	req := proto_pingan.CashOutRequest1318{
		TranWebName:       "TranWebName",
		ThirdCustId:       "ThirdCustId",
		IdType:            "IdType",
		IdCode:            "IdCode",
		TranOutType:       "TranOutType",
		CustAcctId:        "CustAcctId",
		CustName:          "CustName",
		SupAcctId:         "SupAcctId",
		TranType:          "TranType",
		OutAcctId:         "OutAcctId",
		OutAcctIdName:     "OutAcctIdName",
		OutAcctIdBankName: "OutAcctIdBankName",
		OutAcctIdBankCode: "OutAcctIdBankCode",
		Address:           "Address",
		CcyCode:           "CcyCode",
		TranAmount:        "TranAmount",
		FeeOutCustId:      "FeeOutCustId",
	}
	rsp, err := pa.CashOut(req)
	if err != nil {
		log.Fatal(err)
	}
	t.Log(rsp)
}

// TestContract ...
func TestNotifyParse(t *testing.T) {
	var name string
	err := pa.NotifyParse([]byte("jayden"), &name)
	if err != nil {
		log.Fatal(err)
	}
	t.Log(name)
}
