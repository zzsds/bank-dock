package zhongbang

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"

	proto_zhongbang "github.com/zzsds/bank-dock/pkg/zhongbang/proto"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// Service ...
type Service interface {
	Transform([]byte, transform.Transformer) ([]byte, error)
	OpenAcct(*proto_zhongbang.OpenAcctRequest) (*proto_zhongbang.OpenAcctResponse, error)
	OpenUpdate(*proto_zhongbang.OpenAcctRequest) (*proto_zhongbang.OpenAcctResponse, error)
	OpenQuery(*proto_zhongbang.OpenQueryRequest) (*proto_zhongbang.OpenAcctResponse, error)

	VirtPay(*proto_zhongbang.VirtPayRequest) (*proto_zhongbang.VirtPayResponse, error)
	VirtOrder(*proto_zhongbang.VirtOrderRequest) (*proto_zhongbang.VirtPayResponse, error)
	Balance(*proto_zhongbang.BalanceRequest) (*proto_zhongbang.BalanceResponse, error)
	VirtNotify(*proto_zhongbang.VirtNotifyRequest) error
}

// Params ...
type Params map[string]interface{}

// NewParmas ...
func NewParmas(data interface{}) Params {
	var params = make(map[string]interface{})

	b, err := json.Marshal(data)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(b, &params)
	if err != nil {
		return nil
	}
	return params
}

// Set ...
func (p Params) Set(key string, val interface{}) Params {
	p[key] = val
	return p
}

// Get ...
func (p Params) Get(key string) interface{} {
	return p[key]
}

// Zhongbang ...
type Zhongbang struct {
	Host        string
	Merchno     string
	Cardno      string
	Secret      string
	AgentSecret string
	Simplified  encoding.Encoding
}

// Options ...
type Options func(*Zhongbang)

// New ...
func New(opts ...Options) Service {
	var zb Zhongbang
	for _, o := range opts {
		o(&zb)
	}
	if zb.Simplified == nil {
		zb.Simplified = simplifiedchinese.GBK
	}
	return &zb
}

// Transform 编码格式转换
func (h *Zhongbang) Transform(b []byte, t transform.Transformer) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(b), t)
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return b, err
}

// OpenAcct 商户新增白名单的用户信息
func (h *Zhongbang) OpenAcct(msg *proto_zhongbang.OpenAcctRequest) (rsp *proto_zhongbang.OpenAcctResponse, err error) {
	u, _ := url.ParseRequestURI(h.Host)
	u.Path = "openAcct.do"

	params := make(map[string]interface{})
	b, err := json.Marshal(msg)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &params)
	if err != nil {
		return
	}

	params["merchno"] = h.Merchno
	params["signature"] = sign(params, h.Secret)

	data := url.Values{}
	for k, v := range params {
		data.Set(k, v.(string))
	}

	b, err = h.Transform([]byte(data.Encode()), h.Simplified.NewEncoder())
	if err != nil {
		return
	}

	body, err := request(http.MethodPost, u, b)
	if err != nil {
		return nil, err
	}
	b, err = h.Transform(body, h.Simplified.NewDecoder())
	if err := json.Unmarshal(b, &rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}

// OpenUpdate 商户修改白名单的用户信息
func (h *Zhongbang) OpenUpdate(msg *proto_zhongbang.OpenAcctRequest) (rsp *proto_zhongbang.OpenAcctResponse, err error) {
	u, _ := url.ParseRequestURI(h.Host)
	u.Path = "openUpdate.do"

	params := make(map[string]interface{})
	b, err := json.Marshal(msg)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &params)
	if err != nil {
		return
	}

	params["merchno"] = h.Merchno
	params["signature"] = sign(params, h.Secret)

	data := url.Values{}
	for k, v := range params {
		data.Set(k, v.(string))
	}

	b, err = h.Transform([]byte(data.Encode()), h.Simplified.NewEncoder())
	if err != nil {
		return
	}
	body, err := request(http.MethodPost, u, b)
	if err != nil {
		return nil, err
	}
	b, err = h.Transform(body, h.Simplified.NewDecoder())
	if err := json.Unmarshal(b, &rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}

// OpenQuery 查询商户白名单的用户信息
func (h *Zhongbang) OpenQuery(msg *proto_zhongbang.OpenQueryRequest) (rsp *proto_zhongbang.OpenAcctResponse, err error) {
	u, _ := url.ParseRequestURI(h.Host)
	u.Path = "openQuery.do"

	params := make(map[string]interface{})
	b, err := json.Marshal(msg)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &params)
	if err != nil {
		return
	}

	params["merchno"] = h.Merchno
	params["signature"] = sign(params, h.Secret)

	data := url.Values{}
	for k, v := range params {
		data.Set(k, v.(string))
	}

	b, err = h.Transform([]byte(data.Encode()), h.Simplified.NewEncoder())
	if err != nil {
		return
	}
	body, err := request(http.MethodPost, u, b)
	if err != nil {
		return nil, err
	}
	b, err = h.Transform(body, h.Simplified.NewDecoder())
	if err := json.Unmarshal(b, &rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}

