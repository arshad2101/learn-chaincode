package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const NODATA_ERROR_CODE string = "400"
const NODATA_ERROR_MSG string = "No data found"

const INVALID_INPUT_ERROR_CODE string = "401"
const INVALID_INPUT_ERROR_MSG string = "Invalid Input"

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

type B4SCChaincode struct {
}

//custom data models

type Pallet struct {
	PalletId    string
	Modeltype   string
	CartonId    []string
	ShipmentIds []string
}

type Carton struct {
	CartonId    string
	PalletId    string
	AssetId     []string
	ShipmentIds []string
}

type Asset struct {
	AssetId     string
	Modeltype   string
	Color       string
	CartonId    string
	PalletId    string
	ShipmentIds []string
}

/************** Arshad This new struct for AssetDetails , CartonDetails , PalletDetails  is added by Arshad as to incorporate new LLD published orginal structure
are not touched as of now to avoid break of any functionality devloped by Kartik 20/3/2017***************/

type AssetDetails struct {
	assetSerialNo      string
	assetModel         string
	assetType          string
	assetMake          string
	assetCOO           string
	assetMaufacture    string
	assetStatus        string
	createdBy          string
	createdDate        string
	modifiedBy         string
	modifiedDate       string
	palletSerialNumber string
	cartonSerialNumber string
	mshipmentNumber    string
	dcShipmentNumber   string
	mwayBillNumber     string
	dcWayBillNumber    string
	ewWayBillNumber    string
	mShipmentDate      string
	dcShipmentDate     string
	mWayBillDate       string
	dcWayBillDate      string
	ewWayBillDate      string
}
type CartonDetails struct {
	cartonSerialNumber string
	cartonModel        string
	cartonStatus       string
	cartonCreationDate string
	palletSerialNumber string
	assetsSerialNumber []string
	bnmshipmentNumber  string
	dcShipmentNumber   string
	mwayBillNumber     string
	dcWayBillNumber    string
	ewWayBillNumber    string
	dimensions         string
	weight             string
	mShipmentDate      string
	dcShipmentDate     string
	mWayBillDate       string
	dcWayBillDate      string
	ewWayBillDate      string
}
type PalletDetails struct {
	palletSerialNumber string
	palletModel        string
	palletStatus       string
	cartonSerialNumber []string
	palletCreationDate string
	assetsSerialNumber []string
	mshipmentNumber    string
	dcShipmentNumber   string
	mwayBillNumber     string
	dcWayBillNumber    string
	ewWayBillNumber    string
	dimensions         string
	weight             string
	mShipmentDate      string
	dcShipmentDate     string
	mWayBillDate       string
	dcWayBillDate      string
	ewWayBillDate      string
}
type WayBill struct {
	WayBillNumber         string
	ShipmentNumber        string
	CountryFrom           string
	CountryTo             string
	Consigner             string
	Consignee             string
	Custodian             string
	CustodianHistory      []string
	PersonConsigningGoods string
	Comments              string
	TpComments            string
	VehicleNumber         string
	VehicleType           string
	PickupDate            string
	PalletsSerialNumber   []string
	AddressOfConsigner    string
	AddressOfConsignee    string
	ConsignerRegNumber    string
	Carrier               string
	VesselType            string
	VesselNumber          string
	ContainerNumber       string
	ServiceType           string
	ShipmentModel         string
	PalletsQuantity       string
	CartonsQuantity       string
	AssetsQuantity        string
	ShipmentValue         string
	EntityName            string
	ShipmentCreationDate  string
	EWWayBillNumber       string
	SupportiveDocuments   []string
	ShipmentCreatedBy     string
	ShipmentModifiedDate  string
	ShipmentModifiedBy    string
	WayBillCreationDate   string
	WayBillCreatedBy      string
	WayBillModifiedDate   string
	WayBillModifiedBy     string
}
type CreateWayBillRequest struct {
	WayBillNumber         string
	ShipmentNumber        string
	CountryFrom           string
	CountryTo             string
	Consigner             string
	Consignee             string
	Custodian             string
	CustodianHistory      []string
	PersonConsigningGoods string
	Comments              string
	TpComments            string
	VehicleNumber         string
	VehicleType           string
	PickupDate            string
	PalletsSerialNumber   []string
	AddressOfConsigner    string
	AddressOfConsignee    string
	ConsignerRegNumber    string
	Carrier               string
	VesselType            string
	VesselNumber          string
	ContainerNumber       string
	ServiceType           string
	ShipmentModel         string
	PalletsQuantity       string
	CartonsQuantity       string
	AssetsQuantity        string
	ShipmentValue         string
	EntityName            string
	ShipmentCreationDate  string
	EWWayBillNumber       string
	SupportiveDocuments   []string
	ShipmentCreatedBy     string
	ShipmentModifiedDate  string
	ShipmentModifiedBy    string
	WayBillCreationDate   string
	WayBillCreatedBy      string
	WayBillModifiedDate   string
	WayBillModifiedBy     string
}
type CreateWayBillResponse struct {
	Err     string `json:"err"`
	ErrMsg  string `json:"errMsg"`
	Message string `json:"message"`
}
type EWWayBill struct {
	EwWayBillNumber       string
	WayBillsNumber        []string
	ShipmentsNumber       []string
	CountryFrom           string
	CountryTo             string
	Consigner             string
	Consignee             string
	Custodian             string
	CustodianHistory      []string
	CustodianTime         string
	PersonConsigningGoods string
	Comments              string
	PalletsSerialNumber   []string
	AddressOfConsigner    string
	AddressOfConsignee    string
	ConsignerRegNumber    string
	VesselType            string
	VesselNumber          string
	ContainerNumber       string
	ServiceType           string
	SupportiveDocuments   []string
	EwWayBillCreationDate string
	EwWayBillCreatedBy    string
	EwWayBillModifiedDate string
	EwWayBillModifiedBy   string
}
type CreateEWWayBillRequest struct {
	EwWayBillNumber       string
	WayBillsNumber        []string
	ShipmentsNumber       []string
	CountryFrom           string
	CountryTo             string
	Consigner             string
	Consignee             string
	Custodian             string
	CustodianHistory      []string
	CustodianTime         string
	PersonConsigningGoods string
	Comments              string
	PalletsSerialNumber   []string
	AddressOfConsigner    string
	AddressOfConsignee    string
	ConsignerRegNumber    string
	VesselType            string
	VesselNumber          string
	ContainerNumber       string
	ServiceType           string
	SupportiveDocuments   []string
	EwWayBillCreationDate string
	EwWayBillCreatedBy    string
	EwWayBillModifiedDate string
	EwWayBillModifiedBy   string
}
type CreateEWWayBillResponse struct {
	Err     string `json:"err"`
	ErrMsg  string `json:"errMsg"`
	Message string `json:"message"`
}
type EntityWayBillMapping struct {
	WayBillsNumber []string
}
type CreateEntityWayBillMappingRequest struct {
	EntityName     string
	WayBillsNumber []string
}
type EntityWayBillMappingResponse struct {
	Err     string `json:"err"`
	ErrMsg  string `json:"errMsg"`
	Message string `json:"message"`
}
type EntityDetails struct {
	EntityName      string
	EntityType      string
	EntityAddress   string
	EntityRegNumber string
	EntityCountry   string
	Latitude        string
	Longitude       string
}

