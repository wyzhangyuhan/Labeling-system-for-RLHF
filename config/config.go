package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type conf struct {
	AckdgeDomain string `yaml:"ackedge_domain"`
	ServerPort   string `yaml:"server_port"`
	DatabaseDsn  string `yaml:"database_dsn"`
	S3Domain     string `yaml:"s3_domain"`
	BucketName   string `yaml:"bucket_name"`
	S3AccessId   string `yaml:"s3_access_id"`
	S3AccessKey  string `yaml:"s3_access_key"`
	SqlUsername  string `yaml:"DBusername"`
	SqlPassword  string `yaml:"DBpassword"`
	SqlHost      string `yaml:"DBhost"`
	SqlName      string `yaml:"DBname"`

	OssEndpoint   string `yaml:"oss_endpoint"`
	OssBucketName string `yaml:"oss_bucket_name"`
	OssAccessId   string `yaml:"oss_access_id"`
	OssAccessKey  string `yaml:"oss_access_key"`
}

var Conf conf

var APPENV string

func init() {
	APPENV = os.Getenv("APPENV")
	fmt.Printf("APPENV:%s\n", APPENV)
	file, _ := ioutil.ReadFile(fmt.Sprintf("config/%s.yaml", APPENV))
	yaml.Unmarshal(file, &Conf)
	fmt.Printf("Conf:%+v\n", Conf)
}
