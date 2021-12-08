package gateway

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/agclqq/goencryption"
	"io/ioutil"
	"net/http"
)

var key = []byte("0123456789ABCDEF")  //Key的长度16, 24, 32 分别对应 AES-128, AES-192, AES-256
var iv = []byte("0123456789ABCDEF")


/**
通过HTTP请求获取信息
*/
func HttpGetBytes(url string,) ([]byte, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cannot send request : %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 status code: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body err, %v\n", err)
	}
	return body, nil

}

func AES7(plainText string) string {

	cryptText, err := goencryption.AesCBCPkcs7Encrypt([]byte(plainText), key, iv)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return hex.EncodeToString(cryptText)

}


/**
通过HTTP请求获取信息
*/
func HttpPostBytes(url string, bytesData []byte) ([]byte, error) {

	reader := bytes.NewReader(bytesData)

	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cannot send request : %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 status code: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body err, %v\n", err)
	}

	return body, nil

}

