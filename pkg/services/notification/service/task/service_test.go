// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package task

import (
	"testing"

	"openpitrix.io/logger"
)

//func TestNewService(t *testing.T) {
//
//	//taskservice := NewService()
//
//	//go taskservice.ExtractTasks()
//	//go taskservice.HandleTask("A")
//	//go taskservice.HandleTask("B")
//	////
//	//for {
//	//	//println("...")
//	//	time.Sleep(2 * time.Second)
//	//}
//}

func TestGetTaskbyID(t *testing.T) {
	tasksc := &taskService{}
	task, _ := tasksc.getTaskbyID("task-6J1BEDx9wJ94")
	logger.Infof(nil, task.EmailAddr)
}

func TestGetTaskwithNfContentbyID(t *testing.T) {
	tasksc := &taskService{}
	task, _ := tasksc.GetTaskwithNfContentbyID("task-mqY0kxG9yl98")
	logger.Infof(nil, "EmailAddr=[%+s]", task.EmailAddr)
}

func TestUpdateStatusById(t *testing.T) {
	tasksc := &taskService{}
	task, _ := tasksc.GetTaskwithNfContentbyID("task-QvQEG9n5BkZO")
	tasksc.UpdateStatus2SendingByIds(*task)
}