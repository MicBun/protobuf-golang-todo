package schema

type TodoCreateOneBody struct {
	Task string `json:"task" validate:"required"`
}

type TodoFindAllQuery struct {
	Limit  *int  `query:"limit"`
	Offset *int  `query:"offset"`
	Status *bool `query:"status"`
}

type TodoGetOneQuery struct {
	ID uint `query:"id" validate:"required"`
}

type TodoUpdateOneQuery struct {
	ID uint `query:"id" validate:"required"`
}

type TodoUpdateOneBody struct {
	Status bool `json:"status" validate:"required"`
}

type TodoDeleteOneQuery struct {
	ID uint `query:"id" validate:"required"`
}
