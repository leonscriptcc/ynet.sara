package service

import "github.com/leonscriptcc/ynet.sara/domain/entity"

// TaskTemplate 任务模板
type TaskTemplate interface {
	ImportTaskTemplate(dir, emPath string, plan entity.ProjectPlan, roster map[string]string) error
}
