/*****All structs used in the chaincode*****
Structs Involved
BlockchainResponse
AssetDetails
CartonDetails
PalletDetails
ShipmentWayBill
EWWayBill
EntityWayBillMapping
CreateEntityWayBillMappingRequest
WayBillShipmentMapping

Author: Mohd Arshad
Dated: 30/7/2017
/*****************************************************/
package main

const (
	SHIPMENT   = "SHIPMENT"
	WAYBILL    = "WAYBILL"
	DCSHIPMENT = "DCSHIPMENT"
	DCWAYBILL  = "DCWAYBILL"
	EWWAYBILL  = "EWWAYBILL"
)

type BlockchainResponse struct {
	Err     string `json:"err"`
	ErrMsg  string `json:"errMsg"`
	Message string `json:"message"`
}

type AssetDetails struct {
	AssetSerialNo      string
	AssetModel         string
	AssetType          string
	AssetMake          string
	AssetCOO           string
	AssetMaufacture    string
	AssetStatus        string
	CreatedBy          string
	CreatedDate        string
	ModifiedBy         string
	ModifiedDate       string
	PalletSerialNumber string
	CartonSerialNumber string
	MshipmentNumber    string
	DcShipmentNumber   string
	MwayBillNumber     string
	DcWayBillNumber    string
	EwWayBillNumber    string
	MShipmentDate      string
	DcShipmentDate     string
	MWayBillDate       string
	DcWayBillDate      string
	EwWayBillDate      string
}

type CartonDetails struct {
	CartonSerialNo     string
	CartonModel        string
	CartonStatus       string
	CartonCreationDate string
	PalletSerialNumber string
	AssetsSerialNumber []string
	MshipmentNumber    string
	DcShipmentNumber   string
	MwayBillNumber     string
	DcWayBillNumber    string
	EwWayBillNumber    string
	Dimensions         string
	Weight             string
	MShipmentDate      string
	DcShipmentDate     string
	MWayBillDate       string
	DcWayBillDate      string
	EwWayBillDate      string
}

type PalletDetails struct {
	PalletSerialNo     string
	PalletModel        string
	PalletStatus       string
	CartonSerialNumber []string
	PalletCreationDate string
	AssetsSerialNumber []string
	MshipmentNumber    string
	DcShipmentNumber   string
	MwayBillNumber     string
	DcWayBillNumber    string
	EwWayBillNumber    string
	Dimensions         string
	Weight             string
	MShipmentDate      string
	DcShipmentDate     string
	MWayBillDate       string
	DcWayBillDate      string
	EwWayBillDate      string
}

/*This is common struct across Shipment and Waybill*/
type ShipmentWayBill struct {
	WayBillNumber         string   `json:"wayBillNumber"`
	ShipmentNumber        string   `json:"shipmentNumber"`
	CountryFrom           string   `json:"countryFrom"`
	CountryTo             string   `json:"countryTo"`
	Consigner             string   `json:"consigner"`
	Consignee             string   `json:"consignee"`
	Custodian             string   `json:"custodian"`
	CustodianHistory      []string `json:"custodianHistory"`
	PersonConsigningGoods string   `json:"personConsigningGoods"`
	Comments              string   `json:"comments"`
	TpComments            string   `json:"tpComments"`
	VehicleNumber         string   `json:"vehicleNumber"`
	VehicleType           string   `json:"vehicleType"`
	PickupDate            string   `json:"pickupDate"`
	PalletsSerialNumber   []string `json:"palletsSerialNumber"`
	AddressOfConsigner    string   `json:"addressOfConsigner"`
	AddressOfConsignee    string   `json:"addressOfConsignee"`
	ConsignerRegNumber    string   `json:"consignerRegNumber"`
	Carrier               string   `json:"carrier"`
	VesselType            string   `json:"vesselType"`
	VesselNumber          string   `json:"vesselNumber"`
	ContainerNumber       string   `json:"containerNumber"`
	ServiceType           string   `json:"serviceType"`
	ShipmentModel         string   `json:"shipmentModel"`
	PalletsQuantity       string   `json:"palletsQuantity"`
	CartonsQuantity       string   `json:"cartonsQuantity"`
	AssetsQuantity        string   `json:"assetsQuantity"`
	ShipmentValue         string   `json:"shipmentValue"`
	EntityName            string   `json:"entityName"`
	ShipmentCreationDate  string   `json:"shipmentCreationDate"`
	EWWayBillNumber       string   `json:"eWWayBillNumber"`
	SupportiveDocuments   []string `json:"supportiveDocuments"`
	ShipmentCreatedBy     string   `json:"shipmentCreatedBy"`
	ShipmentModifiedDate  string   `json:"shipmentModifiedDate"`
	ShipmentModifiedBy    string   `json:"shipmentModifiedBy"`
	WayBillCreationDate   string   `json:"wayBillCreationDate"`
	WayBillCreatedBy      string   `json:"wayBillCreatedBy"`
	WayBillModifiedDate   string   `json:"wayBillModifiedDate"`
	WayBillModifiedBy     string   `json:"wayBillModifiedBy"`
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

type EntityWayBillMapping struct {
	WayBillsNumber []string
}
type CreateEntityWayBillMappingRequest struct {
	EntityName     string
	WayBillsNumber []string
}
type WayBillShipmentMapping struct {
	DCWayBillsNumber string
	DCShipmentNumber string
}

//storing compliance document mdetadata and hash
type ComplianceDocument struct {
	compliance_id      string
	manufacturer       string
	regulator          string
	documentTitle      string
	document_mime_type string
	documentHash       string
	documentType       string
	createdBy          string
	createdDate        string
}

//mapping for entity and corresponding document
type EntityComplianceDocMapping struct {
	complianceIds []string
}

//collection of all the compliance document ids
type ComplianceIds struct {
	complianceIds []string
}

//list of compliance document
type ComplianceDocumentList struct {
	complianceDocumentList []ComplianceDocument
}
