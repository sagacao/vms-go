package settings

import (
	"flag"
	"io/ioutil"
	"path"
	"time"

	jsoniter "github.com/json-iterator/go"
)

type JwtConfig struct {
	SECRET string
	EXP    time.Duration // 过期时间
	ALG    string        // 算法
}

func GetJwtConfig() *JwtConfig {
	return &JwtConfig{
		SECRET: "qgame_session",
		EXP:    time.Hour,
		ALG:    "HS256",
	}
}

type CookieConfig struct {
	NAME       string
	APP_SECRET string
}

func GetCookieConfig() *CookieConfig {
	return &CookieConfig{
		NAME:       "qgame_session",
		APP_SECRET: "YbskZqLNT6TEVLUA9HWdnHmZErypNJpL",
	}
}

var AppPath string

var SvrConfig struct {
	Env struct {
		SERVER_ADDR string
		DEBUG       bool

		ACCESS_LOG_PATH string
		ACCESS_LOG_NAME string

		URL           string
		TEMPLATE_PATH string
	}

	Mysql struct {
		DBNAME string
		USER   string
		HOST   string
		PASSWD string
	}
}

var userSecret map[string]string

func initEnv(apppath string) {
	path.Join()
	data, err := ioutil.ReadFile(path.Join(apppath, "env.json"))
	if err != nil {
		panic(err)
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(data, &SvrConfig)
	if err != nil {
		panic(err)
	}
}

func initUserSecret(apppath string) {
	path.Join()
	data, err := ioutil.ReadFile(path.Join(apppath, "user.secret.json"))
	if err != nil {
		panic(err)
	}
	userSecret = make(map[string]string)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(data, &userSecret)
	if err != nil {
		panic(err)
	}
}

func GetUser(user string) (string, bool) {
	v, ok := userSecret[user]
	if !ok {
		return "", false
	}

	return v, true
}

func init() {
	flag.StringVar(&AppPath, "app-path", "config", "--app-path")
	flag.Parse()
	initEnv(AppPath)
	initUserSecret(AppPath)
}