/************** Create WayBill Starts ************************/
func CreateWayBill(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering CreateWayBill", args[0])

	wayBillRequest := parseWayBillRequest(args[0])

	return processWayBill(stub, wayBillRequest)

}
func parseWayBillRequest(jsondata string) CreateWayBillRequest {
	res := CreateWayBillRequest{}
	json.Unmarshal([]byte(jsondata), &res)
	fmt.Println(res)
	return res
}
func processWayBill(stub shim.ChaincodeStubInterface, createWayBillRequest CreateWayBillRequest) ([]byte, error) {
	wayBill := WayBill{}
	//	shipmentIndex := ShipmentIndex{}
	wayBill.WayBillNumber = createWayBillRequest.WayBillNumber
	wayBill.ShipmentNumber = createWayBillRequest.ShipmentNumber
	wayBill.CountryFrom = createWayBillRequest.CountryFrom
	wayBill.CountryTo = createWayBillRequest.CountryTo
	wayBill.Consigner = createWayBillRequest.Consigner
	wayBill.Consignee = createWayBillRequest.Consignee
	wayBill.Custodian = createWayBillRequest.Custodian
	wayBill.CustodianHistory = createWayBillRequest.CustodianHistory
	wayBill.PersonConsigningGoods = createWayBillRequest.PersonConsigningGoods
	wayBill.Comments = createWayBillRequest.Comments
	wayBill.TpComments = createWayBillRequest.TpComments
	wayBill.VehicleNumber = createWayBillRequest.VehicleNumber
	wayBill.VehicleType = createWayBillRequest.VehicleType
	wayBill.PickupDate = createWayBillRequest.PickupDate
	wayBill.PalletsSerialNumber = createWayBillRequest.PalletsSerialNumber
	wayBill.AddressOfConsigner = createWayBillRequest.AddressOfConsigner
	wayBill.AddressOfConsignee = createWayBillRequest.AddressOfConsignee
	wayBill.ConsignerRegNumber = createWayBillRequest.ConsignerRegNumber
	wayBill.Carrier = createWayBillRequest.Carrier
	wayBill.VesselType = createWayBillRequest.VesselType
	wayBill.VesselNumber = createWayBillRequest.VesselNumber
	wayBill.ContainerNumber = createWayBillRequest.ContainerNumber
	wayBill.ServiceType = createWayBillRequest.ServiceType
	wayBill.ShipmentModel = createWayBillRequest.ShipmentModel
	wayBill.PalletsQuantity = createWayBillRequest.PalletsQuantity
	wayBill.CartonsQuantity = createWayBillRequest.CartonsQuantity
	wayBill.AssetsQuantity = createWayBillRequest.AssetsQuantity
	wayBill.ShipmentValue = createWayBillRequest.ShipmentValue
	wayBill.EntityName = createWayBillRequest.EntityName
	wayBill.ShipmentCreationDate = createWayBillRequest.ShipmentCreationDate
	wayBill.EWWayBillNumber = createWayBillRequest.EWWayBillNumber
	wayBill.SupportiveDocuments = createWayBillRequest.SupportiveDocuments
	wayBill.ShipmentCreatedBy = createWayBillRequest.ShipmentCreatedBy
	wayBill.ShipmentModifiedDate = createWayBillRequest.ShipmentModifiedDate
	wayBill.ShipmentModifiedBy = createWayBillRequest.ShipmentModifiedBy
	wayBill.WayBillCreationDate = createWayBillRequest.WayBillCreationDate
	wayBill.WayBillCreatedBy = createWayBillRequest.WayBillCreatedBy
	wayBill.WayBillModifiedDate = createWayBillRequest.WayBillModifiedDate
	wayBill.WayBillModifiedBy = createWayBillRequest.WayBillModifiedBy
	dataToStore, _ := json.Marshal(wayBill)
	fmt.Println("WayBill Number From Request", createWayBillRequest.WayBillNumber)

	fmt.Println("WayBill Number", wayBill.WayBillNumber)
	fmt.Println("WayBill to Store", dataToStore)

	err := stub.PutState(wayBill.WayBillNumber, []byte(dataToStore))
	if err != nil {
		fmt.Println("Could not save WayBill to ledger", err)
		return nil, err
	}

	resp := CreateWayBillResponse{}
	resp.Err = "000"
	resp.Message = wayBill.WayBillNumber
	respString, _ := json.Marshal(resp)

	fmt.Println("Successfully saved Way Bill")
	return []byte(respString), nil

}

