package model

// AccessLog
// Nginx log_format 参数
// log_format json '{"remoteUser":"$remote_user","request":"$request","upstreamResponseTime":"$upstream_response_time","requestLength":"$request_length","httpUserAgent":"$http_user_agent","requestTime":"$request_time","httpHost":"$http_host","httpReferer":"$http_referer","remoteAddr":"$remote_addr","timeLocal":"$time_local","status":"$status","upstreamStatus":"$upstream_status","bodyBytesDent":"$body_bytes_sent"}'
type AccessLog struct {
	RemoteUser           string `json:"remoteUser"`
	Request              string `json:"request"`
	UpstreamResponseTime string `json:"upstreamResponseTime"`
	RequestLength        string `json:"requestLength"`
	HttpUserAgent        string `json:"httpUserAgent"`
	RequestTime          string `json:"requestTime"`
	HttpHost             string `json:"httpHost"`
	HttpReferer          string `json:"httpReferer"`
	RemoteAddr           string `json:"remoteAddr"`
	TimeLocal            string `json:"timeLocal"`
	Status               string `json:"status"`
	BodyBytesDent        string `json:"bodyBytesDent"`
	UpstreamStatus       string `json:"upstreamStatus"`
}

// SectionLog
//  {
//		"code": {},
//		"upsCode": {},
//		"reqLen": 0,
//		"reqTime": 0,
//		"upsRespTime": 0
//	}
type SectionLog struct {
	Code        map[string]int `json:"code"`
	UpsCode     map[string]int `json:"upsCode"`
	ReqLen      int            `json:"reqLen"`
	ReqTime     float64        `json:"reqTime"`
	UpsRespTime float64        `json:"upsRespTime"`
}
