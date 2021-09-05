package oss

import (
	"github.com/bitwormhole/starter/application"
)

const (
	pAliyunOssEndpoint          = "aliyun.oss.endpoint"
	pAliyunOssAccessKeyId       = "aliyun.oss.access-key-id"
	pAliyunOssAccessKeySecret   = "aliyun.oss.access-key-secret"
	pAliyunOssDefaultBucketName = "aliyun.oss.default-bucket-name"
)

func ExportConfig(cb application.ConfigBuilder) error {

	return autoGenConfig(cb)
}
