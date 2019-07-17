package httputils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	HTTPGet              = "GET"
	HTTPPost             = "POST"
	HTTPDelete           = "DELETE"
	HTTPPut              = "PUT"
	HeaderContentType    = "Content-Type"
	HeaderAuthentication = "Authentication"
	HeaderSignature      = "Signature"
	HeaderTimestamp      = "Timestamp"
	ContentTypeJSON      = "application/json"
	ContentTypeForm      = "application/x-www-form-urlencoded; param=value"
)

func ParseURL(s string, m map[string]string) string {
	url := s
	for k, v := range m {
		url = strings.Replace(url, k, v, 1)
	}
	return url
}

type Authentication struct {
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

type RestClient struct {
	auth      *Authentication
	basicAuth bool
	user      string
	password  string
	headers   map[string]string
}

func NewRestClientWithBasicAuth(user, password string) *RestClient {
	client := new(RestClient)
	client.basicAuth = true
	client.user = user
	client.password = password
	client.headers = make(map[string]string, 0)
	client.SetHeader(HeaderContentType, ContentTypeJSON)
	return client
}

func NewRestClientWithAuthentication(auth *Authentication) *RestClient {
	client := new(RestClient)
	if auth != nil {
		client.auth = auth
	}
	client.headers = make(map[string]string, 0)
	client.SetHeader(HeaderContentType, ContentTypeJSON)
	return client
}

func (c *RestClient) SetHeader(key, value string) {
	c.headers[key] = value
}

func (c *RestClient) applyHeader(req *http.Request) {
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}
}

func (c *RestClient) signature(url string) {
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	s := fmt.Sprintf("%s%s%s", c.auth.Key, ts, c.auth.Secret)
	hash := md5.New()
	hash.Write([]byte(s))
	signature := hex.EncodeToString(hash.Sum(nil))
	c.SetHeader(HeaderAuthentication, c.auth.Key)
	c.SetHeader(HeaderSignature, signature)
	c.SetHeader(HeaderTimestamp, ts)
}

func (c *RestClient) Get(uri string) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(HTTPGet, uri, nil)

	if c.basicAuth {
		req.SetBasicAuth(c.user, c.password)
	}

	if c.auth != nil {
		c.signature(uri)
	}

	c.applyHeader(req)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *RestClient) Post(uri string, data []byte) ([]byte, error) {
	buffer := bytes.NewBuffer(data)
	client := &http.Client{}
	req, _ := http.NewRequest(HTTPPost, uri, buffer)

	if c.basicAuth {
		req.SetBasicAuth(c.user, c.password)
	}

	if c.auth != nil {
		c.signature(uri)
	}

	c.applyHeader(req)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *RestClient) Put(uri string, data []byte) ([]byte, error) {
	buffer := bytes.NewBuffer(data)
	client := &http.Client{}
	req, _ := http.NewRequest(HTTPPut, uri, buffer)

	if c.basicAuth {
		req.SetBasicAuth(c.user, c.password)
	}

	if c.auth != nil {
		c.signature(uri)
	}

	c.applyHeader(req)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *RestClient) Delete(uri string) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(HTTPDelete, uri, nil)

	if c.basicAuth {
		req.SetBasicAuth(c.user, c.password)
	}

	if c.auth != nil {
		c.signature(uri)
	}

	c.applyHeader(req)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
