package main

import (
	"flag"
	"github.com/leonscriptcc/ynet.sara/domain/service"
	"github.com/leonscriptcc/ynet.sara/reference"
	"log"
)

func main() {
	//声明变量用于接收命令行参数
	var dir string
	var planPath string

	// 声明
	flag.StringVar(&planPath, "plan", "", "计拆拆模板路径")
	flag.StringVar(&dir, "dir", "", "生成模板文件夹")
	// 开始解析命令行参数
	flag.Parse()
	if dir == "" || planPath == "" {
		log.Println("Illegal parameter")
		return
	}

	// 初始化
	var (
		projectPlan  service.ProjectPlan
		taskTemplate service.TaskTemplate
	)
	projectPlan = reference.NewProjectPlan()
	taskTemplate = reference.NewTaskTemplate()

	// 读取数据
	plan, err := projectPlan.LoadProjectPlan(planPath)
	if err != nil {
		log.Println("LoadProjectPlan fail " + err.Error())
		return
	}

	// 导入数据
	err = taskTemplate.ImportTaskTemplate(dir,
		"./example.xlsx", plan)
	if err != nil {
		log.Println("ImportTaskTemplate fail " + err.Error())
		return
	}
}
