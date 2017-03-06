package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type B4SCChaincode struct {
}

//custom data models

type Pallet struct {
	PalletId  string
	Modeltype string
	CartonId  []string
	WayBill   []string
}

type Carton struct {
	CartonId string
	AssetId  []string
}

type Asset struct {
	AssetId   string
	Modeltype string
	color     string
}

type WayBill struct {
	WayBillId string `json:"Id"`
	Consigner
	ConsignerAddress string
	Consignee        string
	ConsigneeAddress string
	ConsigneeRegNo   string
	LastModifiedDate string   `json:"lastModifiedDate"`
	Quantity         int      `json:"quantity"`
	Assets           []Assets `json:"Assets"`
	Cartons          []Cartons
	Pallets          []Pallets
}

type MWayBill struct {
	MWayBillId       string `json:"wayBillId"`
	CreatedDate      string `json:"createdDate"`
	LastModifiedDate string `json:"lastModifiedDate"`
	Status           string `json:"status"`
	ConsignerAddress string
	Consignee        string
	ConsigneeAddress string
	ConsigneeRegNo   string
	ModelNo          string
	VehicleNumber    string
	VehicleType      string
	PickUpTime       string
	ValuesOfGood     string
	ConsignerNotes   string
	CreatedBy        string     `json:"createdBy"`
	PendingWith      string     `json:"pendingWith"`
	Palettes         []Palettes `json:"palettes"`
	Cartons          []Cartons
	Assets           []Assets `json:"Assets"`
}

type MMWayBill struct {
	MMWayBillId          string `json:"wayBillId"`
	CreatedDate          string `json:"createdDate"`
	LastModifiedDate     string `json:"lastModifiedDate"`
	Status               string `json:"status"`
	ConatinerNo          string
	ConsignerAddress     string
	Consignee            string
	ConsigneeAddress     string
	ConsigneeRegNo       string
	PersonConsigning     string
	VehicleId            string
	ExportWarehouseNotes string
	CreatedBy            string `json:"createdBy"`
	PendingWith          string `json:"pendingWith"`
	Conatiner            string
	MWayBills            []MWayBill `json:"mWayBill"`
	Palettes             []Palettes
	Cartons              []Cartons
	Assets               []Assets `json:"Assets"`
}
type WayBillHistory struct {
	name      string
	address   string
	timestamp string
	lat       string
	log       string
}

type Note struct {
	Org   string
	Notes string
}

type Shipment struct {
	ShipmentNumber        string
	WayBillNo             string
	WayBillType           string
	PersonConsigningGoods string
	Consigner             string
	ConsignerAddress      string
	Consignee             string
	ConsigneeAddress      string
	ConsigneeRegNo        string
	Quantity              string
	Pallets               []string
	Cartons               []string
	Status                string
	ModelNo               string
	VehicleNumber         string
	VehicleType           string
	PickUpTime            string
	ValueOfGoods          string
	AllNotes              []Note
	ContainerId           string
	WayBills              []string
	WayBillHistorys       []WayBillHistory
	Carrier               string
	Acl                   []string
	CreatedBy             string
	Custodian             string
	CreatedTimeStamp      string
	UpdatedTimeStamp      string
}

type ShipmentIndex struct {
	ShipmentNumber string
	Status         string
	Acl            []string
}

type AllShipment struct {
	ShipmentIndexArr []ShipmentIndex
}

/************** Create Shipment Starts ************************/
/**
	Expected Input is
	{
		"shipmentNumber"" : "123456",
		"personConsigningGoods" : "KarthikS",
		"consigner" : "HCL",
		"consignerAddress" : "Chennai",
		"consignee" : "HCL-AM",
		"consigneeAddress" : "Dallas",
		"consigneeRegNo" : "12122222222",
		"ModelNo" : "IA1a1222",
		"quantity" : "50",
		"pallets" : ["11111111","22222222","333333"],
		"status" : "intra",
		"notes" : "ha haha ha ha",
		"CreatedBy" : "KarthikSukumaram",
		"custodian" : "HCL",
		"createdTimeStamp" : "2017-03-02"
	}
**/

