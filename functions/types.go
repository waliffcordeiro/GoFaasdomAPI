package function

type Message struct {
	Success bool    `json:"success,omitempty"`
	Payload Payload `json:"payload,omitempty"`
	Metrics Metrics `json:"metrics,omitempty"`
}

type Payload struct {
	Test      string `json:"test,omitempty"`
	N         int    `json:"n,omitempty"`
	Size      int    `json:"size,omitempty"`
	TimeWrite string `json:"timewrite,omitempty"`
	TimeRead  string `json:"timeread,omitempty"`
	Time      string `json:"time,omitempty"`
	Result    []int  `json:"result,omitempty"`
}

type Metrics struct {
	MachineId  string `json:"machineid,omitempty"`
	InstanceId string `json:"instanceid,omitempty"`
	Cpu        string `json:"cpu,omitempty"`
	Mem        string `json:"mem,omitempty"`
	Uptime     string `json:"uptime,omitempty"`
}
