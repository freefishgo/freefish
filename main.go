package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// 实际中应该用更好的变量名
var (
	h bool

	new bool
	mvc bool
)

func init() {
	flag.BoolVar(&h, "-h", false, "freeFishGo 帮助信息")

	flag.BoolVar(&new, "new", false, "创建一个新的项目 如:freefish new [ProjectName]")

	flag.BoolVar(&mvc, "mvc", false, "创建一个新的mvc项目 如:freefish new mvc [ProjectName]")

	// 改变默认的 Usage
	flag.Usage = usage
}

func main() {
	lens := len(os.Args)
	if lens < 2 {
		flag.Usage()
		return
	}
	switch os.Args[1] {
	case "-h":
		cmdHelp()
		return
	case "new":
		if lens == 3 {
			return
		} else if os.Args[2] == "mvc" && lens == 4 {
			createMvc(os.Args[3])
		} else {
			flag.Usage()
		}

	}
}

func createMvc(mvcName string) {
	if err := os.Mkdir(filepath.Join(os.Args[0], mvcName), os.ModeDir); err != nil {
		panic(err.Error())
	} else {
		log.Println("mkdir Suff")
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `freefishgo version: `+freeFish.VERSION+`
命令错误
Usage: freefish h look help

Options:
`)
	flag.VisitAll(MyVisit)

}

func cmdHelp() {
	fmt.Fprintf(os.Stderr, `freefishgo version: `+freeFish.VERSION+`
Usage: freefish h look help

Options:
`)
	flag.VisitAll(MyVisit)
}

func MyVisit(flag2 *flag.Flag) {
	s := fmt.Sprintf("  -%s", flag2.Name) // Two spaces before -; see next two comments.
	name, usage := flag.UnquoteUsage(flag2)
	s = strings.Replace(s, "-", " ", 1)
	if len(name) > 0 {
		s += " " + name
	}
	// Boolean flags of one ASCII letter are so common we
	// treat them specially, putting their usage on the same line.
	if len(s) <= 4 { // space, space, '-', 'x'.
		s += "\t"
	} else {
		// Four spaces before the tab triggers good alignment
		// for both 4- and 8-space tab stops.
		s += "\n    \t"
	}
	s += strings.ReplaceAll(usage, "\n", "\n    \t")
	fmt.Fprint(os.Stderr, s, "\n")
}
