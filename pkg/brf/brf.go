package brf

import (
	"bytes"
	"errors"
	"fmt"
	"os/user"
	"path"
	"strings"

	"github.com/jychri/git-in-sync/pkg/flags"
)

// Printv calls prints to standard output if not running in "oneline" mode.
func Printv(f flags.Flags, s string, z ...interface{}) (err error) {

	if f.Mode != "oneline" {
		fmt.Println(fmt.Sprintf(s, z...))
		return
	}

	return errors.New("N/A")
}

// Single returns a string slice with no duplications.
func Single(ssl []string) (sl []string) {

	smap := make(map[string]bool)

	for i := range ssl {
		if smap[ssl[i]] == true {
		} else {
			smap[ssl[i]] = true
			sl = append(sl, ssl[i])
		}
	}

	return sl
}

// Summary returns a set length string summarizing the contents of a string slice.
func Summary(sl []string, l int) string {
	if len(sl) == 0 {
		return ""
	}

	var csl []string // check slice
	var b bytes.Buffer

	for _, s := range sl {
		lc := len(strings.Join(csl, ", ")) // (l)ength(c)heck
		switch {
		case lc <= l-10 && len(s) <= 20: //
			csl = append(csl, s)
		case lc <= l && len(s) <= 12:
			csl = append(csl, s)
		}
	}

	b.WriteString(strings.Join(csl, ", "))

	if len(sl) != len(csl) {
		b.WriteString("...")
	}

	return b.String()
}

// First returns the first line from a multi line string.
func First(s string) string {
	lines := strings.Split(strings.TrimSuffix(s, "\n"), "\n")

	if len(lines) >= 1 {
		return lines[0]
	} else {
		return ""
	}
}

// Relative returns a path relative to the current user
func Relative(s string) (t string, err error) {
	var u *user.User

	u, err = user.Current()

	if err != nil {
		return "", errors.New("Unable to identify current user")
	}

	t = strings.TrimPrefix(s, "~/")

	if t != s {
		t = strings.Join([]string{u.HomeDir, "/", t}, "")
		return path.Clean(t), nil
	}

	return s, nil
}