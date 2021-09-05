package ossclient

import "github.com/bitwormhole/starter/markup"

////////////////////////////////////////////////////////////////////////////////

// OssDataSource OSS 数据源
type OssDataSource struct {
	markup.Component `id:"aliyun-oss-data-source"`

	Endpoint          string `inject:"${aliyun.oss.endpoint}"`
	AccessKeyID       string `inject:"${aliyun.oss.access-key-id}"`
	AccessKeySecret   string `inject:"${aliyun.oss.access-key-secret}"`
	DefaultBucketName string `inject:"${aliyun.oss.default-bucket-name}"`

	Connector Connector `inject:"#aliyun-oss-connector"`

	bucketNameList []string
}

func (inst *OssDataSource) _Impl() DataSource {
	return inst
}

// GetDefaultBucket  取默认的 bucket 名称
func (inst *OssDataSource) GetDefaultBucket() string {
	return inst.DefaultBucketName
}

// GetBucketList 取本地缓存的 bucket 列表
func (inst *OssDataSource) GetBucketList() []string {
	return inst.bucketNameList
}

// FetchBucketList 从服务端获取bucket列表
func (inst *OssDataSource) FetchBucketList() ([]string, error) {

	// TODO ...
	return inst.bucketNameList, nil
}

// Connect 连接到 OSS 服务
func (inst *OssDataSource) Connect(bucket string) (Connection, error) {

	if bucket == "" {
		bucket = inst.DefaultBucketName
	}

	cfg := &Configuration{}
	cfg.Endpoint = inst.Endpoint
	cfg.AccessKeyID = inst.AccessKeyID
	cfg.AccessKeySecret = inst.AccessKeySecret
	cfg.BucketName = bucket

	return inst.Connector.Connect(inst, cfg)
}

////////////////////////////////////////////////////////////////////////////////