type CreateShipmentRequest struct {
	ShipmentNumber        string   `json:"shipmentNumber"`
	PersonConsigningGoods string   `json:"personConsigningGoods"`
	Consigner             string   `json:"consigner"`
	ConsignerAddress      string   `json:"consignerAddress"`
	Consignee             string   `json:"consignee"`
	ConsigneeAddress      string   `json:"consigneeAddress"`
	ConsigneeRegNo        string   `json:"consigneeRegNo"`
	ModelNo               string   `json:"modelNo"`
	Quantity              string   `json:"quantity"`
	Pallets               []string `json:"pallets"`
	Carrier               string   `json:"status"`
	Notes                 string   `json:"notes"`
	CreatedBy             string   `json:"createdBy"`
	Custodian             string   `json:"custodian"`
	CreatedTimeStamp      string   `json:"createdTimeStamp"`
	CallingEntityName     string   `json:"callingEntityName"`
}

type CreateShipmentResponse struct {
	Err     string `json:"err"`
	ErrMsg  string `json:"errMsg"`
	Message string `json:"message"`
}

func CreateMMWayBill(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering MMWayBill")

	shipmentRequest := parseCreateMMWaybillRequest(args[0])

	return processShipment(stub, shipmentRequest)

}
func parseCreateMMWaybillRequest(jsondata string) CreateShipmentRequest {
	res := CreateShipmentRequest{}
	json.Unmarshal([]byte(jsondata), &res)
	fmt.Println(res)
	return res
}


