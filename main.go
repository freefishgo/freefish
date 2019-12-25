package main

import (
	"errors"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
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

const VERSION = "1.00"

func init() {
	flag.BoolVar(&h, "-h", false, "freeFishGo 帮助信息")

	flag.BoolVar(&new, "new", false, "创建一个新的项目 如:freefish new [ProjectName]")

	flag.BoolVar(&mvc, "mvc", false, "创建一个新的mvc项目 如:freefish new mvc [ProjectName]")

	// 改变默认的 Usage
	flag.Usage = usage
}

var ProjectName string

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
			ProjectName = os.Args[2]
			return
		} else if os.Args[2] == "mvc" && lens == 4 {
			ProjectName = os.Args[3]
			createMvc(os.Args[3])
		} else {
			flag.Usage()
		}

	}
}

func createMvc(mvcName string) {
	GOPATH := os.Getenv("GOPATH")
	if GOPATH == "" {
		log.Println("未设置GOPATH环境变量")
		log.Println("创建项目失败")
		return
	}
	path := filepath.Join(GOPATH, mvcName)

	log.Println(filepath.Join(GOPATH, "src/github.com/freefishgo/freefish/template"))
	if err := os.Mkdir(path, os.ModeDir); err != nil {
		panic(err.Error())
	} else {
		log.Println("mkdir Suff")
	}
	CopyDir(filepath.Join(GOPATH, "src/github.com/freefishgo/freefish/template"), path)

	//if err := os.Mkdir(path, os.ModeDir); err != nil {
	//	panic(err.Error())
	//} else {
	//	log.Println("mkdir Suff")
	//}
}

/**
 * 拷贝文件夹,同时拷贝文件夹中的文件
 * @param srcPath  		需要拷贝的文件夹路径: D:/test
 * @param destPath		拷贝到的位置: D:/backup/
 */
func CopyDir(srcPath string, destPath string) error {
	//检测目录正确性
	if srcInfo, err := os.Stat(srcPath); err != nil {
		fmt.Println(err.Error())
		return err
	} else {
		if !srcInfo.IsDir() {
			e := errors.New("srcPath不是一个正确的目录！")
			fmt.Println(e.Error())
			return e
		}
	}
	if destInfo, err := os.Stat(destPath); err != nil {
		fmt.Println(err.Error())
		return err
	} else {
		if !destInfo.IsDir() {
			e := errors.New("destInfo不是一个正确的目录！")
			fmt.Println(e.Error())
			return e
		}
	}

	err := filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if !f.IsDir() {
			path := strings.Replace(path, "\\", "/", -1)
			srcPath = strings.Replace(srcPath, "\\", "/", -1)
			destPath = strings.Replace(destPath, "\\", "/", -1)
			destNewPath := strings.Replace(path, srcPath, destPath, -1)
			fmt.Println("复制文件:" + path + " 到 " + destNewPath)
			copyFile(path, destNewPath)
		}
		return nil
	})
	if err != nil {
		fmt.Printf(err.Error())
	}
	return err
}

//生成目录并拷贝文件
func copyFile(src, dest string) (w int64, err error) {
	if filepath.Ext(src) == "" || strings.ToLower(filepath.Ext(src)) == ".exe" {
		return
	}
	//分割path目录
	destSplitPathDirs := strings.Split(dest, "/")

	//检测时候存在目录
	destSplitPath := ""
	for index, dir := range destSplitPathDirs {
		if index < len(destSplitPathDirs)-1 {
			destSplitPath = destSplitPath + dir + "/"
			b, _ := pathExists(destSplitPath)
			if b == false {
				fmt.Println("创建目录:" + destSplitPath)
				//创建目录
				err := os.Mkdir(destSplitPath, os.ModePerm)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
	dstFile, err := os.Create(dest)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	b, _ := ioutil.ReadFile(src)
	t, e := template.New(src).Delims("{{[", "]}}").Parse(string(b))
	if e != nil {
		log.Println(e.Error())
		dstFile.Write(b)
		return
	} else {
		data := map[string]interface{}{}
		data["ProjectName"] = ProjectName
		t.Execute(dstFile, data)

		defer dstFile.Close()
		return
	}

}

//检测文件夹路径时候存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func usage() {
	fmt.Fprintf(os.Stderr, `freefishgo version: `+VERSION+`
命令错误
Usage: freefish h look help

Options:
`)
	flag.VisitAll(MyVisit)

}

func cmdHelp() {
	fmt.Fprintf(os.Stderr, `freefishgo version: `+VERSION+`
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
