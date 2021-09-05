// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package oss

import (
	ossclient0x35aae4 "github.com/bitwormhole/starter-aliyun-oss/ossclient"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
    
)

func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()

	// component: aliyun-oss-agent
	cominfobuilder.Next()
	cominfobuilder.ID("aliyun-oss-agent").Class("aliyun-oss-starter").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComAliyunOssStarter{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: aliyun-oss-connector
	cominfobuilder.Next()
	cominfobuilder.ID("aliyun-oss-connector").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDefaultConnector{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: aliyun-oss-data-source
	cominfobuilder.Next()
	cominfobuilder.ID("aliyun-oss-data-source").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComOssDataSource{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComAliyunOssStarter : the factory of component: aliyun-oss-agent
type comFactory4pComAliyunOssStarter struct {

    mPrototype * ossclient0x35aae4.AliyunOssStarter

	
	mDSSelector config.InjectionSelector
	mEnableSelector config.InjectionSelector

}

func (inst * comFactory4pComAliyunOssStarter) init() application.ComponentFactory {

	
	inst.mDSSelector = config.NewInjectionSelector("#aliyun-oss-data-source",nil)
	inst.mEnableSelector = config.NewInjectionSelector("${aliyun.oss.enable}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComAliyunOssStarter) newObject() * ossclient0x35aae4.AliyunOssStarter {
	return & ossclient0x35aae4.AliyunOssStarter {}
}

func (inst * comFactory4pComAliyunOssStarter) castObject(instance application.ComponentInstance) * ossclient0x35aae4.AliyunOssStarter {
	return instance.Get().(*ossclient0x35aae4.AliyunOssStarter)
}

func (inst * comFactory4pComAliyunOssStarter) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComAliyunOssStarter) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComAliyunOssStarter) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComAliyunOssStarter) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Start()
}

func (inst * comFactory4pComAliyunOssStarter) Destroy(instance application.ComponentInstance) error {
	return inst.castObject(instance).Stop()
}

func (inst * comFactory4pComAliyunOssStarter) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.DS = inst.getterForFieldDSSelector(context)
	obj.Enable = inst.getterForFieldEnableSelector(context)
	return context.LastError()
}

//getterForFieldDSSelector
func (inst * comFactory4pComAliyunOssStarter) getterForFieldDSSelector (context application.InstanceContext) ossclient0x35aae4.DataSource {

	o1 := inst.mDSSelector.GetOne(context)
	o2, ok := o1.(ossclient0x35aae4.DataSource)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "aliyun-oss-agent")
		eb.Set("field", "DS")
		eb.Set("type1", "?")
		eb.Set("type2", "ossclient0x35aae4.DataSource")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldEnableSelector
func (inst * comFactory4pComAliyunOssStarter) getterForFieldEnableSelector (context application.InstanceContext) bool {
    return inst.mEnableSelector.GetBool(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDefaultConnector : the factory of component: aliyun-oss-connector
type comFactory4pComDefaultConnector struct {

    mPrototype * ossclient0x35aae4.DefaultConnector

	

}

func (inst * comFactory4pComDefaultConnector) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDefaultConnector) newObject() * ossclient0x35aae4.DefaultConnector {
	return & ossclient0x35aae4.DefaultConnector {}
}

func (inst * comFactory4pComDefaultConnector) castObject(instance application.ComponentInstance) * ossclient0x35aae4.DefaultConnector {
	return instance.Get().(*ossclient0x35aae4.DefaultConnector)
}

func (inst * comFactory4pComDefaultConnector) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDefaultConnector) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDefaultConnector) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDefaultConnector) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDefaultConnector) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDefaultConnector) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComOssDataSource : the factory of component: aliyun-oss-data-source
type comFactory4pComOssDataSource struct {

    mPrototype * ossclient0x35aae4.OssDataSource

	
	mEndpointSelector config.InjectionSelector
	mAccessKeyIDSelector config.InjectionSelector
	mAccessKeySecretSelector config.InjectionSelector
	mDefaultBucketNameSelector config.InjectionSelector
	mConnectorSelector config.InjectionSelector

}

func (inst * comFactory4pComOssDataSource) init() application.ComponentFactory {

	
	inst.mEndpointSelector = config.NewInjectionSelector("${aliyun.oss.endpoint}",nil)
	inst.mAccessKeyIDSelector = config.NewInjectionSelector("${aliyun.oss.access-key-id}",nil)
	inst.mAccessKeySecretSelector = config.NewInjectionSelector("${aliyun.oss.access-key-secret}",nil)
	inst.mDefaultBucketNameSelector = config.NewInjectionSelector("${aliyun.oss.default-bucket-name}",nil)
	inst.mConnectorSelector = config.NewInjectionSelector("#aliyun-oss-connector",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComOssDataSource) newObject() * ossclient0x35aae4.OssDataSource {
	return & ossclient0x35aae4.OssDataSource {}
}

func (inst * comFactory4pComOssDataSource) castObject(instance application.ComponentInstance) * ossclient0x35aae4.OssDataSource {
	return instance.Get().(*ossclient0x35aae4.OssDataSource)
}

func (inst * comFactory4pComOssDataSource) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComOssDataSource) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComOssDataSource) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComOssDataSource) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComOssDataSource) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComOssDataSource) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Endpoint = inst.getterForFieldEndpointSelector(context)
	obj.AccessKeyID = inst.getterForFieldAccessKeyIDSelector(context)
	obj.AccessKeySecret = inst.getterForFieldAccessKeySecretSelector(context)
	obj.DefaultBucketName = inst.getterForFieldDefaultBucketNameSelector(context)
	obj.Connector = inst.getterForFieldConnectorSelector(context)
	return context.LastError()
}

//getterForFieldEndpointSelector
func (inst * comFactory4pComOssDataSource) getterForFieldEndpointSelector (context application.InstanceContext) string {
    return inst.mEndpointSelector.GetString(context)
}

//getterForFieldAccessKeyIDSelector
func (inst * comFactory4pComOssDataSource) getterForFieldAccessKeyIDSelector (context application.InstanceContext) string {
    return inst.mAccessKeyIDSelector.GetString(context)
}

//getterForFieldAccessKeySecretSelector
func (inst * comFactory4pComOssDataSource) getterForFieldAccessKeySecretSelector (context application.InstanceContext) string {
    return inst.mAccessKeySecretSelector.GetString(context)
}

//getterForFieldDefaultBucketNameSelector
func (inst * comFactory4pComOssDataSource) getterForFieldDefaultBucketNameSelector (context application.InstanceContext) string {
    return inst.mDefaultBucketNameSelector.GetString(context)
}

//getterForFieldConnectorSelector
func (inst * comFactory4pComOssDataSource) getterForFieldConnectorSelector (context application.InstanceContext) ossclient0x35aae4.Connector {

	o1 := inst.mConnectorSelector.GetOne(context)
	o2, ok := o1.(ossclient0x35aae4.Connector)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "aliyun-oss-data-source")
		eb.Set("field", "Connector")
		eb.Set("type1", "?")
		eb.Set("type2", "ossclient0x35aae4.Connector")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}




