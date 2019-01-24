package main

import (
	"fmt"

	"github.com/jychri/git-in-sync/pkg/brf"
	"github.com/jychri/git-in-sync/pkg/conf"
	"github.com/jychri/git-in-sync/pkg/emoji"
	"github.com/jychri/git-in-sync/pkg/flags"
	"github.com/jychri/git-in-sync/pkg/repos"
	"github.com/jychri/git-in-sync/pkg/timer"
)

func initRun() (e emoji.Emoji, f flags.Flags, rs repos.Repos, t *timer.Timer) {

	// clear the screen
	emoji.ClearScreen()

	// initialize Timer, Emoji and Flags
	t = timer.Init()
	f = flags.Init()
	t.MarkMoment("init-flags")
	e = emoji.Init()
	t.MarkMoment("init-emoji")

	// print "start"
	brf.Printv(f, "%v start", e.Clapper)

	// print "flag(s) set..."
	if ft, err := t.GetMoment("init-flags"); err == nil {
		brf.Printv(f, "%v parsing flags", e.FlagInHole)
		brf.Printv(f, "%v [%v] flags set (%v) {%v / %v}", e.Flag, f.Count, f.Summary, ft.Split, ft.Start)
	}

	// print "emoji..."
	if et, err := t.GetMoment("init-emoji"); err == nil {
		brf.Printv(f, "%v initializing emoji", e.CrystalBall)
		brf.Printv(f, "%v [%v] emoji {%v / %v}", e.DirectHit, e.Count, et.Split, et.Start)
	}

	// print "reading ~/.gisrc.json"
	brf.Printv(f, "%v reading ~/.gisrc.json", e.Books)

	// initialize Config from ~/.gisrc.json
	c := conf.Init(f)
	t.MarkMoment("init-config")

	// print "read /Users/user/.gisrc.json..."
	brf.Printv(f, "%v read {%v / %v}", e.Book, t.Split(), t.Time())

	// print "parsing Repos..."
	brf.Printv(f, "parsing repos...")

	// initialize Repos
	rs = repos.Init(c)

	// print "parsed repos..."
	brf.Printv(f, "read repos...")

	return e, f, rs, t
}

func main() {
	e, f, rs, t := initRun()
	t.MarkMoment("OK")
	rs.VerifyDivs(e, f)
	fmt.Println("OK")
	// rs.verifyCloned(e, f)
	// rs.verifyRepos(e, f)
	// rs.verifyChanges(e, f)
	// rs.submitChanges(e, f)
}
