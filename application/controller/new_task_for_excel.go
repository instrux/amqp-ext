package controller

import (
	"github.com/gin-gonic/gin"
)

type NewTaskForExcelBody struct {
	SheetsDef []string `json:"sheets_name"`
}

func (c *controller) NewTaskForExcel(ctx *gin.Context) interface{} {
	var body NewTaskForExcelBody
	var err error
	if err = ctx.BindJSON(&body); err != nil {
		return err
	}
	var taskId string
	if taskId, err = c.dep.Excel.NewTask(body.SheetsDef); err != nil {
		return err
	}
	return gin.H{
		"task_id": taskId,
	}
}
