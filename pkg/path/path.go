package path

import (
	"os"
	"path/filepath"
)

// postgre
var (
	PostgreHome = os.Getenv("POSTGRE_HOME")
)

// mate
var (
	PostgreMatePath    = filepath.FromSlash(PostgreHome + "/mate")
	PostgreScripts     = filepath.FromSlash(PostgreMatePath + "/scripts")
	PostgreStartScript = filepath.FromSlash(PostgreScripts + "/start-postgre.sh")
)
