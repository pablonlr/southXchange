package request

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//Client Basic SX Client struct
type Client struct {
	api       string
	secretAPI string
	http.Client
}

const baseURL = "https://www.southxchange.com/api/"

func NewSouthXchangeClient(API, SECRET_API string, timeout time.Duration) *Client {
	return &Client{
		Client:    *&http.Client{Timeout: timeout},
		api:       API,
		secretAPI: SECRET_API,
	}
}

//GetReq Make a GET request following the endopoints received
func (client *Client) GetReq(endpoints ...string) ([]byte, error) {
	url := baseURL + strings.Join(endpoints, "/")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return client.doReq(req)
}
func (client *Client) doReq(req *http.Request) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 {
		return nil, errors.New(string(body))
	}
	return body, nil
}

func (client *Client) PostReq(endpoint string, body map[string]interface{}) ([]byte, error) {
	body["key"] = client.api
	body["nonce"] = timeStamp()
	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", baseURL+endpoint, bytes.NewBuffer(reqBody))
	req.Header.Set("Hash", signBody(reqBody, client.secretAPI))
	req.Header.Set("content-type", "application/json")
	return client.doReq(req)
}

func signBody(body []byte, secret string) string {
	hm := hmac.New(sha512.New, []byte(secret))
	hm.Write(body)
	sha := hex.EncodeToString(hm.Sum(nil))
	return sha
}
func timeStamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
