package engines

const VERSION = "4.4.0-SNAPSHOT"

type Engines interface {
	CreateConnection()
	CreateLibrary()
}

type DataCollector struct {
	Name       string
	Connection bool
	Library    bool
	Root       string
}
