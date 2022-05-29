package usercall

type ConstructCall struct {

	//Basic ID for the constructor
	ConstructorID string

	//Inputs from the Terminal
	codebase_path     string
	temperorydir_path string
	makefile_command  string

	//Inputs for database
	connection_url string
	database_type  string
	user_name      string
	password       string

	//build info
	buildrequired bool
}

type AnalyzeCall struct {
	AnalyzerID   string //Unique Id
	ConstructoID string //Id of the constructor, usefull in the metadata

	debug_level int //Debug level, taken from the user call

	run_commands string //Command which will be used to run the test file
	run_dir      string //Directory from where do we are required to run the tests
}

type Constructor interface {
	generateConstructorID()
	initialize()
	takeinput()
	generateCodeBaseCopy()
	buildCode()
	run()
}

type Analyzer interface {
	generateAnalyzerID()
	initailize()
	run()
	takeinput()
}

type displayUnit struct {
	displayID       string
	functionName    string
	VariableDetails string
}

type display interface {
	next()
	previous()
	getAll()
}
