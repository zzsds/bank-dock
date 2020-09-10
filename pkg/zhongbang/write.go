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

// Sign 加签
func (h *Zhongbang) Sign(params map[string]interface{}) string {
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
	if h.Secret == "" {
		log.Printf("商家密钥不能为空")
	} else {
		p = append(p, h.Secret)
	}
	// log.Println(strings.Join(p, "&"))
	return fmt.Sprintf("%x", md5.Sum([]byte(strings.Join(p, "&"))))
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
	params["signature"] = h.Sign(params)

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
	params["signature"] = h.Sign(params)

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
	params["signature"] = h.Sign(params)

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

	signStr := fmt.Sprintf("cardno=%s&traceno=%s&amount=%s&accountno=%s&mobile=%s&bankno=%s&key=%s", h.Cardno, msg.GetTraceno(), msg.GetAmount(), msg.GetAccountno(), msg.GetPayMode(), msg.GetBankno(), h.AgentSecret)
	sign := []byte(signStr)
	// sign, err = h.Transform(sign, h.Simplified.NewEncoder())
	// if err != nil {
	// 	return nil, err
	// }
	params["cardno"] = h.Cardno
	params["signature"] = fmt.Sprintf("%x", md5.Sum(sign))

	// log.Fatal(params, string(sign))
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
