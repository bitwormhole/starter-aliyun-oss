package ossclient

import (
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// Configuration 表示一个连接的配置
type Configuration struct {
	BucketName      string
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
}

// Connection 表示一个连接
type Connection interface {
	io.Closer
	Client() *oss.Client
	Bucket() *oss.Bucket
	BucketName() string
	DataSource() DataSource
	Test() error
}

// ConnectionHolder 持有一个 OSS 连接
type ConnectionHolder interface {
	GetConnection() Connection
}

// Connector 表示一个连接器
type Connector interface {
	Connect(ds DataSource, c *Configuration) (Connection, error)
}

// DataSource 表示一个 oss 数据源
type DataSource interface {
	GetDefaultBucket() string
	GetBucketList() []string
	FetchBucketList() ([]string, error)
	Connect(bucket string) (Connection, error)
}
