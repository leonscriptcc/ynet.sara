package reference

import (
	"fmt"
	"github.com/leonscriptcc/ynet.sara/domain/entity"
	"github.com/xuri/excelize/v2"
)

// TaskTemplate 任务模板
type TaskTemplate struct {
}

func NewTaskTemplate() *TaskTemplate {
	return &TaskTemplate{}
}

// ImportTaskTemplate 数据导入文件模板
func (t *TaskTemplate) ImportTaskTemplate(dir, emPath string, plan entity.ProjectPlan) (err error) {
	for _, milestone := range plan.Order.Milestones {
		// 里程碑不是录入参数,订单开始录入
		for _, order := range milestone.WorkOrders {
			// 打开样本excel
			f, err := excelize.OpenFile(emPath)
			if err != nil {
				return err
			}
			// 录入任务
			for i, task := range order.Tasks {
				dataArray := []interface{}{task.Name, nil, nil, task.Executor, nil, task.StartDate, task.EndDate, task.Workload, task.Name}
				err = f.SetSheetRow(entity.TASK_SHEET, fmt.Sprintf("A%d", i+2), &dataArray)
				if err != nil {
					fmt.Println(task.Name + " import fail")
				}
			}
			// 文件另存
			if err := f.SaveAs(dir + order.Name + ".xlsx"); err != nil {
				fmt.Println(dir + order.Name + ".xlsx save fail")
			}
		}

	}
	return err
}
