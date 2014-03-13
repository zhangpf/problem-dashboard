package models

import (
	"github.com/revel/revel"
)

type Problem struct {
	ProblemId			int
	Title				string
	HustojId			int
}

func (problem *Problem) Validate(v *revel.Validation) {
	v.Check(problem.Title,
		revel.Required{},
		revel.MaxSize{50},
	)

	v.Check(problem.HustojId,
		revel.Required{},
	)
}
