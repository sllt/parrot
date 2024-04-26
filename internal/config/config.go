package config

var (
	Version       = "v0.0.1"
	WireCmd       = "github.com/google/wire/cmd/wire@latest"
	ParrotCmd     = "github.com/sllt/parrot@latest"
	RepoProto     = "https://github.com/sllt/parrot-layout-proto.git"
	RepoAPI       = "https://github.com/sllt/parrot-layout-api.git"
	RunExcludeDir = ".git,.idea,tmp,vendor,.vscode"
	RunIncludeExt = "go,html,yaml,yml,toml,ini,json,xml,tpl,tmpl"
)
