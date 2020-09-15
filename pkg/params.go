package pkg

import (
	"encoding/json"
	"net/url"
)

// Params ...
type Params map[string]string

// NewParmas ...
func NewParmas(data interface{}) Params {
	var params = make(map[string]string)

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
func (p Params) Set(key string, val string) Params {
	p[key] = val
	return p
}

// Get ...
func (p Params) Get(key string) (string, bool) {
	val, ok := p[key]
	return val, ok
}

// Delete ...
func (p Params) Delete(key string) {
	delete(p, key)
}

// URL ...
func (p Params) URL() string {
	var u = url.Values{}
	for k, v := range p {
		u.Set(k, v)
	}
	return u.Encode()
}