func processShipment(stub shim.ChaincodeStubInterface, shipmentRequest CreateShipmentRequest) ([]byte, error) {
	shipment := Shipment{}
	shipmentIndex := ShipmentIndex{}

	shipment.ShipmentNumber = shipmentRequest.ShipmentNumber
	shipment.PersonConsigningGoods = shipmentRequest.PersonConsigningGoods
	shipment.Consigner = shipmentRequest.Consigner
	shipment.ConsignerAddress = shipmentRequest.ConsignerAddress
	shipment.Consignee = shipmentRequest.Consignee
	shipment.ConsigneeAddress = shipmentRequest.ConsigneeAddress
	shipment.ConsigneeRegNo = shipmentRequest.ConsigneeRegNo
	shipment.ModelNo = shipmentRequest.ModelNo
	shipment.Quantity = shipmentRequest.Quantity
	shipment.Pallets = shipmentRequest.Pallets
	shipment.Carrier = shipmentRequest.Carrier
	shipment.CreatedBy = shipmentRequest.CreatedBy
	shipment.Custodian = shipmentRequest.Custodian
	shipment.CreatedTimeStamp = shipmentRequest.CreatedTimeStamp
	shipment.Status = "Created"

	manufacturerNotes := Note{}
	manufacturerNotes.Org = shipmentRequest.CallingEntityName
	manufacturerNotes.Notes = shipmentRequest.Notes
	var allNotes []Note
	allNotes = append(allNotes, manufacturerNotes)
	shipment.AllNotes = allNotes

	var acl []string
	acl = append(acl, shipmentRequest.CallingEntityName)
	shipment.Acl = acl

	shipmentIndex.ShipmentNumber = shipmentRequest.ShipmentNumber
	shipmentIndex.Status = shipment.Status
	shipmentIndex.Acl = acl

	dataToStore, _ := json.Marshal(shipment)

	err := stub.PutState(shipment.ShipmentNumber, []byte(dataToStore))
	if err != nil {
		fmt.Println("Could not save Shipment to ledger", err)
		return nil, err
	}

	

func CreateShipment(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering CreateShipment")

	shipmentRequest := parseCreateShipmentRequest(args[0])

	return processShipment(stub, shipmentRequest)

}

func processShipment(stub shim.ChaincodeStubInterface, shipmentRequest CreateShipmentRequest) ([]byte, error) {
	shipment := Shipment{}
	shipmentIndex := ShipmentIndex{}

	shipment.ShipmentNumber = shipmentRequest.ShipmentNumber
	shipment.PersonConsigningGoods = shipmentRequest.PersonConsigningGoods
	shipment.Consigner = shipmentRequest.Consigner
	shipment.ConsignerAddress = shipmentRequest.ConsignerAddress
	shipment.Consignee = shipmentRequest.Consignee
	shipment.ConsigneeAddress = shipmentRequest.ConsigneeAddress
	shipment.ConsigneeRegNo = shipmentRequest.ConsigneeRegNo
	shipment.ModelNo = shipmentRequest.ModelNo
	shipment.Quantity = shipmentRequest.Quantity
	shipment.Pallets = shipmentRequest.Pallets
	shipment.Carrier = shipmentRequest.Carrier
	shipment.CreatedBy = shipmentRequest.CreatedBy
	shipment.Custodian = shipmentRequest.Custodian
	shipment.CreatedTimeStamp = shipmentRequest.CreatedTimeStamp
	shipment.Status = "Created"

	manufacturerNotes := Note{}
	manufacturerNotes.Org = shipmentRequest.CallingEntityName
	manufacturerNotes.Notes = shipmentRequest.Notes
	var allNotes []Note
	allNotes = append(allNotes, manufacturerNotes)
	shipment.AllNotes = allNotes

	var acl []string
	acl = append(acl, shipmentRequest.CallingEntityName)
	shipment.Acl = acl

	shipmentIndex.ShipmentNumber = shipmentRequest.ShipmentNumber
	shipmentIndex.Status = shipment.Status
	shipmentIndex.Acl = acl

	dataToStore, _ := json.Marshal(shipment)

	err := stub.PutState(shipment.ShipmentNumber, []byte(dataToStore))
	if err != nil {
		fmt.Println("Could not save Shipment to ledger", err)
		return nil, err
	}

	addShipmentIndex(stub, shipmentIndex)

	resp := CreateShipmentResponse{}
	resp.Err = "000"
	resp.Message = shipment.ShipmentNumber

	respString, _ := json.Marshal(resp)

	fmt.Println("Successfully saved way bill")
	return []byte(respString), nil

}

func addShipmentIndex(stub shim.ChaincodeStubInterface, shipmentIndex ShipmentIndex) error {
	indexByte, err := stub.GetState("SHIPMENT_INDEX")
	if err != nil {
		fmt.Println("Could not retrive Shipment Index", err)
		return err
	}
	allShipmentIndex := AllShipment{}

	if marshErr := json.Unmarshal(indexByte, &allShipmentIndex); marshErr != nil {
		fmt.Println("Could not save Shipment to ledger", marshErr)
		return marshErr
	}

	allShipmentIndex.ShipmentIndexArr = append(allShipmentIndex.ShipmentIndexArr, shipmentIndex)
	dataToStore, _ := json.Marshal(allShipmentIndex)

	addErr := stub.PutState("SHIPMENT_INDEX", []byte(dataToStore))
	if addErr != nil {
		fmt.Println("Could not save Shipment to ledger", addErr)
		return addErr
	}

	return nil
}

func parseCreateShipmentRequest(jsondata string) CreateShipmentRequest {
	res := CreateShipmentRequest{}
	json.Unmarshal([]byte(jsondata), &res)
	fmt.Println(res)
	return res
}

/************** Create Shipment Ends ************************/

/************** View Shipment Starts ************************/

type ViewShipmentRequest struct {
	CallingEntityName string `json:"callingEntityName"`
	ShipmentNumber    string `json:"shipmentNumber"`
}

func ViewShipment(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering ViewShipment " + args[0])

	request := parseViewShipmentRequest(args[0])

	shipmentData, dataerr := fetchShipmentData(stub, request.ShipmentNumber)
	if dataerr == nil {
		if hasPermission(shipmentData.Acl, request.CallingEntityName) {
			dataToStore, _ := json.Marshal(shipmentData)
			return []byte(dataToStore), nil
		} else {
			return []byte("{ \"errMsg\": \"No data found\" }"), nil
		}
	}

	return nil, dataerr

}

func parseViewShipmentRequest(jsondata string) ViewShipmentRequest {
	res := ViewShipmentRequest{}
	json.Unmarshal([]byte(jsondata), &res)
	fmt.Println(res)
	return res
}

/************** View Shipment Ends ************************/

/************** Inbox Service Starts ************************/

/**
	Expected Input is
	{
		"callingEntityName" : "INTEL",
		"status" : "Created"
	}
**/

type InboxRequest struct {
	CallingEntityName string `json:"callingEntityName"`
	Status            string `json:"status"`
}

type InboxResponse struct {
	Data []Shipment `json:"data"`
}

func Inbox(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering Inbox " + args[0])

	request := parseInboxRequest(args[0])

	return fetchShipmentIndex(stub, request.CallingEntityName, request.Status)

}

func parseInboxRequest(jsondata string) InboxRequest {
	res := InboxRequest{}
	json.Unmarshal([]byte(jsondata), &res)
	fmt.Println(res)
	return res
}

func hasPermission(acl []string, currUser string) bool {
	lenOfArray := len(acl)

	for i := 0; i < lenOfArray; i++ {
		if acl[i] == currUser {
			return true
		}
	}

	return false
}

func fetchShipmentData(stub shim.ChaincodeStubInterface, shipmentNumber string) (Shipment, error) {
	var shipmentData Shipment

	indexByte, err := stub.GetState(shipmentNumber)
	if err != nil {
		fmt.Println("Could not retrive Shipment Index", err)
		return shipmentData, err
	}

	if marshErr := json.Unmarshal(indexByte, &shipmentData); marshErr != nil {
		fmt.Println("Could not save Shipment to ledger", marshErr)
		return shipmentData, marshErr
	}

	return shipmentData, nil

}

func fetchShipmentIndex(stub shim.ChaincodeStubInterface, callingEntityName string, status string) ([]byte, error) {
	allShipmentIndex := AllShipment{}
	var shipmentIndexArr []ShipmentIndex
	var tmpShipmentIndex ShipmentIndex
	var shipmentDataArr []Shipment
	resp := InboxResponse{}

	indexByte, err := stub.GetState("SHIPMENT_INDEX")
	if err != nil {
		fmt.Println("Could not retrive Shipment Index", err)
		return nil, err
	}

	if marshErr := json.Unmarshal(indexByte, &allShipmentIndex); marshErr != nil {
		fmt.Println("Could not save Shipment to ledger", marshErr)
		return nil, marshErr
	}

	shipmentIndexArr = allShipmentIndex.ShipmentIndexArr

	lenOfArray := len(shipmentIndexArr)

	for i := 0; i < lenOfArray; i++ {
		tmpShipmentIndex = shipmentIndexArr[i]
		if tmpShipmentIndex.Status == status {
			if hasPermission(tmpShipmentIndex.Acl, callingEntityName) {
				shipmentData, dataerr := fetchShipmentData(stub, tmpShipmentIndex.ShipmentNumber)
				if dataerr == nil {
					shipmentDataArr = append(shipmentDataArr, shipmentData)
				}
			}
		}
	}

	resp.Data = shipmentDataArr
	dataToStore, _ := json.Marshal(resp)

	return []byte(dataToStore), nil
}

/************** Inbox Service Ends ************************/

/************** View Data for Key Starts ************************/

func ViewDataForKey(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering ViewDataForKey " + args[0])

	return stub.GetState(args[0])

}

/************** View Data for Key Ends ************************/

/************** DumpData Start ************************/

func DumpData(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering DumpData " + args[0] + "  " + args[1])

	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		fmt.Println("Could not save the Data", err)
		return nil, err
	}

	return nil, nil
}

