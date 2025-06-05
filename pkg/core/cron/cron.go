package cron

import (
	cron "github.com/robfig/cron/v3"
)

var c *cron.Cron

func Init() {
	c = cron.New()
	c.Start()
}

func Add(spec string, job func()) (cron.EntryID, error) {
	return c.AddFunc(spec, job)
}
