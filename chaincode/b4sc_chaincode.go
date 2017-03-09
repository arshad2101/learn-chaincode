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
	WayBillId        string
	Consigner        string
	ConsignerAddress string
	Consignee        string
	ConsigneeAddress string
	ConsigneeRegNo   string
	LastModifiedDate string
	Quantity         int
	Assets           []string
	Cartons          []string
	Pallets          []string
}

type CreateWayBillRequest struct {
	WayBillId        string
	Consigner        string
	ConsignerAddress string
	Consignee        string
	ConsigneeAddress string
	ConsigneeRegNo   string
	LastModifiedDate string
	Quantity         int
	Assets           []string
	Cartons          []string
	Pallets          []string
}

type CreateWayBillResponse struct {
	Err     string `json:"err"`
	ErrMsg  string `json:"errMsg"`
	Message string `json:"message"`
}

type MWayBill struct {
	MWayBillId       string
	CreatedDate      string
	LastModifiedDate string
	Status           string
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
	CreatedBy        string
	PendingWith      string
	Pallets          []string
	Cartons          []string
	Assets           []string
}

type CreateMWayBillRequest struct {
	MWayBillId       string
	CreatedDate      string
	LastModifiedDate string
	Status           string
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
	CreatedBy        string
	PendingWith      string
	Pallets          []string
	Cartons          []string
	Assets           []string
}
type CreateMWayBillResponse struct {
	Err     string `json:"err"`
	ErrMsg  string `json:"errMsg"`
	Message string `json:"message"`
}

type MMWayBill struct {
	MMWayBillId          string
	CreatedDate          string
	LastModifiedDate     string
	Status               string
	ConatinerNo          string
	ConsignerAddress     string
	Consignee            string
	ConsigneeAddress     string
	ConsigneeRegNo       string
	PersonConsigning     string
	VehicleId            string
	ExportWarehouseNotes string
	CreatedBy            string
	PendingWith          string
	Conatiner            string
	MWayBills            []string
	Pallets              []string
	Cartons              []string
	Assets               []string
}

type CreateMMWayBillRequest struct {
	MMWayBillId          string
	CreatedDate          string
	LastModifiedDate     string
	Status               string
	ConatinerNo          string
	ConsignerAddress     string
	Consignee            string
	ConsigneeAddress     string
	ConsigneeRegNo       string
	PersonConsigning     string
	VehicleId            string
	ExportWarehouseNotes string
	CreatedBy            string
	PendingWith          string
	Conatiner            string
	MWayBills            []string
	Pallets              []string
	Cartons              []string
	Assets               []string
}

