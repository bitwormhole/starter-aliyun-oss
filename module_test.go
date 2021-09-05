package ossstarter

import (
	"embed"
	"testing"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/tests"
)

func TestModule(t *testing.T) {

	i := tests.TestingStarter(t)
	i.Use(Module())
	i.Use(testingModule())

	i.RunEx()
}

//go:embed src/test/resources
var theTestResFS embed.FS

func testingModule() application.Module {
	mb := &application.ModuleBuilder{}
	mb.Name(myName + "#testing").Version(myVersion).Revision(myRevision)
	mb.Resources(collection.LoadEmbedResources(&theTestResFS, "src/test/resources"))
	mb.Dependency(Module())
	return mb.Create()
}
