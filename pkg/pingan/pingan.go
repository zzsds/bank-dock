package pingan

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	"github.com/zzsds/bank-dock/pkg"
	proto_pingan "github.com/zzsds/bank-dock/pkg/pingan/proto"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// Service ...
type Service interface {
	GetThirdLogNo() string
	Contract(proto_pingan.ContractRequest1303) (*proto_pingan.ContractResponse1303, error)
	NotifyParse([]byte, interface{}) error
	CashOut(proto_pingan.CashOutRequest1318) (*proto_pingan.CashOutResponse1318, error)
}

// Pingan ...
type Pingan struct {
	options Options
	conn    net.Conn
}

// Sftp 银行 sftp
type Sftp struct {
	// SFTP 银行SFTP 地址
	Sftp string
	// SFTP 端口
	Port int
	// SFTP 用户账号
	User string
	// SFTP 密码
	Password string
	// 上传清算文件SFTP 路径
	ReconPath string
}

// Options ...
type Options struct {
	// 外置机地址
	Host string
	// 交易所代码
	QYdm string
	// 资金主账号ID
	SupAcctiD string
	// 交易所名称
	TranwebName string
	// 银行 sftp
	Sftp       Sftp
	Simplified encoding.Encoding
}

// Option ...
type Option func(*Options)

// New ...
func New(opt ...Option) Service {
	// 设置默认值
	var opts = Options{
		Simplified: simplifiedchinese.GBK,
		// WriteDead:
	}
	for _, o := range opt {
		o(&opts)
	}
	addr, err := net.ResolveTCPAddr("tcp4", opts.Host)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	conn.SetKeepAlive(true)
	conn.SetNoDelay(true)
	conn.SetWriteBuffer(4096)
	return &Pingan{
		options: opts,
		conn:    conn,
	}
}

// NotifyParse ...
func (h *Pingan) NotifyParse(b []byte, rsp interface{}) error {

	return json.Unmarshal(b, &rsp)
}

// Contract 签约
func (h *Pingan) Contract(req proto_pingan.ContractRequest1303) (*proto_pingan.ContractResponse1303, error) {
	tranFunc := "1303"
	thirdLogNoget := getThirdLogNo()
	headerMsg := h.getHeader(req.XXX_Size(), proto_pingan.HeaderMessage{
		TranFunc:   tranFunc,
		ThirdLogNo: thirdLogNoget,
	})
	fmt.Println(headerMsg + "123&12321")

	return &proto_pingan.ContractResponse1303{
		ThirdLogNo: thirdLogNoget,
		Teserve:    "",
	}, nil
}

// CashOut 出金
func (h *Pingan) CashOut(req proto_pingan.CashOutRequest1318) (*proto_pingan.CashOutResponse1318, error) {
	tranFunc := "1318"
	var buff bytes.Buffer
	req.SupAcctId = h.options.SupAcctiD
	u := pkg.NewParmas(req).URL()
	thirdLogNoget := getThirdLogNo()
	headerMsg := h.getHeader(len(u), proto_pingan.HeaderMessage{
		TranFunc:   tranFunc,
		ThirdLogNo: thirdLogNoget,
	})
	buff.WriteString(headerMsg)
	h.conn.SetWriteDeadline(time.Now().Add(time.Second * 3))
	h.conn.SetReadDeadline(time.Now().Add(time.Second * 5))

	b, err := Transform([]byte(u), h.options.Simplified.NewEncoder())
	if err != nil {
		return nil, err
	}
	buff.Write(b)
	h.conn.Write(buff.Bytes())
	readbyte := make([]byte, 4096)
	_, err = h.conn.Read(readbyte)
	if err != nil {
		return nil, err
	}
	var rsp proto_pingan.CashOutResponse1318
	if err = h.NotifyParse(readbyte, &rsp); err != nil {
		return nil, err
	}

	return &rsp, nil
}

// GetThirdLogNo ...
func (h *Pingan) GetThirdLogNo() string {
	return getThirdLogNo()
}

// GetHeader 获取header 头信息
func (h *Pingan) getHeader(length int, data proto_pingan.HeaderMessage) string {
	data.Qydm = h.options.QYdm
	dateFormat := time.Now().Format("20060102150405")
	length1 := strconv.Itoa(length + 122)
	length2 := strconv.Itoa(length)

	newLength1 := length1
	newLength2 := length2
	for i := 0; i < 10-len(length1); i++ {
		newLength1 = "0" + newLength1
	}

	for i := 0; i < 8-len(length2); i++ {
		newLength2 = "0" + newLength2
	}
	var s string
	for i := 0; i < 20-len(data.Qydm); i++ {
		s += " "
	}

	head := "A001" + // 报文版本
		"03" + // 目标系统
		"01" + // 报文编码
		"01" + // 通讯协议
		data.Qydm + //平台代码
		s +
		newLength1 + // 报文长度
		"000000" + // 交易码
		"PA001" + // 操作袁代码
		"01" + // 服务类型
		dateFormat + // 交易日期+交易时间
		data.ThirdLogNo + // 请求方系统流水号
		"000000" + // 返回码
		fmt.Sprintf("%100s", "") + // 返回描述
		"0" + // 后续包标志
		"000" + // 请求次数
		"0" + // 签名标识
		" " + // 签名数据包格式
		fmt.Sprintf("%12s", "") + // 签名算法
		"0000000000" + // 签名数据长度
		"0" + // 附件数目
		data.TranFunc +
		"01" +
		"                " +
		dateFormat +
		"000000" +
		"                                          " +
		"0" +
		newLength2 +
		"PA001" +
		data.ThirdLogNo +
		data.Qydm
	return head
}

// Transform 编码格式转换
func Transform(b []byte, t transform.Transformer) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(b), t)
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return b, err
}

// 获取请求方系统流水号
func getThirdLogNo() string {
	now := time.Now()
	rand.Seed(now.UnixNano())
	return fmt.Sprintf("%s%d", now.Format("20060102150405"), rand.Intn(10000))
}
