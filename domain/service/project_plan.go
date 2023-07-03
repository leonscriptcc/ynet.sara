package service

import "github.com/leonscriptcc/ynet.sara/domain/entity"

type ProjectPlan interface {
	// LoadProjectPlanFromFile 从计拆拆模版中导出数据
	LoadProjectPlanFromFile(filePath string) (entity.ProjectPlan, error)

	// LoadRoaster 读取人员信息
	LoadRoaster() (map[string]string, error)
}
