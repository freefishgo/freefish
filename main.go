package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/template"
)

// 实际中应该用更好的变量名
var (
	h bool

	new         bool
	view        bool
	controller  bool
	scontroller bool
)

const VERSION = "1.00"

func init() {
	flag.BoolVar(&h, "-h", true, "freeFishGo 帮助信息")

	flag.BoolVar(&new, "new", false, `创建一个新的mvc项目 具体有:
freefish new [ProjectName]                :在当前目录下创建mvc项目
freefish new [ProjectName] -path [dirPath]:在当前目录dirPath创建mvc项目
freefish new -gopath [ProjectName]        :在GOPATH下创建一个新的mvc项目`)

	flag.BoolVar(&view, "-v", false, `在freefish生成的项目中操作视图 具体命令有:
freefish -v check ：检查Mvc视图文件是否存在，打印缺视图的控制器和视图,仅供参考
freefish -v create：遍历Mvc控制器文件，创建缺失的视图,仅供参考`)

	flag.BoolVar(&controller, "-c", false, `在freefish生成的项目中控制器 具体命令有:
freefish -c [controllerName] ：在controllers文件夹下生成 controllerName+"Controller" 控制器`)
	flag.BoolVar(&scontroller, "-cs", false, `在freefish生成的项目中http状态处理控制器 具体命令有:
freefish -c [StatusControllerName] ：在controllers文件夹下生成 controllerName+"StatusController" 控制器`)

	// 改变默认的 Usage
	flag.Usage = usage
}

var ProjectName string
var WorkDir string
var importPath string

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
		if lens == 3 && os.Args[2][0] != '-' {
			path, _ := filepath.Abs(os.Args[2])
			WorkDir = path
			importPath = os.Args[2]
			ProjectName = os.Args[2]
			createMvc(os.Args[2])
		} else if lens == 4 && os.Args[2] == "-gopath" {
			GOPATH := os.Getenv("GOPATH")
			path := filepath.Join(GOPATH, "src", os.Args[3])
			WorkDir = path
			ProjectName = os.Args[3]
			importPath = ProjectName
			createMvc(os.Args[3])
		} else if (lens == 5 || lens == 4) && os.Args[3] == "-path" {
			tmp, _ := filepath.Abs("")
			if lens == 4 {
				os.Args = append(os.Args, "")
			}
			path, _ := filepath.Abs(filepath.Join(os.Args[4], os.Args[2]))
			WorkDir = path
			importPath = strings.Replace(filepath.Join(filepath.Base(tmp), os.Args[4], os.Args[2]), "\\", "/", -1)
			ProjectName = os.Args[2]
			createMvc(os.Args[2])
		} else {
			flag.Usage()
		}
	case "-v":
		if lens == 3 && os.Args[2] == "check" {
			viewCheck("", "", os.Args[2])
		} else if lens == 3 && os.Args[2] == "create" {
			viewCheck("", "", os.Args[2])
		} else {
			flag.Usage()
		}
	case "-c":
		if lens == 3 {
			path := filepath.Join("controllers", os.Args[2]+"Controller.go")
			if b, err := pathExists(path); err == nil {
				if b {
					log.Println("freeFish:->Controller:" + os.Args[2] + "创建失败,由于已有" + os.Args[2] + "Controller.go已经存在")
				} else {
					if f, err := os.Create(path); err == nil {
						defer f.Close()
						f.Write([]byte(strings.Replace(strings.Replace(controllerText, "{{[Controller]}}", os.Args[2]+"Controller", -1), "{{[Name]}}", os.Args[2], -1)))
						log.Println("freeFish:->Controller:" + os.Args[2] + "Controller创建成功,文件地址为:" + path)
					} else {
						panic(err)
					}
				}
			} else {
				panic(err)
			}
		} else {
			flag.Usage()
		}
	case "-cs":
		if lens == 3 {
			path := filepath.Join("controllers", os.Args[2]+"StatusController.go")
			if b, err := pathExists(path); err == nil {
				if b {
					log.Println("freeFish:->Controller:" + os.Args[2] + "创建失败,由于已有" + os.Args[2] + "StatusController.go已经存在")
				} else {
					if f, err := os.Create(path); err == nil {
						defer f.Close()
						f.Write([]byte(strings.Replace(strings.Replace(stateCodeControllerText, "{{[Controller]}}", os.Args[2]+"StatusCodeController", -1), "{{[Name]}}", os.Args[2], -1)))
						log.Println("freeFish:->Controller:" + os.Args[2] + "StatusCodeController创建成功,文件地址为:" + path)
					} else {
						panic(err)
					}
				}
			} else {
				panic(err)
			}
		} else {
			flag.Usage()
		}
	default:
		flag.Usage()

	}
}

