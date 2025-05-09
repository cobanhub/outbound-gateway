package outbound

import "time"

type Outbound struct {
	timeout time.Duration
}

type OutboundOptions struct {
	Timeout time.Duration
}

type OutboundInterface interface {
	HandleJson(integrationName string, coreRequest map[string]interface{}) (map[string]interface{}, error)
	HandleXML(integrationName string, coreRequest map[string]interface{}) (map[string]interface{}, error)
}

func NewOutbound(opts OutboundOptions) *Outbound {
	return &Outbound{
		timeout: opts.Timeout,
	}
}
