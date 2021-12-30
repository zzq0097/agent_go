package crontab

import (
	"agent/consts"
	"agent/dao"
	"agent/model"
	"agent/service"
	"agent/util"
	"encoding/json"
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
		fmt.Println(err)
	}
	_, err = c.AddFunc(spec, ngxStatus)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Start()
}

func readLog() {
	offset, err := dao.GetLogOffset()
	if err != nil {
		fmt.Println(err)
		return
	}
	dao.SaveOrUpdateParameter(consts.AccessLogOffset, strconv.FormatInt(util.FileSize(consts.NgxRpmAccessLog), 10))
	sectionLog := model.SectionLog{}
	service.ReadOffset(consts.NgxRpmAccessLog, offset, &sectionLog)
	marshal, err := json.Marshal(sectionLog)
	if err != nil {
		fmt.Println(err)
		return
	}
	if dao.InsertSectionLog(string(marshal)) < 1 {
		fmt.Println("ngx_log insert 0")
	}
}

func ngxStatus() {
	url := "http://localhost:10110/ngx_status"
	client := http.Client{Timeout: 1000}
	if resp, err := client.Get(url); err != nil {
		fmt.Println(err)
	} else if resp.StatusCode == http.StatusOK {
		if all, err := ioutil.ReadAll(resp.Body); err != nil {
			fmt.Println(err)
		} else {
			lines := strings.Split(string(all), "\n")
			trim := strings.Trim(lines[0], "Active connections: ")
			connections, err := strconv.Atoi(strings.Trim(trim, " "))
			if err != nil {
				fmt.Println(err)
				return
			}
			split := strings.Split(lines[2], " ")
			requests, err := strconv.Atoi(split[2])
			if err != nil {
				fmt.Println(err)
				return
			}
			if dao.InsertNgxStatus(connections, requests) < 1 {
				fmt.Println("ngx_status insert 0")
			}
		}
	}
}