/************** DumpData Ends ************************/

// Init resets all the things
func (t *B4SCChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Inside INIT for test chaincode")

	allShipment := AllShipment{}
	var tmpShipmentIndex []ShipmentIndex
	allShipment.ShipmentIndexArr = tmpShipmentIndex

	dataToStore, _ := json.Marshal(allShipment)

	err := stub.PutState("SHIPMENT_INDEX", []byte(dataToStore))
	if err != nil {
		fmt.Println("Could not save Shipment to ledger", err)
		return nil, err
	}

	return nil, nil
}

func (t *B4SCChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	/*if function == "Init" {
		return Init(stub, function, args)
	}else*/
	if function == "CreateShipment" {
		return CreateShipment(stub, args)
	} else {
		return nil, errors.New("Invalid function name " + function)
	}
	//return nil, nil
}

func (t *B4SCChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "ViewShipment" {
		return ViewShipment(stub, args)
	} else if function == "ViewDataForKey" {
		return ViewDataForKey(stub, args)
	} else if function == "Inbox" {
		return Inbox(stub, args)
	} else if function == "DumpData" {
		return DumpData(stub, args)
	} else {
		return nil, errors.New("Invalid function name " + function)
	}
}

func main() {
	err := shim.Start(new(B4SCChaincode))
	if err != nil {
		fmt.Println("Could not start B4SCChaincode")
	} else {
		fmt.Println("B4SCChaincode successfully started")
	}
}
