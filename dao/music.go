package dao

type Music struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	Album      string `json:"album"`
	Type       int    `json:"type"`
	From       string `json:"from"`
	BitRate    int    `json:"bitRate"`
	SampleRate int    `json:"sampleRate"`
	Idx        int    `json:"idx"`
}
