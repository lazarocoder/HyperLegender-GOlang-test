package main

import (
	"github.com/hyperledger-labs/hyper-fabric-test-GOlang/chaincode/eventtypes"
	"github.com/hyperledger-labs/cc-tools/events"
)

var eventTypeList = []events.Event{
	eventtypes.CreateLibraryLog,
}