type CreateMMWayBillResponse struct {
	Err     string `json:"err"`
	ErrMsg  string `json:"errMsg"`
	Message string `json:"message"`
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

func CreateWayBill(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering Master Master WayBill")

	wayBillRequest := parseWayBillRequest(args[0])

	return processWayBill(stub, wayBillRequest)

}

func CreateMWayBill(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering Master Master WayBill")

	mWayBillRequest := parseMWayBillRequest(args[0])

	return processMWayBill(stub, mWayBillRequest)

}

func CreateMMWayBill(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering Master Master WayBill")

	mmWayBillRequest := parseMMWayBillRequest(args[0])

	return processMMWayBill(stub, mmWayBillRequest)

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

func processWayBill(stub shim.ChaincodeStubInterface, createWayBillRequest CreateWayBillRequest) ([]byte, error) {
	wayBill := WayBill{}
	//	shipmentIndex := ShipmentIndex{}

	wayBill.WayBillId = createWayBillRequest.WayBillId
	wayBill.Consigner = createWayBillRequest.Consigner
	wayBill.ConsignerAddress = createWayBillRequest.ConsignerAddress
	wayBill.Consignee = createWayBillRequest.Consignee
	wayBill.ConsigneeAddress = createWayBillRequest.ConsigneeAddress
	wayBill.LastModifiedDate = createWayBillRequest.LastModifiedDate
	wayBill.Quantity = createWayBillRequest.Quantity
	wayBill.Assets = createWayBillRequest.Assets
	wayBill.Cartons = createWayBillRequest.Cartons
	wayBill.Pallets = createWayBillRequest.Pallets

	dataToStore, _ := json.Marshal(wayBill)

	err := stub.PutState(wayBill.WayBillId, []byte(dataToStore))
	if err != nil {
		fmt.Println("Could not save WayBill to ledger", err)
		return nil, err
	}

	resp := CreateWayBillResponse{}
	resp.Err = "000"
	resp.Message = wayBill.WayBillId

	respString, _ := json.Marshal(resp)

	fmt.Println("Successfully saved way bill")
	return []byte(respString), nil

}

func processMWayBill(stub shim.ChaincodeStubInterface, createMWayBillRequest CreateMWayBillRequest) ([]byte, error) {
	mWayBill := MWayBill{}
	//	shipmentIndex := ShipmentIndex{}

	mWayBill.MWayBillId = createMWayBillRequest.MWayBillId
	mWayBill.CreatedDate = createMWayBillRequest.CreatedDate
	mWayBill.LastModifiedDate = createMWayBillRequest.LastModifiedDate
	mWayBill.Status = createMWayBillRequest.Status
	mWayBill.ConsignerAddress = createMWayBillRequest.ConsignerAddress
	mWayBill.ConsigneeRegNo = createMWayBillRequest.ConsigneeRegNo
	mWayBill.ModelNo = createMWayBillRequest.ModelNo
	mWayBill.VehicleNumber = createMWayBillRequest.VehicleNumber
	mWayBill.VehicleType = createMWayBillRequest.VehicleType
	mWayBill.PickUpTime = createMWayBillRequest.PickUpTime
	mWayBill.ValuesOfGood = createMWayBillRequest.ValuesOfGood
	mWayBill.ConsignerNotes = createMWayBillRequest.ConsignerNotes
	mWayBill.CreatedBy = createMWayBillRequest.CreatedBy
	mWayBill.PendingWith = createMWayBillRequest.PendingWith
	mWayBill.Pallets = createMWayBillRequest.Pallets
	mWayBill.Cartons = createMWayBillRequest.Cartons
	mWayBill.Assets = createMWayBillRequest.Assets

	dataToStore, _ := json.Marshal(mWayBill)

	err := stub.PutState(mWayBill.MWayBillId, []byte(dataToStore))
	if err != nil {
		fmt.Println("Could not save master Way Bill to ledger", err)
		return nil, err
	}

	resp := CreateMWayBillResponse{}
	resp.Err = "000"
	resp.Message = mWayBill.MWayBillId

	respString, _ := json.Marshal(resp)

	fmt.Println("Successfully saved master way bill")
	return []byte(respString), nil

}

func processMMWayBill(stub shim.ChaincodeStubInterface, createMMWayBillRequest CreateMMWayBillRequest) ([]byte, error) {
	mmWayBill := MMWayBill{}

	mmWayBill.MMWayBillId = createMMWayBillRequest.MMWayBillId
	mmWayBill.CreatedDate = createMMWayBillRequest.CreatedDate
	mmWayBill.LastModifiedDate = createMMWayBillRequest.LastModifiedDate
	mmWayBill.Status = createMMWayBillRequest.Status
	mmWayBill.ConatinerNo = createMMWayBillRequest.ConatinerNo
	mmWayBill.ConsignerAddress = createMMWayBillRequest.ConsignerAddress
	mmWayBill.Consignee = createMMWayBillRequest.Consignee
	mmWayBill.ConsigneeAddress = createMMWayBillRequest.ConsigneeAddress
	mmWayBill.ConsigneeRegNo = createMMWayBillRequest.ConsigneeRegNo
	mmWayBill.PersonConsigning = createMMWayBillRequest.PersonConsigning
	mmWayBill.VehicleId = createMMWayBillRequest.VehicleId
	mmWayBill.ExportWarehouseNotes = createMMWayBillRequest.ExportWarehouseNotes
	mmWayBill.CreatedBy = createMMWayBillRequest.CreatedBy
	mmWayBill.PendingWith = createMMWayBillRequest.PendingWith
	mmWayBill.Conatiner = createMMWayBillRequest.Conatiner
	mmWayBill.MWayBills = createMMWayBillRequest.MWayBills
	mmWayBill.Pallets = createMMWayBillRequest.Pallets
	mmWayBill.Cartons = createMMWayBillRequest.Cartons
	mmWayBill.Assets = createMMWayBillRequest.Assets

	dataToStore, _ := json.Marshal(mmWayBill)

	err := stub.PutState(mmWayBill.MMWayBillId, []byte(dataToStore))
	if err != nil {
		fmt.Println("Could not save Master Master Way Bill to ledger", err)
		return nil, err
	}

	resp := CreateMMWayBillResponse{}
	resp.Err = "000"
	resp.Message = mWayBill.MMWayBillId

	respString, _ := json.Marshal(resp)

	fmt.Println("Successfully saved Master Master way bill")
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

func parseWayBillRequest(jsondata string) CreateWayBillRequest {
	res := CreateWayBillRequest{}
	json.Unmarshal([]byte(jsondata), &res)
	fmt.Println(res)
	return res
}

func parseMWayBillRequest(jsondata string) CreateMWayBillRequest {
	res := CreateMWayBillRequest{}
	json.Unmarshal([]byte(jsondata), &res)
	fmt.Println(res)
	return res
}

func parseMMWayBillRequest(jsondata string) CreateMMWayBillRequest {
	res := CreateMMWayBillRequest{}
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
func ViewWayBill(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering ViewWayBill " + args[0])

	WayBillId := args[0]
	//request := parseViewShipmentRequest(args[0])

	waybilldata, dataerr := fetchWayBillData(stub, WayBillId)
	if dataerr == nil {

		dataToStore, _ := json.Marshal(waybilldata)
		return []byte(dataToStore), nil

	}

	return nil, dataerr

}

func ViewMWayBill(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering ViewMWayBill " + args[0])

	mWayBillId := args[0]
	//request := parseViewShipmentRequest(args[0])

	mWaybilldata, dataerr := fetchMWayBillData(stub, mWayBillId)
	if dataerr == nil {

		dataToStore, _ := json.Marshal(mWaybilldata)
		return []byte(dataToStore), nil

	}

	return nil, dataerr

}

func ViewMMWayBill(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering ViewMMWayBill " + args[0])

	mmWayBillId := args[0]

	mmWaybilldata, dataerr := fetchMWayBillData(stub, mmWayBillId)
	if dataerr == nil {

		dataToStore, _ := json.Marshal(mmWaybilldata)
		return []byte(dataToStore), nil

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

func fetchWayBillData(stub shim.ChaincodeStubInterface, WayBillId string) (WayBill, error) {
	var wayBill WayBill

	indexByte, err := stub.GetState(WayBillId)
	if err != nil {
		fmt.Println("Could not retrive WayBill ", err)
		return wayBill, err
	}

	if marshErr := json.Unmarshal(indexByte, &wayBill); marshErr != nil {
		fmt.Println("Could not retrieve WayBill from ledger", marshErr)
		return wayBill, marshErr
	}

	return wayBill, nil

}

func fetchMWayBillData(stub shim.ChaincodeStubInterface, mWayBillId string) (MWayBill, error) {
	var mWayBill MWayBill

	indexByte, err := stub.GetState(mWayBillId)
	if err != nil {
		fmt.Println("Could not retrive MWayBill ", err)
		return mWayBill, err
	}

	if marshErr := json.Unmarshal(indexByte, &mWayBill); marshErr != nil {
		fmt.Println("Could not retrieve master WayBill from ledger", marshErr)
		return mWayBill, marshErr
	}

	return mWayBill, nil

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
	} else if function == "CreateWayBill" {
		return CreateWayBill(stub, args)
	} else if function == "CreateMWayBill" {
		return CreateMWayBill(stub, args)
	} else if function == "CreateMWayBill" {
		return CreateMMWayBill(stub, args)
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
	} else if function == "ViewWayBill" {
		return ViewWayBill(stub, args)
	} else if function == "ViewMWayBill" {
		return ViewMWayBill(stub, args)
	} else if function == "ViewMMWayBill" {
		return ViewMMWayBill(stub, args)
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
