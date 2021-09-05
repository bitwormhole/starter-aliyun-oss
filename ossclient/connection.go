package ossclient

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/bitwormhole/starter/vlog"
)

////////////////////////////////////////////////////////////////////////////////

// defaultConnection 表示一个连接器
type defaultConnection struct {
	bucketName string
	client     *oss.Client
	bucket     *oss.Bucket
	ds         DataSource
}

func (inst *defaultConnection) _Impl() Connection {
	return inst
}

func (inst *defaultConnection) init() (Connection, error) {
	return inst, nil
}

func (inst *defaultConnection) BucketName() string {
	return inst.bucketName
}

func (inst *defaultConnection) DataSource() DataSource {
	return inst.ds
}

func (inst *defaultConnection) Client() *oss.Client {
	return inst.client
}

func (inst *defaultConnection) Bucket() *oss.Bucket {
	return inst.bucket
}

func (inst *defaultConnection) Close() error {
	inst.bucket = nil
	inst.client = nil
	inst.ds = nil
	return nil
}

func (inst *defaultConnection) Test() error {

	bucket := inst.bucket
	name := "index.html"
	isExist, err := bucket.IsObjectExist(name)

	if err != nil {
		return err
	}

	if isExist {
		vlog.Debug("the file is exists, name=" + name)
	} else {
		vlog.Debug("the file is NOT exists, name=" + name)
	}

	return nil
}
