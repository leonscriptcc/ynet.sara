package main

import (
	"github.com/leonscriptcc/ynet.sara/reference"
	"log"
)

func main() {
	p := reference.NewProjectPlan()
	plan, err := p.LoadProjectPlan("/Users/leonscript/Documents/YNET/sara/project.xlsm")
	if err != nil {
		log.Println(plan)
		log.Println(err)
	}
}
