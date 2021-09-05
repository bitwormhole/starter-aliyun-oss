package ossclient

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/bitwormhole/starter/markup"
)

// DefaultConnector 表示一个连接器
type DefaultConnector struct {
	markup.Component `id:"aliyun-oss-connector"`
}

func (inst *DefaultConnector) _Impl() Connector {
	return inst
}

// Connect 连接到 OSS 服务
func (inst *DefaultConnector) Connect(ds DataSource, cfg *Configuration) (Connection, error) {

	client, err := oss.New(cfg.Endpoint, cfg.AccessKeyID, cfg.AccessKeySecret)
	if err != nil {
		return nil, err
	}

	bucket, err := client.Bucket(cfg.BucketName)
	if err != nil {
		return nil, err
	}

	conn := &defaultConnection{}
	conn.ds = ds
	conn.bucket = bucket
	conn.bucketName = cfg.BucketName
	conn.client = client
	return conn.init()
}
