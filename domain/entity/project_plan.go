package entity

// ProjectPlan 项目计划总表-计拆拆模板表格
type ProjectPlan struct {
	path  string
	Order Order
}

// Order 订单
type Order struct {
	Name       string
	Milestones []Milestone
}

// Milestone 里程碑/节点
type Milestone struct {
	Name       string
	WorkOrders []WorkOrder
}

// WorkOrder 工单
type WorkOrder struct {
	Name  string
	Tasks []Task
}

// Task 任务
type Task struct {
	Name      string
	Workload  string
	StartDate string
	EndDate   string
	Executor  string
}
