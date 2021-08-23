package model

type Log struct {
	remoteUser           string
	request              string
	upstreamResponseTime string
	requestLength        string
	httpUserAgent        string
	requestTime          string
	httpHost             string
	httpReferer          string
	remoteAddr           string
	timeLocal            string
	status               string
	bodyBytesDent        string
	upstreamStatus       string
}
