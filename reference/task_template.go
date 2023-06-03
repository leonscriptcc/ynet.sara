package reference

import (
	"github.com/leonscriptcc/ynet.sara/domain/entity"
	"log"
)

// TaskTemplate 任务模板
type TaskTemplate struct {
}

func (t *TaskTemplate) ImportTaskTemplate(dir string, plan entity.ProjectPlan) (err error) {
	for _, milestone := range plan.Order.Milestones {
		// 里程碑不是录入参数,订单开始录入
		for _, order := range milestone.WorkOrders {
			//TODO 创建excel
			for _, task := range order.Tasks {
				//TODO 录入任务
				log.Println(task)
			}
		}

	}
	return err
}
