package oss_curd

import (
	"fmt"
	"io/ioutil"
	"label_system/config"
	"label_system/utils/genrandom"
	"log"
	"os"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type OssClient struct {
	Conn   *oss.Client
	Bucket *oss.Bucket
}

var Oc *OssClient = OssInit()

func OssInit() *OssClient {

	client, err := oss.New(config.Conf.OssEndpoint, config.Conf.OssAccessId, config.Conf.OssAccessKey)
	if err != nil {
		log.Println("Error:", err)
		return nil
	}

	bucket, err := client.Bucket(config.Conf.OssBucketName)
	if err != nil {
		log.Println("Error:", err)
		return nil
	}

	return &OssClient{Conn: client, Bucket: bucket}
}

func (oc *OssClient) LoadDataset(objectpath string) []QuestionRaw {

	objectpath = fmt.Sprintf("Input/%s", objectpath)
	var filelist = make([]string, 0)
	continueToken := ""
	prefix := oss.Prefix(objectpath)
	for {
		lsRes, err := oc.Bucket.ListObjectsV2(prefix, oss.ContinuationToken(continueToken))
		if err != nil {
			fmt.Printf("%v", err)
		}
		for _, object := range lsRes.Objects {
			if strings.Contains(object.Key, ".json") {
				filelist = append(filelist, object.Key)
			}
		}
		if lsRes.IsTruncated {
			continueToken = lsRes.NextContinuationToken
			prefix = oss.Prefix(lsRes.Prefix)
		} else {
			break
		}
	}
	randFile := filelist[genrandom.GenrandInt(len(filelist))]
	body, err := oc.Bucket.GetObject(randFile)
	if err != nil {
		log.Println("Error:", err)
		return nil
	}
	defer body.Close()

	data, _ := ioutil.ReadAll(body)
	results := strings.Split(string(data), "\n")
	question := results[genrandom.GenrandInt(len(results))]
	queJson := Stream2Json([]string{question})
	return queJson
}

func (oc *OssClient) SaveLabeled(filename, username string, label string) bool {
	env := os.Getenv("APPENV")
	objectpath := fmt.Sprintf("Output/%s/%s/%s.json", env, filename, username)
	if err := oc.Bucket.PutObject(objectpath, strings.NewReader(label)); err != nil {
		fmt.Printf("err: %v", err)
		log.Printf("err: %v", err)
		return false
	}
	return true
}
