package main

import (
	"github.com/hyperledger-labs/hyper-fabric-test-GOlang/chaincode/assettypes"
	"github.com/hyperledger-labs/cc-tools/assets"
)

var assetTypeList = []assets.AssetType{
	assettypes.Person,
	assettypes.Book,
	assettypes.Library,
	assettypes.Secret,
}
