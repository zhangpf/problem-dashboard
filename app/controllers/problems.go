package controllers

import (
	"github.com/revel/revel"
	//"problem-dashboard/app/models"
)

type Problems struct {
	App
}

func (c Problems) Index() revel.Result {
	//results, err := c.Txn.Select(models.Booking)
	return nil
}
