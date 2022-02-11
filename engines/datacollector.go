package engines

import (
	"fmt"

	"github.com/streamsets/sdk-cli/util"
)

func NewDataCollector(name string, connection bool, library bool) *DataCollector {
	d := new(DataCollector)
	d.Connection = connection
	d.Library = library
	d.Name = name

	return d
}

func (dc *DataCollector) CreateLibrary() {
	if dc.Library {
		util.CreateProjectDirectory(fmt.Sprintf("%s-%s", dc.Name, "lib"))
	}


}

func (dc *DataCollector) CreateConnection() {
	if dc.Connection {
		util.CreateProjectDirectory(fmt.Sprintf("%s-%s", dc.Name, "connection"))
	}

	
}
