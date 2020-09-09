package zhongbang

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"

	proto_zhongbang "github.com/zzsds/zhongbang/pkg/zhongbang/proto"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// Service ...
type Service interface {
	Transform([]byte, transform.Transformer) ([]byte, error)
	OpenAcct(*proto_zhongbang.OpenAcctMessage) (*proto_zhongbang.OpenAcctResponse, error)
}

// Zhongbang ...
type Zhongbang struct {
	Host       string
	Merchno    string
	Simplified encoding.Encoding
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
func (h *Zhongbang) Sign(msg interface{}) (string, error) {
	params := make(map[string]interface{})
	b, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(b, &params); err != nil {
		return "", err
	}
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

	return "123123", nil
}

// OpenAcct 商户新增白名单的用户信息
func (h *Zhongbang) OpenAcct(msg *proto_zhongbang.OpenAcctMessage) (rsp *proto_zhongbang.OpenAcctResponse, err error) {
	route := h.Host + "/openAcct.do"
	msg.Merchno = h.Merchno
	msg.Signature, err = h.Sign(msg)
	if err != nil {
		return
	}

	b, err := json.Marshal(msg)
	if err != nil {
		return
	}

	b, err = h.Transform(b, h.Simplified.NewEncoder())
	if err != nil {
		return
	}
	resq, err := http.NewRequest("POST", route, bytes.NewBuffer(b))
	if err != nil {
		return
	}
	resq.Header.Set("Content-Type", "text/html;charset=GBK")
	client := &http.Client{}
	resp, err := client.Do(resq)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status fail")
	}
	fmt.Println("status:", resp.Status)
	fmt.Println("response:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	b, err = h.Transform(body, h.Simplified.NewDecoder())
	if err := json.Unmarshal(b, &rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}
