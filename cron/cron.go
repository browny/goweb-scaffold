// Package cron defines those tasks running in the background
package cron

import (
	"fmt"

	"goweb-scaffold/config"
	"goweb-scaffold/logger"

	"github.com/robfig/cron"
)

var GlobalCron *cron.Cron

// TaskRunner manages the background tasks
type TaskRunner struct {
	Ctx *config.AppContext `inject:""`
}

// GlobalRun runs cluster-wise background task, only one node is allowed to run this
func (tr *TaskRunner) GlobalRun() {
	logger.Debugf("Run global taskRunner: Env[%s]", tr.Ctx.Env)

	GlobalCron = cron.New()

	GlobalCron.AddFunc(everyNminutes(1), func() {
		logger.Debug("cron job run every 1 minute")
	})

	GlobalCron.Start()
}

func everyNminutes(num int) string {
	return fmt.Sprintf("@every %dm", num)
}
