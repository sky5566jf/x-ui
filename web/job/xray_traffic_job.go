package job

import (
	"x-ui/logger"
	"x-ui/web/service"
)

type XrayTrafficJob struct {
	xrayService    service.XrayService
	inboundService service.InboundService
}

func NewXrayTrafficJob() *XrayTrafficJob {
	return new(XrayTrafficJob)
}

func (j *XrayTrafficJob) Run() {
	if !j.xrayService.IsXrayRunning() {
		return
	}

	// get Client Traffic

	clientTraffics, err := j.xrayService.GetXrayClientTraffic()
	if err != nil {
		logger.Warning("get xray client traffic failed:", err)
		return
	}
	err = j.inboundService.AddClientTraffic(clientTraffics)
	if err != nil {
		logger.Warning("add client traffic failed:", err)
	}
	
	traffics, err := j.xrayService.GetXrayTraffic()
	if err != nil {
		logger.Warning("get xray traffic failed:", err)
		return
	}
	err = j.inboundService.AddTraffic(traffics)
	if err != nil {
		logger.Warning("add traffic failed:", err)
	}
	


}
