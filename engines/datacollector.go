package engines

import (
	"fmt"
	"os"
	"strings"

	"github.com/markbates/pkger"
	"github.com/streamsets/sdk-cli/util"
)

func NewDataCollector(name string, connection bool, library bool) *DataCollector {
	d := new(DataCollector)
	d.Connection = connection
	d.Library = library
	d.Name = name
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	d.Root = dir

	return d
}

func (dc *DataCollector) Build() {

	if dc.Connection {
		util.WriteStatus("--createConnection flag found. Building connection project")
		dir := dc.createConnection()
		util.WriteStatus(fmt.Sprintf("%s created", dir))
		dc.createConnectionSubDirectories(dir)
		dc.createConnectionPOM(dir)
		util.WriteStatus(fmt.Sprintf("%s pom.xml file created", dir))
		dc.createConnectionResources(dir)
		util.WriteStatus(fmt.Sprintf("%s resources created", dir))
		dc.createConnectionJava(dir)
		util.WriteStatus(fmt.Sprintf("%s java classes created", dir))
		util.WriteStatus(fmt.Sprintf("%s project creation completed", dir))
	} else {
		util.WriteStatus("--createConnection flag not found.")
	}

	if dc.Library {
		util.WriteStatus("--createLibrary flag found. Building library project")
		dir := dc.createLibrary()
		util.WriteStatus(fmt.Sprintf("%s created", dir))
		dc.createLibrarySubDirectories(dir)
		dc.createLibraryPOM(dir)
		util.WriteStatus(fmt.Sprintf("%s pom.xml file created", dir))
		dc.createLibraryResources(dir)
		util.WriteStatus(fmt.Sprintf("%s resources created", dir))
		dc.createLibraryJava(dir)
		util.WriteStatus(fmt.Sprintf("%s java classes created", dir))
		util.WriteStatus(fmt.Sprintf("%s project creation completed", dir))
	} else {
		util.WriteStatus("--createLibrary flag not found")
	}
}

