package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	s "strings"
)

var p = fmt.Println

type point struct {
	x, y int
}

func Stringtest() {
	p("Contents: ", s.Contains("test", "es"))
	p("Count: ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index: ", s.HasSuffix("test", "e"))
	p("Join: ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat: ", s.Repeat("a", 5))
	p("Replace: ", s.Replace("foo", "o", "0", -1))
	p("Replace: ", s.Replace("foo", "o", "0", 1))
	p("Split: ", s.Split("a-b-c-d-e", "-"))
	p("ToLower: ", s.ToLower("TEST"))
	p("ToUpper: ", s.ToUpper("test"))
	p()
	p("Len: ", len("Hello"))
	p("Char: ", "hello"[1])

	p := point{1, 2}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)

	fmt.Printf("%T\n", p)
	fmt.Printf("%t\n", true)
	fmt.Printf("%d\n", 123)

	fmt.Printf("%b\n", 14)
	fmt.Printf("%c\n", 33)

	fmt.Printf("%x\n", 456)
	fmt.Printf("%f\n", 78.9)

	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	fmt.Printf("%s\n", "\"String\"")
	fmt.Printf("%q\n", "\"String\"")
	fmt.Printf("%x\n", "hex this")

	fmt.Printf("%p\n", &p)
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	fmt.Printf("|%-6.2f|%-6.2f|", 1.2, 3.45)
	fmt.Printf("|%6s|%6s|", "foo", "b")

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")

	//RegExp
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("peach puch"))
	fmt.Println(r.FindStringIndex("peach puch"))
	fmt.Println(r.FindStringSubmatch("peach puch"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
