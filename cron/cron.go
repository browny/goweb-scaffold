// Package cron defines those tasks running in the background
package cron

import (
	"fmt"

	log "github.com/cihub/seelog"
	"github.com/robfig/cron"
	"goweb-scaffold/config"
)

var GlobalCron *cron.Cron

// TaskRunner manages the background tasks
type TaskRunner struct {
	Config *config.GlobalConfig `inject:""`
}

// GlobalRun runs cluster-wise background task, only one node is allowed to run this
func (tr *TaskRunner) GlobalRun() {
	log.Debugf("Run global taskRunner: projectID[%s]", tr.Config.ProjectId)

	GlobalCron = cron.New()

	GlobalCron.AddFunc(everyNminutes(1), func() {
		log.Debug("cron job run every 1 minute")
	})

	GlobalCron.Start()
}

func everyNminutes(num int) string {
	return fmt.Sprintf("@every %dm", num)
}