/************** Create WayBill Ends *************************/

/************** View WayBill Starts ************************/
func ViewWayBill(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering ViewWayBill " + args[0])

	wayBillNumber := args[0]

	wayBilldata, dataerr := fetchWayBillData(stub, wayBillNumber)
	if dataerr == nil {

		dataToStore, _ := json.Marshal(wayBilldata)
		return []byte(dataToStore), nil

	}

	return nil, dataerr

}
func fetchWayBillData(stub shim.ChaincodeStubInterface, wayBillNumber string) (WayBill, error) {
	var wayBill WayBill

	indexByte, err := stub.GetState(wayBillNumber)
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

/************** View WayBill Ends ************************/

/************** Create Export Warehouse WayBill Starts ************************/
func CreateEWWayBill(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering Export Warehouse WayBill ")

	ewWayBillRequest := parseEWWayBillRequest(args[0])

	return processEWWayBill(stub, ewWayBillRequest)

}
func parseEWWayBillRequest(jsondata string) CreateEWWayBillRequest {
	res := CreateEWWayBillRequest{}
	json.Unmarshal([]byte(jsondata), &res)
	fmt.Println(res)
	return res
}
func processEWWayBill(stub shim.ChaincodeStubInterface, createEWWayBillRequest CreateEWWayBillRequest) ([]byte, error) {
	ewWayBill := EWWayBill{}
	//	shipmentIndex := ShipmentIndex{}

	ewWayBill.EwWayBillNumber = createEWWayBillRequest.EwWayBillNumber
	ewWayBill.WayBillsNumber = createEWWayBillRequest.WayBillsNumber
	ewWayBill.ShipmentsNumber = createEWWayBillRequest.ShipmentsNumber
	ewWayBill.CountryFrom = createEWWayBillRequest.CountryFrom
	ewWayBill.CountryTo = createEWWayBillRequest.CountryTo
	ewWayBill.Consigner = createEWWayBillRequest.Consigner
	ewWayBill.Consignee = createEWWayBillRequest.Consignee
	ewWayBill.Custodian = createEWWayBillRequest.Custodian
	ewWayBill.CustodianHistory = createEWWayBillRequest.CustodianHistory
	ewWayBill.CustodianTime = createEWWayBillRequest.CustodianTime
	ewWayBill.PersonConsigningGoods = createEWWayBillRequest.PersonConsigningGoods
	ewWayBill.Comments = createEWWayBillRequest.Comments
	ewWayBill.PalletsSerialNumber = createEWWayBillRequest.PalletsSerialNumber
	ewWayBill.AddressOfConsigner = createEWWayBillRequest.AddressOfConsigner
	ewWayBill.AddressOfConsignee = createEWWayBillRequest.AddressOfConsignee
	ewWayBill.ConsignerRegNumber = createEWWayBillRequest.ConsignerRegNumber
	ewWayBill.VesselType = createEWWayBillRequest.VesselType
	ewWayBill.VesselNumber = createEWWayBillRequest.VesselNumber
	ewWayBill.ContainerNumber = createEWWayBillRequest.ContainerNumber
	ewWayBill.ServiceType = createEWWayBillRequest.ServiceType
	ewWayBill.SupportiveDocuments = createEWWayBillRequest.SupportiveDocuments
	ewWayBill.EwWayBillCreationDate = createEWWayBillRequest.EwWayBillCreationDate
	ewWayBill.EwWayBillCreatedBy = createEWWayBillRequest.EwWayBillCreatedBy
	ewWayBill.EwWayBillModifiedDate = createEWWayBillRequest.EwWayBillModifiedDate
	ewWayBill.EwWayBillModifiedBy = createEWWayBillRequest.EwWayBillModifiedBy

	dataToStore, _ := json.Marshal(ewWayBill)

	err := stub.PutState(ewWayBill.EwWayBillNumber, []byte(dataToStore))
	if err != nil {
		fmt.Println("Could not save Export Warehouse Way Bill to ledger", err)
		return nil, err
	}

	resp := CreateEWWayBillResponse{}
	resp.Err = "000"
	resp.Message = ewWayBill.EwWayBillNumber

	respString, _ := json.Marshal(resp)

	fmt.Println("Successfully saved Export Warehouse Way Bill")
	return []byte(respString), nil

}

/************** Create Export Warehouse WayBill Ends ************************/

/************** View Export Warehouse WayBill Starts ************************/
func ViewEWWayBill(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering ViewEWWayBill " + args[0])

	ewWayBillNumber := args[0]

	emWayBilldata, dataerr := fetchEWWayBillData(stub, ewWayBillNumber)
	if dataerr == nil {

		dataToStore, _ := json.Marshal(emWayBilldata)
		return []byte(dataToStore), nil

	}

	return nil, dataerr

}
func fetchEWWayBillData(stub shim.ChaincodeStubInterface, ewWayBillNumber string) (EWWayBill, error) {
	var ewWayBill EWWayBill

	indexByte, err := stub.GetState(ewWayBillNumber)
	if err != nil {
		fmt.Println("Could not retrive MWayBill ", err)
		return ewWayBill, err
	}

	if marshErr := json.Unmarshal(indexByte, &ewWayBill); marshErr != nil {
		fmt.Println("Could not retrieve master WayBill from ledger", marshErr)
		return ewWayBill, marshErr
	}

	return ewWayBill, nil

}

/************** View Export Warehouse WayBill Ends ************************/

/************** Insert Entity Mapping Starts ************************/
func InsertEntityWayBillMapping(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering Insert Entity Mapping")
	jsondata := args[0]
	createEntityWayBillMappingRequest := CreateEntityWayBillMappingRequest{}
	json.Unmarshal([]byte(jsondata), &createEntityWayBillMappingRequest)
	fmt.Println(createEntityWayBillMappingRequest)

	entityWayBillMapping := EntityWayBillMapping{}
	entityWayBillMapping.WayBillsNumber = createEntityWayBillMappingRequest.WayBillsNumber

	dataToStore, _ := json.Marshal(entityWayBillMapping)

	err := stub.PutState(createEntityWayBillMappingRequest.EntityName, []byte(dataToStore))
	if err != nil {
		fmt.Println("Could not save Entity Mapping to ledger", err)
		return nil, err
	}

	resp := EntityWayBillMappingResponse{}
	resp.Err = "000"
	resp.Message = createEntityWayBillMappingRequest.EntityName

	respString, _ := json.Marshal(resp)

	fmt.Println("Successfully saved Entity Mapping Way Bill")
	return []byte(respString), nil

}

/************** Update Entity Mapping Ends ************************/

/************** Update Entity Mapping Starts ************************/
func UpdateEntityWayBillMapping(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering Update Entity Mapping")
	updatedEntityWayBillMapping := EntityWayBillMapping{}
	entityName := args[0]
	wayBillNumber := args[1]
	entityWayBillMapping, _ := fetchEntityWayBillMappingData(stub, entityName)
	updatedEntityWayBillMapping.WayBillsNumber = append(entityWayBillMapping.WayBillsNumber, wayBillNumber)
	fmt.Println("Updated Entity", updatedEntityWayBillMapping)
	dataToStore, _ := json.Marshal(updatedEntityWayBillMapping)
	err := stub.PutState(entityName, []byte(dataToStore))
	if err != nil {
		fmt.Println("Could not save Entity Mapping to ledger", err)
		return nil, err
	}

	resp := EntityWayBillMappingResponse{}
	resp.Err = "000"
	resp.Message = entityName

	respString, _ := json.Marshal(resp)

	fmt.Println("Successfully saved Entity Mapping Way Bill")
	return []byte(respString), nil

}

/************** Update Entity Mapping Ends ************************/

/************** Get Entity Mapping Starts ************************/
func GetEntityWayBillMapping(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering Get Entity Mapping")
	entityName := args[0]
	wayBillEntityMappingData, dataerr := fetchEntityWayBillMappingData(stub, entityName)
	if dataerr == nil {

		dataToStore, _ := json.Marshal(wayBillEntityMappingData)
		return []byte(dataToStore), nil

	}

	return nil, dataerr
}

func fetchEntityWayBillMappingData(stub shim.ChaincodeStubInterface, entityName string) (EntityWayBillMapping, error) {
	var entityWayBillMapping EntityWayBillMapping

	indexByte, err := stub.GetState(entityName)
	if err != nil {
		fmt.Println("Could not retrive MWayBill ", err)
		return entityWayBillMapping, err
	}

	if marshErr := json.Unmarshal(indexByte, &entityWayBillMapping); marshErr != nil {
		fmt.Println("Could not retrieve master WayBill from ledger", marshErr)
		return entityWayBillMapping, marshErr
	}

	return entityWayBillMapping, nil

}

/************** Get Entity Mapping Ends ************************/

/**************Arshad End new code as per LLD****************/

type WayBillHistory struct {
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Status    string  `json:"status"`
	Timestamp string  `json:"timestamp"`
	Notes     string  `json:"notes"`
	Lat       float64 `json:"lat"`
	Log       float64 `json:"log"`
}

type Shipment struct {
	ShipmentNumber        string           `json:"shipmentNumber"`
	WayBillNo             string           `json:"wayBillNo"`
	WayBillType           string           `json:"wayBillType"`
	PersonConsigningGoods string           `json:"personConsigningGoods"`
	Consigner             string           `json:"consigner"`
	ConsignerAddress      string           `json:"consignerAddress"`
	Consignee             string           `json:"consignee"`
	ConsigneeAddress      string           `json:"consigneeAddress"`
	ConsigneeRegNo        string           `json:"consigneeRegNo"`
	Quantity              string           `json:"quantity"`
	Pallets               []string         `json:"pallets"`
	Cartons               []string         `json:"cartons"`
	Status                string           `json:"status"`
	ModelNo               string           `json:"modelNo"`
	VehicleNumber         string           `json:"vehicleNumber"`
	VehicleType           string           `json:"vehicleType"`
	PickUpTime            string           `json:"pickUpTime"`
	ValueOfGoods          string           `json:"valueOfGoods"`
	ContainerId           string           `json:"containerId"`
	MasterWayBillRef      []string         `json:"masterWayBillRef"`
	WayBillHistorys       []WayBillHistory `json:"wayBillHistorys"`
	Carrier               string           `json:"carrier"`
	Acl                   []string         `json:"acl"`
	CreatedBy             string           `json:"createdBy"`
	Custodian             string           `json:"custodian"`
	CreatedTimeStamp      string           `json:"createdTimeStamp"`
	UpdatedTimeStamp      string           `json:"updatedTimeStamp"`
}

type ShipmentIndex struct {
	ShipmentNumber string
	Status         string
	Acl            []string
}

type AllShipment struct {
	ShipmentIndexArr []ShipmentIndex
}

type AllShipmentDump struct {
	ShipmentIndexArr []string `json:"shipmentIndexArr"`
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

/************** Asset Search Service Starts ************************/

type SearchAssetRequest struct {
	CallingEntityName string `json:"callingEntityName"`
	AssetId           string `json:"assetId"`
}

type SearchAssetResponse struct {
	AssetId        string     `json:"assetId"`
	Modeltype      string     `json:"modeltype"`
	Color          string     `json:"color"`
	CartonId       string     `json:"cartonId"`
	PalletId       string     `json:"palletId"`
	ShipmentDetail []Shipment `json:"shipmentDetail"`
	ErrCode        string     `json:"errCode"`
	ErrMessage     string     `json:"errMessage"`
}

func parseAsset(stub shim.ChaincodeStubInterface, assetId string) (Asset, error) {
	var asset Asset

	assetBytes, err := stub.GetState(assetId)
	if err != nil {
		return asset, err
	} else {
		if marshErr := json.Unmarshal(assetBytes, &asset); marshErr != nil {
			fmt.Println("Could not Unmarshal Asset", marshErr)
			return asset, marshErr
		}
		return asset, nil
	}

}

func retrieveShipment(stub shim.ChaincodeStubInterface, shipmentId string) (Shipment, error) {
	var shipment Shipment

	shipmentBytes, err := stub.GetState(shipmentId)
	if err != nil {
		return shipment, err
	} else {
		if marshErr := json.Unmarshal(shipmentBytes, &shipment); marshErr != nil {
			fmt.Println("Could not Unmarshal Asset", marshErr)
			return shipment, marshErr
		}
		return shipment, nil
	}
}
func PrepareSearchAssetResponse(stub shim.ChaincodeStubInterface, asset Asset) ([]byte, error) {
	var resp SearchAssetResponse
	var shipmentArr []Shipment
	var tmpShipment Shipment
	var err error

	resp.AssetId = asset.AssetId
	resp.Modeltype = asset.Modeltype
	resp.Color = asset.Color
	resp.CartonId = asset.CartonId
	resp.PalletId = asset.PalletId

	lenOfArray := len(asset.ShipmentIds)

	for i := 0; i < lenOfArray; i++ {
		tmpShipment, err = retrieveShipment(stub, asset.ShipmentIds[i])
		if err != nil {
			fmt.Println("Error while retrieveing the Shipment Details", err)
			return nil, err
		} else {
			shipmentArr = append(shipmentArr, tmpShipment)
		}
	}

	resp.ShipmentDetail = shipmentArr
	return json.Marshal(resp)

}

func parseSearchAssetRequest(requestParam string) (SearchAssetRequest, error) {
	var request SearchAssetRequest

	if marshErr := json.Unmarshal([]byte(requestParam), &request); marshErr != nil {
		fmt.Println("Could not Unmarshal Asset", marshErr)
		return request, marshErr
	}
	return request, nil

}

func SearchAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering SearchAsset " + args[0])
	var asset Asset
	var err error
	var request SearchAssetRequest
	var resp SearchAssetResponse

	request, err = parseSearchAssetRequest(args[0])
	if err != nil {
		resp.ErrCode = INVALID_INPUT_ERROR_CODE
		resp.ErrMessage = INVALID_INPUT_ERROR_MSG
		return json.Marshal(resp)
	}

	asset, err = parseAsset(stub, request.AssetId)

	if err != nil {
		resp.ErrCode = NODATA_ERROR_CODE
		resp.ErrMessage = NODATA_ERROR_MSG
		return json.Marshal(resp)
	}

	return PrepareSearchAssetResponse(stub, asset)

}

/************** Asset Search Service Ends ************************/

/************** Carton Search Service Starts ************************/

type SearchCartonRequest struct {
	CallingEntityName string `json:"callingEntityName"`
	CartonId          string `json:"cartonId"`
}

type SearchCartonResponse struct {
	CartonId       string     `json:"cartonId"`
	PalletId       string     `json:"palletId"`
	ShipmentDetail []Shipment `json:"shipmentDetail"`
	ErrCode        string     `json:"errCode"`
	ErrMessage     string     `json:"errMessage"`
}

func parseSearchCartonRequest(requestParam string) (SearchCartonRequest, error) {
	var request SearchCartonRequest

	if marshErr := json.Unmarshal([]byte(requestParam), &request); marshErr != nil {
		fmt.Println("Could not Unmarshal Asset", marshErr)
		return request, marshErr
	}
	return request, nil

}

func parseCarton(stub shim.ChaincodeStubInterface, cartonId string) (Carton, error) {
	var carton Carton

	cartonBytes, err := stub.GetState(cartonId)
	if err != nil {
		return carton, err
	} else {
		if marshErr := json.Unmarshal(cartonBytes, &carton); marshErr != nil {
			fmt.Println("Could not Unmarshal Asset", marshErr)
			return carton, marshErr
		}
		return carton, nil
	}

}

func PrepareSearchCartontResponse(stub shim.ChaincodeStubInterface, carton Carton) ([]byte, error) {
	var resp SearchCartonResponse
	var shipmentArr []Shipment
	var tmpShipment Shipment
	var err error

	resp.CartonId = carton.CartonId
	resp.PalletId = carton.PalletId

	lenOfArray := len(carton.ShipmentIds)

	for i := 0; i < lenOfArray; i++ {
		tmpShipment, err = retrieveShipment(stub, carton.ShipmentIds[i])
		if err != nil {
			fmt.Println("Error while retrieveing the Shipment Details", err)
			return nil, err
		}
		shipmentArr = append(shipmentArr, tmpShipment)
	}

	resp.ShipmentDetail = shipmentArr
	return json.Marshal(resp)

}

func SearchCarton(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering SearchCarton " + args[0])
	var carton Carton
	var err error
	var request SearchCartonRequest
	var resp SearchCartonResponse

	request, err = parseSearchCartonRequest(args[0])
	if err != nil {
		resp.ErrCode = INVALID_INPUT_ERROR_CODE
		resp.ErrMessage = INVALID_INPUT_ERROR_MSG
		return json.Marshal(resp)
	}

	carton, err = parseCarton(stub, request.CartonId)

	if err != nil {
		resp.ErrCode = NODATA_ERROR_CODE
		resp.ErrMessage = NODATA_ERROR_MSG
		return json.Marshal(resp)
	}

	return PrepareSearchCartontResponse(stub, carton)

}

/************** Carton Search Service Ends ************************/

/************** Pallet Search Service Starts ************************/

type SearchPalletRequest struct {
	CallingEntityName string `json:"callingEntityName"`
	PalletId          string `json:"palletId"`
}

type SearchPalletResponse struct {
	PalletId       string     `json:"palletId"`
	ShipmentDetail []Shipment `json:"shipmentDetail"`
	ErrCode        string     `json:"errCode"`
	ErrMessage     string     `json:"errMessage"`
}

func parseSearchPalletRequest(requestParam string) (SearchPalletRequest, error) {
	var request SearchPalletRequest

	if marshErr := json.Unmarshal([]byte(requestParam), &request); marshErr != nil {
		fmt.Println("Could not Unmarshal Asset", marshErr)
		return request, marshErr
	}
	return request, nil

}

func parsePallet(stub shim.ChaincodeStubInterface, palletId string) (Pallet, error) {

	var pallet Pallet

	palletBytes, err := stub.GetState(palletId)
	if err != nil {
		return pallet, err
	} else {
		if marshErr := json.Unmarshal(palletBytes, &pallet); marshErr != nil {
			fmt.Println("Could not Unmarshal Asset", marshErr)
			return pallet, marshErr
		}
		return pallet, nil
	}

}

func PrepareSearchPalletResponse(stub shim.ChaincodeStubInterface, pallet Pallet) ([]byte, error) {
	var resp SearchPalletResponse
	var shipmentArr []Shipment
	var tmpShipment Shipment
	var err error

	resp.PalletId = pallet.PalletId

	lenOfArray := len(pallet.ShipmentIds)

	for i := 0; i < lenOfArray; i++ {
		tmpShipment, err = retrieveShipment(stub, pallet.ShipmentIds[i])
		if err != nil {
			fmt.Println("Error while retrieveing the Shipment Details", err)
			return nil, err
		}
		shipmentArr = append(shipmentArr, tmpShipment)
	}

	resp.ShipmentDetail = shipmentArr
	return json.Marshal(resp)

}

func SearchPallet(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering SearchPallet " + args[0])
	var pallet Pallet
	var err error
	var request SearchPalletRequest
	var resp SearchPalletResponse

	request, err = parseSearchPalletRequest(args[0])
	if err != nil {
		resp.ErrCode = INVALID_INPUT_ERROR_CODE
		resp.ErrMessage = INVALID_INPUT_ERROR_MSG
		return json.Marshal(resp)
	}

	pallet, err = parsePallet(stub, request.PalletId)

	if err != nil {
		resp.ErrCode = NODATA_ERROR_CODE
		resp.ErrMessage = NODATA_ERROR_MSG
		return json.Marshal(resp)
	}

	return PrepareSearchPalletResponse(stub, pallet)

}

/************** Pallet Search Service Ends ************************/

/************** Date Search Service Starts ************************/

type SearchDateRequest struct {
	CallingEntityName string `json:"callingEntityName"`
	StartDate         string `json:"startDate"`
	EndDate           string `json:"endDate"`
}

type SearchDateResponse struct {
	ShipmentDetail []Shipment `json:"shipmentDetail"`
}

func parseAllShipmentDump() (AllShipmentDump, error) {
	var dump AllShipmentDump

	if marshErr := json.Unmarshal([]byte("ALL_SHIPMENT_DUMP"), &dump); marshErr != nil {
		fmt.Println("Could not Unmarshal Asset", marshErr)
		return dump, marshErr
	}
	return dump, nil

}

func SearchDateRange(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	//var shipmentDump AllShipmentDump
	var err error
	var shipmentArr []Shipment
	var tmpShipment Shipment
	var resp SearchDateResponse

	/*shipmentDump, err = parseAllShipmentDump()
	if err != nil {
		return nil, err
	}*/

	lenOfArray := len(args)

	for i := 0; i < lenOfArray; i++ {
		tmpShipment, err = retrieveShipment(stub, args[i])
		if err != nil {
			fmt.Println("Error while retrieveing the Shipment Details", err)
			return nil, err
		}
		shipmentArr = append(shipmentArr, tmpShipment)
	}
	resp.ShipmentDetail = shipmentArr

	return json.Marshal(resp)

}

/************** Date Search Service Ends ************************/

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

func Initialize(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

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
	} else if function == "DumpData" {
		return DumpData(stub, args)
	} else if function == "CreateWayBill" {
		return CreateWayBill(stub, args)
	} else if function == "CreateEWWayBill" {
		return CreateEWWayBill(stub, args)
	} else if function == "InsertEntityWayBillMapping" {
		return InsertEntityWayBillMapping(stub, args)
	} else if function == "UpdateEntityWayBillMapping" {
		return UpdateEntityWayBillMapping(stub, args)
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
	} else if function == "SearchAsset" {
		return SearchAsset(stub, args)
	} else if function == "SearchCarton" {
		return SearchCarton(stub, args)
	} else if function == "SearchPallet" {
		return SearchPallet(stub, args)
	} else if function == "SearchDateRange" {
		return SearchDateRange(stub, args)
	} else if function == "ViewWayBill" {
		return ViewWayBill(stub, args)
	} else if function == "ViewEWWayBill" {
		return ViewEWWayBill(stub, args)
	} else if function == "GetEntityWayBillMapping" {
		return GetEntityWayBillMapping(stub, args)
	} else {
		return nil, errors.New("Invalid function name " + function)
	}
}

func main() {
	Initialize(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	err := shim.Start(new(B4SCChaincode))
	if err != nil {
		fmt.Println("Could not start B4SCChaincode")
	} else {
		fmt.Println("B4SCChaincode successfully started")
	}
}
