package crontab

import (
	"agent/dao"
	"fmt"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func RunCron() {
	spec := "0/1 * * * *"
	c := cron.New()
	_, err := c.AddFunc(spec, readLog)
	if err != nil {
		fmt.Println("[cron:error]:" + err.Error())
	}
	_, err = c.AddFunc(spec, ngxStatus)
	if err != nil {
		return
	}
	c.Start()
}

func readLog() {
	dao.InsertSectionLog()
}

func ngxStatus() {
	url := "http://localhost:10110/ngx_status"
	if resp, err := http.Get(url); err == nil {
		if resp.StatusCode == http.StatusOK {
			if all, err := ioutil.ReadAll(resp.Body); err == nil {
				lines := strings.Split(string(all), "\n")
				trim := strings.Trim(lines[0], "Active connections: ")
				connections, err := strconv.Atoi(strings.Trim(trim, " "))
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				split := strings.Split(lines[2], " ")
				requests, err := strconv.Atoi(split[2])
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				dao.InsertNgxStatus(connections, requests)
			}
		}
	}
}
