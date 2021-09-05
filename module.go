package ossstarter

import (
	"embed"

	"github.com/bitwormhole/starter"
	etcoss "github.com/bitwormhole/starter-aliyun-oss/etc/oss"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	myName     = "github.com/bitwormhole/starter-aliyun-oss"
	myVersion  = "v0.0.1"
	myRevision = 1
)

//go:embed src/main/resources
var theResFS embed.FS

// Module 定义并导出本模块(starter-aliyun-oss)
func Module() application.Module {

	mb := &application.ModuleBuilder{}

	mb.Name(myName).Version(myVersion).Revision(myRevision)
	mb.OnMount(etcoss.ExportConfig)
	mb.Resources(collection.LoadEmbedResources(&theResFS, "src/main/resources"))
	mb.Dependency(starter.Module())

	return mb.Create()
}
