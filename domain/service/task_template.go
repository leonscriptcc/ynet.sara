package service

import "github.com/leonscriptcc/ynet.sara/domain/entity"

// TaskTemplate 任务模板
type TaskTemplate interface {
	ImportTaskTemplate(dir string, plan entity.ProjectPlan) error
}
