// Forked from https://git.dim13.org/itermcolors.git/tree/main.go
// Inspired by https://github.com/F1LT3R/itermcolors-to-hex/blob/master/index.js
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"howett.net/plist"
)

type Scheme struct {
	Name         string
	Ansi0        Color `plist:"Ansi 0 Color"`
	Ansi1        Color `plist:"Ansi 1 Color"`
	Ansi2        Color `plist:"Ansi 2 Color"`
	Ansi3        Color `plist:"Ansi 3 Color"`
	Ansi4        Color `plist:"Ansi 4 Color"`
	Ansi5        Color `plist:"Ansi 5 Color"`
	Ansi6        Color `plist:"Ansi 6 Color"`
	Ansi7        Color `plist:"Ansi 7 Color"`
	Ansi8        Color `plist:"Ansi 8 Color"`
	Ansi9        Color `plist:"Ansi 9 Color"`
	Ansi10       Color `plist:"Ansi 10 Color"`
	Ansi11       Color `plist:"Ansi 11 Color"`
	Ansi12       Color `plist:"Ansi 12 Color"`
	Ansi13       Color `plist:"Ansi 13 Color"`
	Ansi14       Color `plist:"Ansi 14 Color"`
	Ansi15       Color `plist:"Ansi 15 Color"`
	Background   Color `plist:"Background Color"`
	Badge        Color `plist:"Badge Color"`
	Bold         Color `plist:"Bold Color"`
	Cursor       Color `plist:"Cursor Color"`
	CursorGuide  Color `plist:"Cursor Guide Color"`
	CursorText   Color `plist:"Cursor Text Color"`
	Foreground   Color `plist:"Foreground Color"`
	Link         Color `plist:"Link Color"`
	SelectedText Color `plist:"Selected Text Color"`
	Selection    Color `plist:"Selection Color"`
}

type Color struct {
	Alpha float64 `plist:"Alpha Component"`
	Red   float64 `plist:"Red Component"`
	Green float64 `plist:"Green Component"`
	Blue  float64 `plist:"Blue Component"`
	Space string  `plist:"Color Space"`
}

func (c Color) String() string {
	r := int(math.MaxUint8 * c.Red)
	g := int(math.MaxUint8 * c.Green)
	b := int(math.MaxUint8 * c.Blue)
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

func name(s string) string {
	return strings.TrimSuffix(filepath.Base(s), filepath.Ext(s))
}

func exitErr(code int, err error) {
	fmt.Println(err)
	os.Exit(code)
}

func main() {
	file := flag.String("file", "", "Color scheme .itemcolors")
	tmpl := flag.String("tmpl", "scheme.tmpl", "Xresources template")
	out := flag.String("out", "-", "output")
	flag.Parse()

	body, err := ioutil.ReadFile(*file)
	if err != nil {
		exitErr(1, err)
	}

	var s Scheme
	if _, err := plist.Unmarshal(body, &s); err != nil {
		exitErr(1, err)
	}
	s.Name = name(*file)
	fd := os.Stdout

	if *out != "-" {
		fd, err = os.Create(*out)
		if err != nil {
			exitErr(1, err)
		}
	}

	t, err := template.ParseFiles(*tmpl)
	if err != nil {
		exitErr(1, err)
	}

	if err := t.Execute(fd, s); err != nil {
		exitErr(1, err)
	}
}