//调用os.MkdirAll递归创建文件夹
func createFile(filePath string) error {
	if !isExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func isExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func viewCheck(viewDir string, controllerDir string, ty string) error {
	if viewDir == "" {
		viewDir = "views"
	}
	if controllerDir == "" {
		controllerDir = "controllers"
	}
	//检测目录正确性
	if srcInfo, err := os.Stat(viewDir); err != nil {
		fmt.Println(err.Error())
		return err
	} else {
		if !srcInfo.IsDir() {
			e := errors.New("srcPath不是一个正确的目录！")
			fmt.Println(e.Error())
			return e
		}
	}
	//检测目录正确性
	if srcInfo, err := os.Stat(controllerDir); err != nil {
		fmt.Println(err.Error())
		return err
	} else {
		if !srcInfo.IsDir() {
			e := errors.New("srcPath不是一个正确的目录！")
			fmt.Println(e.Error())
			return e
		}
	}

	err := filepath.Walk(controllerDir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if !f.IsDir() {
			f, _ := os.Open(path)
			defer f.Close()
			lineIndex := 0
			upCount := 0
			offCount := 0
			tmpStr := ""
			rd := bufio.NewReader(f)
			for {
				line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
				lineIndex++
				if err != nil || io.EOF == err {
					break
				}
				line = strings.Trim(line, "\n")
				line = strings.Trim(line, "\t")
				line = strings.Trim(line, "\n")
				line = strings.Trim(line, "\t")
				line = strings.Trim(line, " ")
				r := regexp.MustCompile(`^[//]`)
				b := r.MatchString(line)
				if b || len(line) == 0 {
					continue
				}
				line = strings.Split(line, "//")[0]
				tmpStr += line + " "
				upCount += strings.Count(line, "{")
				offCount += strings.Count(line, "}")
				r = regexp.MustCompile(`\.UseTplPath[\ ]?\([\ ]?(.*?)[\ ]?\)`)
				if r.MatchString(line) {
					sl := r.FindAllStringSubmatch(line, -1)
					r = regexp.MustCompile(`func[\ ]?\([\w+$]+[\ ]+\*[\ ]?([\w+$]+)[\ ]?\)[\ ]?([\w+$]+)\(.*?\)`)
					for _, v := range sl {
						s := r.FindAllStringSubmatch(tmpStr, -1)
						v[1] = strings.Trim(v[1], `"`)
						if v[1] == "" {
							if len(s) > 0 {
								controllerIndex := strings.Index(strings.ToLower(s[0][1]), "controller")
								viewPath := filepath.Join(viewDir, s[0][1][0:controllerIndex], replaceActionName(s[0][2])+".fish")
								if b, _ := pathExists(viewPath); !b {
									switch ty {
									case "check":
										log.Println("freeFish:->路径:" + path + " Controller:" + s[0][1] + " Action:" + s[0][2] + " 缺失视图:" + filepath.Join(s[0][1][0:controllerIndex], replaceActionName(s[0][2])+".fish") + " 行号:" + strconv.Itoa(lineIndex))
									case "create":
										createFile(filepath.Dir(viewPath))
										f, err := os.Create(viewPath)
										f.Write([]byte(htmlText))
										defer f.Close()
										if err != nil {
											panic(err)
										}
										log.Println("freeFish:->路径:" + path + " Controller:" + s[0][1] + " Action:" + s[0][2] + " 成功创建视图:" + filepath.Join(s[0][1][0:controllerIndex], replaceActionName(s[0][2])+".fish") + " 行号:" + strconv.Itoa(lineIndex))
									}
								}
							}
						} else {
							viewPath := filepath.Join(viewDir, v[1])
							if b, _ := pathExists(viewPath); !b {
								switch ty {
								case "check":
									if len(s) > 0 && len(s[0]) > 2 {
										log.Println("freeFish:->路径:" + path + " Controller:" + s[0][1] + " Action:" + s[0][2] + " 缺失视图:" + v[1] + " 行号:" + strconv.Itoa(lineIndex))
									} else {
										log.Println("freeFish:->路径:" + path + " 可能缺失视图:" + v[1] + " 行号:" + strconv.Itoa(lineIndex))
									}
								case "create":
									createFile(filepath.Dir(viewPath))
									f, err := os.Create(viewPath)
									f.Write([]byte(htmlText))
									defer f.Close()
									if err != nil {
										panic(err)
									}
									if len(s) > 0 && len(s[0]) > 2 {
										log.Println("freeFish:->路径:" + path + " Controller:" + s[0][1] + " Action:" + s[0][2] + " 成功创建视图:" + v[1] + " 行号:" + strconv.Itoa(lineIndex))
									} else {
										log.Println("freeFish:->路径:" + path + "   成功创建视图:" + v[1] + " 行号:" + strconv.Itoa(lineIndex))
									}
								}
							}
						}
					}
				}
				if upCount == offCount {
					upCount = 0
					offCount = 0
					tmpStr = ""
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf(err.Error())
	}
	return err

}

type HttpMethod string

const (
	MethodGet     HttpMethod = "GET"
	MethodHead    HttpMethod = "HEAD"
	MethodPost    HttpMethod = "POST"
	MethodPut     HttpMethod = "PUT"
	MethodPatch   HttpMethod = "PATCH" // RFC 5789
	MethodDelete  HttpMethod = "DELETE"
	MethodConnect HttpMethod = "CONNECT"
	MethodOptions HttpMethod = "OPTIONS"
	MethodTrace   HttpMethod = "TRACE"
	htmlText      string     = `<!DOCTYPE html>
<html>
<head>
    <title>FreeFishGo</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<body>
	<h1>欢迎使用 fishFreeGoMvc</h1>
</body>
</html>`
	controllerText string = `package controllers

import (
	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

type {{[Controller]}} struct {
	mvc.Controller
}

func init() {
	mvc.AddHandlers(&{{[Controller]}}{})
}

func ({{[Name]}} *{{[Controller]}}) Index() {

}

// 控制器执行前调用
func ({{[Name]}} *{{[Controller]}}) Prepare() {
	//log.Println("子类的Prepare")
}

// 控制器结束时调用
func ({{[Name]}} *{{[Controller]}}) Finish() {
	//log.Println("子类的Finish")
}`
	stateCodeControllerText = `package controllers

import (
	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

func init() {
	mvc.SetStatusCodeHandlers(&{{[Controller]}}{})
}

type {{[Name]}}StatusCodeController struct {
	mvc.StatusCodeController
}

// 500 错误处理函数
func ({{[Name]}} *{{[Controller]}}) Error500() {
	{{[Name]}}.StatusCodeController.Error500()
}

// 403 处理函数
func ({{[Name]}} *{{[Controller]}}) Forbidden403() {
	{{[Name]}}.StatusCodeController.Forbidden403()
}

// 404 处理函数
func ({{[Name]}} *{{[Controller]}}) NotFind404() {
	{{[Name]}}.StatusCodeController.NotFind404()
}
`
)

func replaceActionName(actionName string) string {
	tmp := actionName
	actionName = strings.ToUpper(actionName)
	httpMethodList := []HttpMethod{MethodPost,
		MethodConnect, MethodDelete,
		MethodGet, MethodHead, MethodOptions,
		MethodPatch, MethodPut, MethodTrace}
	for _, v := range httpMethodList {
		f := regexp.MustCompile(string(v) + "$")
		if f.MatchString(actionName) {
			index := f.FindAllStringIndex(actionName, -1)
			return tmp[0:index[0][0]]
		}
	}
	return tmp

}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func createMvc(mvcName string) {
	log.Println("生成MVC项目:" + mvcName + " 中.....")
	GOPATH := os.Getenv("GOPATH")
	if GOPATH == "" {
		log.Println("未设置GOPATH环境变量")
		log.Println("创建项目失败")
		return
	}
	log.Println(filepath.Join(GOPATH, "src/github.com/freefishgo/freefish/template"))
	if err := os.Mkdir(WorkDir, os.ModeDir|os.ModePerm); err != nil {
		panic(err.Error())
	}
	if err := CopyDir(filepath.Join(GOPATH, "src/github.com/freefishgo/freefish/template"), WorkDir); err != nil {
		log.Println("MVC项目:" + mvcName + " 生成失败.....失败原因为:" + err.Error())
	} else {
		log.Println("MVC项目:" + mvcName + " 生成成功.....请查看目录:" + WorkDir)
	}
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
			fmt.Println("创建文件：" + destNewPath + " 成功")
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
	if filepath.Ext(src) == ".go" || filepath.Ext(src) == ".mod" {
		t, e := template.New(src).Delims("{{[", "]}}").Parse(string(b))
		if e != nil {
			log.Println(e.Error())
			dstFile.Write(b)
			return
		} else {
			data := map[string]interface{}{}
			data["ProjectName"] = importPath
			data["Chdir"] = WorkDir
			t.Execute(dstFile, data)

			defer dstFile.Close()
			return
		}
		return
	} else {
		dstFile.Write(b)
	}
	return
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
	s += strings.Replace(usage, "\n", "\n    \t", -1)
	fmt.Fprint(os.Stderr, s, "\n")
}
