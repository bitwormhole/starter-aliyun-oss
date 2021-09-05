// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package oss

import (
	ossclient0x35aae4 "github.com/bitwormhole/starter-aliyun-oss/ossclient"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComAliyunOssStarter struct {
	instance *ossclient0x35aae4.AliyunOssStarter
	 markup0x23084a.Component `id:"aliyun-oss-agent" class:"aliyun-oss-starter" initMethod:"Start" destroyMethod:"Stop"`
	DS ossclient0x35aae4.DataSource `inject:"#aliyun-oss-data-source"`
	Enable bool `inject:"${aliyun.oss.enable}"`
	connection ossclient0x35aae4.Connection ``
}


type pComDefaultConnector struct {
	instance *ossclient0x35aae4.DefaultConnector
	 markup0x23084a.Component `id:"aliyun-oss-connector"`
}


type pComOssDataSource struct {
	instance *ossclient0x35aae4.OssDataSource
	 markup0x23084a.Component `id:"aliyun-oss-data-source"`
	Endpoint string `inject:"${aliyun.oss.endpoint}"`
	AccessKeyID string `inject:"${aliyun.oss.access-key-id}"`
	AccessKeySecret string `inject:"${aliyun.oss.access-key-secret}"`
	DefaultBucketName string `inject:"${aliyun.oss.default-bucket-name}"`
	Connector ossclient0x35aae4.Connector `inject:"#aliyun-oss-connector"`
	bucketNameList []string ``
}

