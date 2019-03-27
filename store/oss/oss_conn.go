package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go-filestore-server/config"
)

var ossCli *oss.Client

// 创建oss client对象
func Client() *oss.Client {
	if ossCli != nil {
		return ossCli
	}

	ossCli, err := oss.New(config.OSSEndpoint, config.OSSAccesskeyID, config.OSSAccesskeySecret)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return ossCli
}

// 获取bucket存储空间
func Bucket() *oss.Bucket {
	cli := Client()
	if cli != nil {
		bucket, err := cli.Bucket(config.OSSBucket)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		return bucket
	}
	return nil
}

// 临时授权下载url
func DownloadURL(objName string) string {
	signedURL, err := Bucket().SignURL(objName, oss.HTTPGet, 3600)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return signedURL
}

// 针对指定bucket设置生命周期规则
func BuildLifeCycleRule(bucketName string) {
	ruleTest1 := oss.BuildLifecycleRuleByDays("rul1", "/test", true, 30)
	rules := []oss.LifecycleRule{ruleTest1}

	Client().SetBucketLifecycle(bucketName, rules)
}
