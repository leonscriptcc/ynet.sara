package reference

import (
	"fmt"
	"github.com/leonscriptcc/ynet.sara/domain/entity"
	"github.com/xuri/excelize/v2"
	"time"
)

type ProjectPlan struct {
	nowTask entity.Task
}

func NewProjectPlan() *ProjectPlan {
	return &ProjectPlan{}
}

// LoadProjectPlan 从计拆拆模版中导出数据
func (p *ProjectPlan) LoadProjectPlan(path string) (plan entity.ProjectPlan, err error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return plan, err
	}
	defer func() {
		// release file descriptor
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// 读取计划中的所有数据
	rows, err := f.GetRows(entity.PLAN_SHEET)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 解析数据并插入结构体
	for i := 1; i < len(rows); i++ {
		for index, colCell := range rows[i] {
			if colCell == "" {
				continue
			}
			// 解析数据
			if err = p.analyseData(index, colCell, &plan); err != nil {
				break
			}
		}
	}

	return plan, err

}

// analyseData 解析数据
func (p *ProjectPlan) analyseData(i int, data string, plan *entity.ProjectPlan) (err error) {
	switch i {
	case entity.ORDER_INDEX:
		plan.Order.Name = data
	case entity.MILESTONE_INDEX:
		plan.Order.Milestones = append(plan.Order.Milestones, entity.Milestone{Name: data})
	case entity.WORKORDER_INDEX:
		ml := len(plan.Order.Milestones) - 1
		plan.Order.Milestones[ml].WorkOrders =
			append(plan.Order.Milestones[ml].WorkOrders, entity.WorkOrder{Name: data})
	case entity.TASK_INDEX:
		p.nowTask = entity.Task{Name: data}
	case entity.WORKLOAD_INDEX:
		p.nowTask.Workload = data
	case entity.STARTDATE_INDEX:
		t, err := time.Parse(entity.PLAN_DATE, data)
		if err != nil {
			return err
		}
		p.nowTask.StartDate = t.Format(entity.TASK_DATE)
	case entity.ENDDATE_INDEX:
		t, err := time.Parse(entity.PLAN_DATE, data)
		if err != nil {
			return err
		}
		p.nowTask.EndDate = t.Format(entity.TASK_DATE)
	case entity.EXECUTOR_INDEX:
		p.nowTask.Executor = data
		ml := len(plan.Order.Milestones) - 1
		wl := len(plan.Order.Milestones[ml].WorkOrders) - 1
		plan.Order.Milestones[ml].WorkOrders[wl].Tasks =
			append(plan.Order.Milestones[ml].WorkOrders[wl].Tasks, p.nowTask)
	}
	return err
}
