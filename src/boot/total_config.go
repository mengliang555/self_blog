package boot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"self_blog/src/util/err_util"
)

type systemConfig struct {
	RateLimitConfig *RateLimit `json:"rate_limit"`
	MysqlConfig     *Mysql     `json:"mysql"`
}

func (c *systemConfig) String() string {
	marshal, err := json.Marshal(c)
	err_util.PanicError(err)
	return string(marshal)
}

type RateLimit struct {
	BucketSize   int `json:"bucket_size"`
	RateForToken int `json:"rate_for_token"`
}

type Mysql struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

var innConfig *systemConfig

func RefConfig() *systemConfig {
	return innConfig
}

func init() {
	innConfig = new(systemConfig)
	str, err := os.Getwd()
	err_util.PanicError(err)

	file, err := ioutil.ReadFile(fmt.Sprintf("%s\\resource\\config\\config.json", str))
	err_util.PanicError(err)

	err_util.PanicError(json.Unmarshal(file, innConfig))
	fmt.Printf("%s\n", innConfig)
}