func (dc *DataCollector) createLibraryJava(dir string) {
	project := strings.ToLower(dc.Name)
	name := strings.ReplaceAll(strings.ReplaceAll(dc.Name, "-", ""), ".", "")
	path := fmt.Sprintf("%s/%s", dir, "src/main/java/com.streamsets.pipeline")

	// Config classes
	f, size := util.GetFile("/assets/library/java/config/CoreConfigBean.java")
	b := make([]byte, size)
	_, err := f.Read(b)
	if err != nil {
		panic(err)
	}
	contents := string(b)
	contents = strings.ReplaceAll(contents, "@project@", project)
	contents = strings.ReplaceAll(contents, "@name@", name)

	err = os.WriteFile(fmt.Sprintf("%s/lib.%s.config/%sCoreConfigBean.java", path, project, name), []byte(contents), os.ModePerm)
	if err != nil {
		panic(err)
	}

	f, size = util.GetFile("/assets/library/java/config/Errors.java")
	b = make([]byte, size)
	_, err = f.Read(b)
	if err != nil {
		panic(err)
	}
	contents = string(b)
	contents = strings.ReplaceAll(contents, "@project@", project)
	contents = strings.ReplaceAll(contents, "@name@", name)

	err = os.WriteFile(fmt.Sprintf("%s/lib.%s.config/%sErrors.java", path, project, name), []byte(contents), os.ModePerm)
	if err != nil {
		panic(err)
	}

	f, size = util.GetFile("/assets/library/java/config/GroupConstants.java")
	b = make([]byte, size)
	_, err = f.Read(b)
	if err != nil {
		panic(err)
	}
	contents = string(b)
	contents = strings.ReplaceAll(contents, "@project@", project)
	contents = strings.ReplaceAll(contents, "@name@", name)

	err = os.WriteFile(fmt.Sprintf("%s/lib.%s.config/%sGroupConstants.java", path, project, name), []byte(contents), os.ModePerm)
	if err != nil {
		panic(err)
	}

	f, size = util.GetFile("/assets/library/java/config/Groups.java")
	b = make([]byte, size)
	_, err = f.Read(b)
	if err != nil {
		panic(err)
	}
	contents = string(b)
	contents = strings.ReplaceAll(contents, "@project@", project)
	contents = strings.ReplaceAll(contents, "@name@", name)

	err = os.WriteFile(fmt.Sprintf("%s/lib.%s.config/%sGroups.java", path, project, name), []byte(contents), os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Example Code
	f, size = util.GetFile("/assets/library/java/example/RecordHandler.java")
	b = make([]byte, size)
	_, err = f.Read(b)
	if err != nil {
		panic(err)
	}
	contents = string(b)
	contents = strings.ReplaceAll(contents, "@project@", project)
	contents = strings.ReplaceAll(contents, "@name@", name)

	err = os.WriteFile(fmt.Sprintf("%s/%s.example/%sRecordHandler.java", path, project, name), []byte(contents), os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Stage Code

	f, size = util.GetFile("/assets/library/java/stage/common/ConnectionVerifier.java")
	b = make([]byte, size)
	_, err = f.Read(b)
	if err != nil {
		panic(err)
	}
	contents = string(b)
	contents = strings.ReplaceAll(contents, "@project@", project)
	contents = strings.ReplaceAll(contents, "@name@", name)

	err = os.WriteFile(fmt.Sprintf("%s/stage/common.%s/%sConnectionVerifier.java", path, project, name), []byte(contents), os.ModePerm)
	if err != nil {
		panic(err)
	}

	f, size = util.GetFile("/assets/library/java/stage/common/Mgr.java")
	b = make([]byte, size)
	_, err = f.Read(b)
	if err != nil {
		panic(err)
	}
	contents = string(b)
	contents = strings.ReplaceAll(contents, "@project@", project)
	contents = strings.ReplaceAll(contents, "@name@", name)

	err = os.WriteFile(fmt.Sprintf("%s/stage/common.%s/%sMgr.java", path, project, name), []byte(contents), os.ModePerm)
	if err != nil {
		panic(err)
	}

	f, size = util.GetFile("/assets/library/java/stage/origin/DSource.java")
	b = make([]byte, size)
	_, err = f.Read(b)
	if err != nil {
		panic(err)
	}
	contents = string(b)
	contents = strings.ReplaceAll(contents, "@project@", project)
	contents = strings.ReplaceAll(contents, "@name@", name)

	err = os.WriteFile(fmt.Sprintf("%s/stage/origin.%s/%sDSource.java", path, project, name), []byte(contents), os.ModePerm)
	if err != nil {
		panic(err)
	}

	f, size = util.GetFile("/assets/library/java/stage/origin/Source.java")
	b = make([]byte, size)
	_, err = f.Read(b)
	if err != nil {
		panic(err)
	}
	contents = string(b)
	contents = strings.ReplaceAll(contents, "@project@", project)
	contents = strings.ReplaceAll(contents, "@name@", name)

	err = os.WriteFile(fmt.Sprintf("%s/stage/origin.%s/%sSource.java", path, project, name), []byte(contents), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func (dc *DataCollector) createLibraryResources(dir string) {
	// Image File
	img, size := util.GetFile("/assets/library/resources/example.png")

	b := make([]byte, size)

	_, err := img.Read(b)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(fmt.Sprintf("%s/%s/%s.png", dir, "src/main/resources", dc.Name), b, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// properties file(s)
	prop, size := util.GetFile("/assets/library/resources/data-collector-library.properties")

	b = make([]byte, size)

	_, err = prop.Read(b)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(fmt.Sprintf("%s/%s/%s", dir, "src/main/resources", "data-collector-library.properties"), b, os.ModePerm)
	if err != nil {
		panic(err)
	}

	prop, size = util.GetFile("/assets/library/resources/data-collector-library-bundle.properties")

	b = make([]byte, size)

	_, err = prop.Read(b)
	if err != nil {
		panic(err)
	}
	bundle := string(b)
	bundle = strings.ReplaceAll(bundle, "@name@", dc.Name)
	err = os.WriteFile(fmt.Sprintf("%s/%s/%s", dir, "src/main/resources", "data-collector-library-bundle.properties"), []byte(bundle), os.ModePerm)
	if err != nil {
		panic(err)
	}

	// upgrader file(s)
	verifier, size := util.GetFile("/assets/library/resources/upgrader/ConnectionVerifier.yaml")
	b = make([]byte, size)
	_, err = verifier.Read(b)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(fmt.Sprintf("%s/%s/%sConnectionVerifier.yaml", dir, "src/main/resources/upgrader", dc.Name), b, os.ModePerm)
	if err != nil {
		panic(err)
	}

	source, size := util.GetFile("/assets/library/resources/upgrader/Source.yaml")
	b = make([]byte, size)
	_, err = source.Read(b)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(fmt.Sprintf("%s/%s/%sSource.yaml", dir, "src/main/resources/upgrader", dc.Name), b, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func (dc *DataCollector) createLibraryPOM(dir string) {
	pomf, size := util.GetFile("/assets/library/pom.xml")

	defer pomf.Close()

	b := make([]byte, size)
	_, err := pomf.Read(b)
	if err != nil {
		panic(err)
	}
	start, end := "", ""
	if !dc.Connection {
		start = "<--"
		end = "-->"
	}
	pom := string(b)
	pom = strings.ReplaceAll(pom, "@project@", strings.ToLower(dc.Name))
	pom = strings.ReplaceAll(pom, "@name@", dc.Name)
	pom = strings.ReplaceAll(pom, "@version@", VERSION)
	pom = strings.ReplaceAll(pom, "@start@", start)
	pom = strings.ReplaceAll(pom, "@end@", end)
	err = os.WriteFile(fmt.Sprintf("%s/%s", dir, "pom.xml"), []byte(pom), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func (dc *DataCollector) createLibrarySubDirectories(base string) {
	resources := fmt.Sprintf("%s/src/main/resources/upgrader", base)
	util.CreateProjectDirectory(dc.Root, resources)

	jpipeline := fmt.Sprintf("%s/src/main/java/com.streamsets.pipeline", base)
	util.CreateProjectDirectory(dc.Root, jpipeline)

	project := strings.ToLower(dc.Name)
	jconfig := fmt.Sprintf("%s/lib.%s.config", jpipeline, project)
	util.CreateProjectDirectory(dc.Root, jconfig)

	jexample := fmt.Sprintf("%s/%s.example", jpipeline, project)
	util.CreateProjectDirectory(dc.Root, jexample)

	jstage := fmt.Sprintf("%s/stage", jpipeline)
	util.CreateProjectDirectory(dc.Root, jstage)

	jcommon := fmt.Sprintf("%s/common.%s", jstage, project)
	util.CreateProjectDirectory(dc.Root, jcommon)

	jorigin := fmt.Sprintf("%s/origin.%s", jstage, project)
	util.CreateProjectDirectory(dc.Root, jorigin)
}

func (dc *DataCollector) createLibrary() string {
	d := fmt.Sprintf("%s-%s", strings.ToLower(dc.Name), "lib")
	util.CreateProjectDirectory(dc.Root, d)

	return d
}

func (dc *DataCollector) createConnection() string {
	d := fmt.Sprintf("%s-%s", strings.ToLower(dc.Name), "connection")
	util.CreateProjectDirectory(dc.Root, d)

	return d
}

func (dc *DataCollector) createConnectionSubDirectories(base string) {
	java := fmt.Sprintf("%s/src/main/java/com.streamsets.pipeline.stage.common.%s", base, strings.ToLower(dc.Name))
	util.CreateProjectDirectory(dc.Root, java)
	resources := fmt.Sprintf("%s/src/main/resources/upgrader", base)
	util.CreateProjectDirectory(dc.Root, resources)
}

func (dc *DataCollector) createConnectionPOM(dir string) {
	f, err := pkger.Open("/assets/connection/pom.xml")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		panic(err)
	}
	b := make([]byte, stat.Size())
	_, err = f.Read(b)
	if err != nil {
		panic(err)
	}
	pom := string(b)
	pom = strings.ReplaceAll(pom, "@name@", strings.ToLower(dc.Name))
	pom = strings.ReplaceAll(pom, "@project@", dc.Name)
	pom = strings.ReplaceAll(pom, "@version@", VERSION)
	err = os.WriteFile(fmt.Sprintf("%s/%s", dir, "pom.xml"), []byte(pom), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func (dc *DataCollector) createConnectionResources(dir string) {
	f, err := pkger.Open("/assets/connection/resources/upgrader.yaml")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		panic(err)
	}

	b := make([]byte, stat.Size())
	_, err = f.Read(b)
	if err != nil {
		panic(err)
	}

	name := strings.ReplaceAll(dc.Name, "-", "")
	name = strings.ReplaceAll(name, ".", "")
	name = fmt.Sprintf("%s%s", name, "ConnectionUpgrader.yaml")
	err = os.WriteFile(fmt.Sprintf("%s/src/main/resources/upgrader/%s", dir, name), b, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func (dc *DataCollector) createConnectionJava(dir string) {
	project := strings.ToLower(dc.Name)
	name := strings.ReplaceAll(strings.ReplaceAll(dc.Name, "-", ""), ".", "")
	path := fmt.Sprintf("%s/src/main/java/com.streamsets.pipeline.stage.common.%s", dir, project)
	/*
	* Create Connection class
	 */

	conn, size := util.GetFile("/assets/connection/java/connection.java")

	defer conn.Close()

	connf := make([]byte, size)
	_, err := conn.Read(connf)
	if err != nil {
		panic(err)
	}

	contents := string(connf)
	contents = strings.ReplaceAll(strings.ReplaceAll(contents, "@name@", name), "@project@", project)

	conn_fn := fmt.Sprintf("%s%s", name, "Connection.java")
	err = os.WriteFile(fmt.Sprintf("%s/%s", path, conn_fn), []byte(contents), os.ModePerm)

	if err != nil {
		panic(err)
	}

	/*
	* Create ConnectionGroupConstants class
	 */

	cgc, size := util.GetFile("/assets/connection/java/connectiongroupconstants.java")
	defer cgc.Close()

	cgcf := make([]byte, size)
	_, err = cgc.Read(cgcf)
	if err != nil {
		panic(err)
	}

	cgcc := []byte(strings.ReplaceAll(strings.ReplaceAll(string(cgcf), "@project@", project), "@name@", name))
	cgc_fn := fmt.Sprintf("%s%s", name, "ConnectionGroupConstants.java")
	err = os.WriteFile(fmt.Sprintf("%s/%s", path, cgc_fn), cgcc, os.ModePerm)

	if err != nil {
		panic(err)
	}
	/*
	* Create ConnectionGroups
	 */

	groups, size := util.GetFile("/assets/connection/java/connectiongroups.java")
	defer groups.Close()

	groupsf := make([]byte, size)
	_, err = groups.Read(groupsf)
	if err != nil {
		panic(err)
	}

	groupsc := []byte(strings.ReplaceAll(strings.ReplaceAll(string(groupsf), "@project@", project), "@name@", name))
	groups_fn := fmt.Sprintf("%s%s", name, "ConnectionGroups.java")
	err = os.WriteFile(fmt.Sprintf("%s/%s", path, groups_fn), groupsc, os.ModePerm)

	if err != nil {
		panic(err)
	}
}
