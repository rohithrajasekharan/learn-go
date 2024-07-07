package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:  ", s.Count("test", "t"))
	p("Has Prefix:  ", s.HasPrefix("test", "tes"))
	p("Has Suffix:  ", s.HasSuffix("test", "est"))
	p("Index:  ", s.Index("test", "s"))
	p("Join:  ", s.Join([]string{"test", "string"}, "-"))
	p("Repeat:  ", s.Repeat("a", 5))
	p("Replace:  ", s.Replace("test", "t", "p", -1))
	p("Replace:  ", s.Replace("test", "t", "p", 1))
	p("Split:  ", s.Split("test-file", "-"))
	p("ToLower:  ", s.ToLower("TesT"))
	p("ToUpper:  ", s.ToUpper("TesT"))
}
