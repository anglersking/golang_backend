package flag

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"
)

var (
	Config  string
	Debug   bool
	DataDir string
)

func init() {
	flag.StringVar(&Config, "c", "config.json", "-c <config file>")
	flag.BoolVar(&Debug, "debug", false, "--debug Using debug mode")
	flag.StringVar(&DataDir, "d", "./", "-d <data dir>")
}

func init() {
	testing.Init()
	flag.Parse()
	validationCheck(DataDir)
}

func validationCheck(dataDir string) {
	if !isDir(dataDir) {
		panic(fmt.Sprintf("%s not a directory", dataDir))
	}

	if !strings.HasSuffix(dataDir, "/") {
		dataDir = dataDir + "/"
	}
}

func isDir(dataDir string) bool {
	if dir, err := os.Stat(dataDir); err == nil && dir != nil {
		return dir.IsDir()
	}
	return false
}
