package employee

//Status 在职员工子状态筛选，可以查询多个状态。不同状态之间使用英文逗号分隔。
//
//2：试用期
//
//3：正式
//
//5：待离职
//
//-1：无状态
type Status string

const (
	// ProbationPeriod 试用期
	ProbationPeriod = Status("2")

	// Formal 正式
	Formal = Status("3")

	// ToBeResigned 待离职
	ToBeResigned = Status("5")

	// Unknown 无状态
	Unknown = Status("-1")
)
