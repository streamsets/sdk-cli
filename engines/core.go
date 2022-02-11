package engines

type Engines interface {
	CreateConnection()
	CreateLibrary()
}

type DataCollector struct {
	Name       string
	Connection bool
	Library    bool
}
