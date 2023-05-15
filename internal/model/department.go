package model

type DepartmentInput struct {
	Id             int
	DepartmentName string
	IdList         []int
}

type DepartmentOutput struct {
	Id             int
	DepartmentName string
}
