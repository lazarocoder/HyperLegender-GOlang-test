package eventtypes

import "github.com/hyperledger-labs/cc-tools/events"

var CreateLibraryLog = events.Event{
	Tag:         "createLibraryLog",
	Label:       "Create Library Log",
	Description: "Log of a library creation",
	Type:        events.EventLog,                 
	BaseLog:     "New library created",           
	Receivers:   []string{"$org2MSP", "$orgMSP"}, 
}
