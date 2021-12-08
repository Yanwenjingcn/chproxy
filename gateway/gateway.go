package gateway

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/common/log"
)

type userPassword struct {
	code int    `json:"code"`
	msg  string `json:"msg"`
	Data *rsaPassword
}

type rsaPassword struct {
	RsaPassword string `json:"rsaPassword"`
}

/**
 校验：用户名+密码
 */
func CheckUserPw(userName string, passwd string) (bool, error) {
	if userName == "" {
		return false, fmt.Errorf("username is blank")
	}

	body, err := HttpGetBytes(fmt.Sprintf("http://localhost:8188/api/gateway/getUserPassword?userName=%s", userName))
	if err != nil {
		return false, err
	}
	var a userPassword
	if err = json.Unmarshal(body, &a); err != nil {
		return false, fmt.Errorf("Unmarshal err, %v\n", err)
	}

	rsaPw := a.Data.RsaPassword
	log.Info(rsaPw)

	if rsaPw == AES7(passwd) {
		return true, nil
	}

	return false, fmt.Errorf("invalid username or password for userName %q", userName)

}

type sqlParseResp struct {
	code int    `json:"code"`
	msg  string `json:"msg"`
	Data []SQLOptResp
}


type SQLOptResp struct {
	opt        string `json:"opt"`
	cluster    string `json:"cluster"`
	db         string `json:"db"`
	objectType string `json:"type"`
	object     string `json:"object"`
}


func sqlParse(sql string) (bool,error)  {
	if sql == "" {
		return false, fmt.Errorf("sql is empty")
	}

	postBody := make(map[string]string)
	postBody["type"] = "clickhouse"
	postBody["sql"] = sql
	bytesData, err := json.Marshal(postBody)

	body, err := HttpPostBytes("http://localhost:8188/api/gateway/sqlParse", bytesData)
	if err != nil {
		return false, err
	}

	var a sqlParseResp
	if err = json.Unmarshal(body, &a); err != nil {
		return false, fmt.Errorf("Unmarshal err, %v\n", err)
	}

	return true,nil
}

func CheckGatewayPermission(reqBody string,username string)  (bool,error){
	return true,nil


}
