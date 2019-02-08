// Package conf implements access to gisrc.json files.
package conf

import (
	"encoding/json"
	"io/ioutil"

	"github.com/jychri/git-in-sync/pkg/flags"
	q "github.com/jychri/git-in-sync/pkg/quit"
)

// private

func read(f flags.Flags) ([]byte, q.Quit) {
	var bs []byte
	var err error

	bs, err = ioutil.ReadFile(f.Config)

	fmats := []string{"Cant' read file at (%v)\n", "Read file at (%v)\n"}
	q := q.Err(err, fmats, f.Config)

	return bs, q
}

func unmarshal(bs []byte, f flags.Flags) (Config, q.Quit) {
	var c Config
	var err error

	err = json.Unmarshal(bs, &c)

	fmats := []string{"Can't unmarshal JSON from (%v)\n", f.Config}
	q := q.Err(err, fmats, f.Config)

	return c, q
}

// public

// Config holds unmrashalled data from gisrc.json.
type Config struct {
	Bundles []struct {
		Path  string `json:"path"`
		Zones []struct {
			User      string   `json:"user"`
			Remote    string   `json:"remote"`
			Workspace string   `json:"workspace"`
			Repos     []string `json:"repositories"`
		} `json:"zones"`
	} `json:"bundles"`
}

// Init returns unmarshalled data from gisrc.json.
// f.Config is validated before reaching Init.
// flags.Init() verifies input with tilde.AbsUser()
func Init(f flags.Flags) (c Config) {

	bs, _ := read(f)
	c, _ = unmarshal(bs, f)

	return c
}
