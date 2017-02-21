package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type SampleChaincode struct {
}

//custom data models
type WayBill struct {
	WayBillID        string `json:"id"`
	AssetType        string `json:"assetType"`
	LastModifiedDate string `json:"lastModifiedDate"`
	Quantity         int    `json:"quantity"`
}

type mWayBill struct {
	WayBillID        string   `json:"wayBillID"`
	CreatedDate      string   `json:"createdDate"`
	LastModifiedDate string   `json:"lastModifiedDate"`
	Status           string   `json:"status"`
	CreatedBy        string   `json:"createdBy"`
	PendingWith      string   `json:"pendingWith"`
	Palettes         []string `json:"palettes"`
}

func GetWayBill(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering GetWayBill")

	if len(args) < 1 {
		fmt.Println("Invalid number of arguments")
		return nil, errors.New("Missing way bill ID")
	}

	var wayBillId = args[0]

	bytes, err := stub.GetState(wayBillId)

	wayBillStruct := WayBill{}

	json.Unmarshal(bytes, &wayBillStruct)

	fmt.Println("WayBillID ", wayBillStruct.WayBillID)
	fmt.Println("AssetType ", wayBillStruct.AssetType)
	fmt.Println("LastModifiedDate ", wayBillStruct.LastModifiedDate)
	fmt.Println("Quantity ", wayBillStruct.Quantity)

	if err != nil {
		fmt.Println("Could not fetch loan way bill with id "+wayBillId+" from ledger", err)
		return nil, err
	}
	return bytes, nil
}

func GetMasterWayBill(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering GetMasterWayBill")

	if len(args) < 1 {
		fmt.Println("Invalid number of arguments")
		return nil, errors.New("Missing master way bill ID")
	}

	var mWayBillId = args[0]
	bytes, err := stub.GetState(mWayBillId)
	if err != nil {
		fmt.Println("Could not fetch loan way bill with id "+mWayBillId+" from ledger", err)
		return nil, err
	}
	return bytes, nil
}

func CreateWayBill(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering CreateWayBill")

	if len(args) < 2 {
		fmt.Println("Invalid number of args")
		return nil, errors.New("Expected atleast two arguments for way bill creation")
	}

	//var wayBillId = args[0]
	//var wayBillInput = args[1]

	var wayBillID = args[0]
	var assetType = args[1]
	var lastModifiedDate = args[2]
	var quantity, _ = strconv.Atoi(args[3])
	wayBillStruct := WayBill{
		WayBillID:        wayBillID,
		AssetType:        assetType,
		LastModifiedDate: lastModifiedDate,
		Quantity:         quantity,
	}
	bytes, _ := json.Marshal(wayBillStruct)

	err := stub.PutState(wayBillID, bytes)
	if err != nil {
		fmt.Println("Could not save way bill to ledger", err)
		return nil, err
	}

	fmt.Println("Successfully saved way bill")
	return nil, nil

}

func CreateMasterWayBill(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering CreateWayBill")

	if len(args) < 2 {
		fmt.Println("Invalid number of args")
		return nil, errors.New("Expected atleast two arguments for master way bill creation")
	}

	var mWayBillId = args[0]
	var mwayBillInput = args[1]

	err := stub.PutState(mWayBillId, []byte(mwayBillInput))
	if err != nil {
		fmt.Println("Could not save master way bill to ledger", err)
		return nil, err
	}

	fmt.Println("Successfully saved master way bill")
	return nil, nil

}

func (t *SampleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Inside INIT for test chaincode")
	return nil, nil
}

func (t *SampleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "GetWayBill" {
		return GetWayBill(stub, args)
	} else if function == "GetMasterWayBill" {
		return GetMasterWayBill(stub, args)
	} else {
		return nil, errors.New("Invalid function name " + function)
	}

	return nil, nil
}

func GetCertAttribute(stub shim.ChaincodeStubInterface, attributeName string) (string, error) {
	fmt.Println("Entering GetCertAttribute")
	attr, err := stub.ReadCertAttribute(attributeName)
	if err != nil {
		return "", errors.New("Couldn't get attribute " + attributeName + ". Error: " + err.Error())
	}
	attrString := string(attr)
	return attrString, nil
}

func (t *SampleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "CreateWayBill" {
		return CreateWayBill(stub, args)
	} else if function == "CreateMasterWayBill" {
		return CreateMasterWayBill(stub, args)
	} else {
		return nil, errors.New("Invalid function name " + function)
	}
	return nil, nil
}

func main() {
	err := shim.Start(new(SampleChaincode))
	if err != nil {
		fmt.Println("Could not start SampleChaincode")
	} else {
		fmt.Println("SampleChaincode successfully started")
	}

	// var str = `{"id":"wb1","assetType":"chip","lastModifiedDate":"21/09/2016 2:30pm","quantity":10}`
	// stub := shim.NewMockStub("mock", new(SampleChaincode))

	// args := []string{"wb1", str}
	// _, err := stub.MockInvoke("123", "CreateWayBill", args)
	// fmt.Println(err)

	// bytes, err := stub.MockQuery("GetWayBill", []string{"wb1"})
	// var wb WayBill
	// err = json.Unmarshal(bytes, &wb)

	// fmt.Println(wb)

}
