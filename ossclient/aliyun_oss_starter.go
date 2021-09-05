package ossclient

import "github.com/bitwormhole/starter/markup"

// AliyunOssStarter 控制OSS数据源
type AliyunOssStarter struct {
	markup.Component `id:"aliyun-oss-agent" class:"aliyun-oss-starter" initMethod:"Start" destroyMethod:"Stop"`

	DS     DataSource `inject:"#aliyun-oss-data-source"`
	Enable bool       `inject:"${aliyun.oss.enable}"`

	connection Connection
}

func (inst *AliyunOssStarter) _Impl() ConnectionHolder {
	return inst
}

func (inst *AliyunOssStarter) GetConnection() Connection {
	return inst.connection
}

func (inst *AliyunOssStarter) Start() error {
	if inst.Enable {
		return inst.doStart()
	}
	return nil
}

func (inst *AliyunOssStarter) doStart() error {
	name := inst.DS.GetDefaultBucket()
	conn, err := inst.DS.Connect(name)
	if err != nil {
		return err
	}
	err = conn.Test()
	if err != nil {
		return err
	}
	inst.connection = conn
	return nil
}

func (inst *AliyunOssStarter) Stop() error {
	conn := inst.connection
	inst.connection = nil
	if conn == nil {
		return nil
	}
	return conn.Close()
}