// VirtPay 实时代付接口
func (h *Zhongbang) VirtPay(msg *proto_zhongbang.VirtPayRequest) (rsp *proto_zhongbang.VirtPayResponse, err error) {
	u, _ := url.ParseRequestURI(h.Host)
	u.Path = "virtPay.do"

	params := make(map[string]interface{})
	b, err := json.Marshal(msg)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &params)
	if err != nil {
		return
	}
	params["cardno"] = h.Cardno

	signStr := fmt.Sprintf("cardno=%s&traceno=%s&amount=%s&accountno=%s&mobile=%s&bankno=%s&key=%s", h.Cardno, msg.GetTraceno(), msg.GetAmount(), msg.GetAccountno(), msg.GetMobile(), msg.GetBankno(), h.AgentSecret)

	sign := fmt.Sprintf("%x", md5.Sum([]byte(signStr)))
	var buff bytes.Buffer
	for k, v := range params {
		buff.WriteString(fmt.Sprintf("%s=%s&", k, v.(string)))
	}
	buff.WriteString(fmt.Sprintf("signature=%s", sign))

	b, err = h.Transform(buff.Bytes(), h.Simplified.NewEncoder())
	if err != nil {
		return
	}

	body, err := request(http.MethodPost, u, b)
	if err != nil {
		return nil, err
	}
	b, err = h.Transform(body, h.Simplified.NewDecoder())
	if err := json.Unmarshal(b, &rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}

// VirtOrder 实时代付接口
func (h *Zhongbang) VirtOrder(msg *proto_zhongbang.VirtOrderRequest) (rsp *proto_zhongbang.VirtPayResponse, err error) {
	u, _ := url.ParseRequestURI(h.Host)
	u.Path = "virtOrder.do"

	params := make(map[string]interface{})
	b, err := json.Marshal(msg)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &params)
	if err != nil {
		return
	}
	params["cardno"] = h.Cardno
	signStr := fmt.Sprintf("cardno=%s&traceno=%s&key=%s", h.Cardno, msg.GetTraceno(), h.AgentSecret)

	sign := fmt.Sprintf("%x", md5.Sum([]byte(signStr)))
	var buff bytes.Buffer
	for k, v := range params {
		buff.WriteString(fmt.Sprintf("%s=%s&", k, v.(string)))
	}
	buff.WriteString(fmt.Sprintf("signature=%s", sign))

	b, err = h.Transform(buff.Bytes(), h.Simplified.NewEncoder())
	if err != nil {
		return
	}

	body, err := request(http.MethodPost, u, b)
	if err != nil {
		return nil, err
	}
	b, err = h.Transform(body, h.Simplified.NewDecoder())
	if err := json.Unmarshal(b, &rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}

// Balance 余额查询接口
func (h *Zhongbang) Balance(msg *proto_zhongbang.BalanceRequest) (rsp *proto_zhongbang.BalanceResponse, err error) {
	u, _ := url.ParseRequestURI(h.Host)
	u.Path = "balance.do"

	params := make(map[string]interface{})
	b, err := json.Marshal(msg)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &params)
	if err != nil {
		return
	}

	signStr := fmt.Sprintf("cardno=%s&key=%s", msg.GetCardno(), h.AgentSecret)

	sign := fmt.Sprintf("%x", md5.Sum([]byte(signStr)))
	var buff bytes.Buffer
	for k, v := range params {
		buff.WriteString(fmt.Sprintf("%s=%s&", k, v.(string)))
	}
	buff.WriteString(fmt.Sprintf("signature=%s", sign))

	b, err = h.Transform(buff.Bytes(), h.Simplified.NewEncoder())
	if err != nil {
		return
	}

	body, err := request(http.MethodPost, u, b)
	if err != nil {
		return nil, err
	}
	b, err = h.Transform(body, h.Simplified.NewDecoder())
	if err := json.Unmarshal(b, &rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}

// VirtNotify ...
func (h *Zhongbang) VirtNotify(msg *proto_zhongbang.VirtNotifyRequest) error {
	var params = NewParmas(msg)
	delete(params, "signature")
	signature := strings.ToUpper(sign(params, h.AgentSecret))
	if msg.GetSignature() != signature {
		return fmt.Errorf("validate sign fail")
	}
	return nil
}

func request(method string, url *url.URL, params []byte) ([]byte, error) {
	resq, err := http.NewRequest(method, url.String(), bytes.NewBuffer(params))
	if err != nil {
		return nil, err
	}

	resq.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=GBK")
	resq.Header.Set("Content-Length", strconv.Itoa(len(string(params))))
	client := &http.Client{}
	resp, err := client.Do(resq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status fail")
	}
	log.Printf("status: %s", resp.Status)
	log.Printf("response: %s", resp.Header)
	return ioutil.ReadAll(resp.Body)
}

// Sign 加签
func sign(params map[string]interface{}, secret string) string {
	var p []string
	for k := range params {
		p = append(p, k)
	}
	if !sort.StringsAreSorted(p) {
		sort.Strings(p)
	}
	for k, v := range p {
		if val, ok := params[v]; ok {
			p[k] = v + "=" + val.(string)
		}
	}
	p = append(p, secret)
	// log.Println(strings.Join(p, "&"))
	return fmt.Sprintf("%x", md5.Sum([]byte(strings.Join(p, "&"))))
}
