package model

import (
	"database/sql"
	"io/ioutil"
	"net/http"
	"strings"
)

type DBCommand struct {
	order   string
	gateway *sql.DB
}

func (self *DBCommand) Execute() string {
	if self.order == "" {
		return ""
	}

	_, err := self.gateway.Exec(self.order)

	if err != nil {
		return err.Error()
	}
	return self.order + "executed!"
}

type HttpCommand struct {
	orderUrl  string
	orderBody string
	gateway   *http.Client
}

func (self *HttpCommand) Execute() string {
	if self.orderUrl == "" {
		return ""
	}
	// POSTリクエストの生成 TODO うまく行かなかったらここを参照 https://qiita.com/yyoshiki41/items/fc878494e19b9de93d56
	req, err := http.NewRequest("POST", self.orderUrl, strings.NewReader(self.orderBody))
	if err != nil {
		return err.Error()
	}

	result, err := self.gateway.Do(req)

	if err != nil {
		return err.Error()
	}

	defer result.Body.Close()
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return err.Error()
	}
	// bodyをstringに変換する方法を調査
	return string(body)
}

// type AeroCommand struct {
// 	order string
// 	gateway
// }
//
// func (self *AeroCommand) Execute() string {
// 	if self.order == "" {
// 		return ""
// 	}
//
// 	_, err := self.gateway.Exec(self.order)
//
// 	if err != nil {
// 		return err.Error()
// 	}
// 	return self.order + "executed!"
// }
