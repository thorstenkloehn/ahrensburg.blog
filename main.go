package main

import (
	"github.com/robfig/cron/v3"
	"github.com/thorstenkloehn/ahrensburg.digital/module/wordpress_pages"
)

func main() {

	c := cron.New()
	c.AddFunc("@hourly", func() { wordpress_pages.Start() })
	c.Start()
	wordpress_pages.Start()

}
