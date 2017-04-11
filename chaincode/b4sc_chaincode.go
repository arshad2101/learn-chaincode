/*****Main Chaicode to start the execution*****

/*****************************************************/
package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type B4SCChaincode struct {
}

func NonDeterministic(stub shim.ChaincodeStubInterface) ([]byte, error) {

	stub.PutState("1", []byte("Arshad"))
	return nil, nil
}

func GetNonDeterministic(stub shim.ChaincodeStubInterface) ([]byte, error) {
	output, _ := stub.GetState("1")
	return output, nil
}

// Init resets all the things
func (t *B4SCChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Inside INIT for test chaincode")

	return nil, nil
}

func (t *B4SCChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	/*if function == "Init" {
		return Init(stub, function, args)
	}else*/
	if function == "GetNonDeterministic" {
		return GetNonDeterministic(stub)
	} else {
		return nil, errors.New("Invalid function name " + function)
	}
	//return nil, nil
}

func (t *B4SCChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "GetNonDeterministic" {
		return GetNonDeterministic(stub)
	}
	return nil, errors.New("Invalid function name " + function)

}

func main() {

	err := shim.Start(new(B4SCChaincode))
	if err != nil {
		fmt.Println("Could not start B4SCChaincode")
	} else {
		fmt.Println("B4SCChaincode successfully started")
	}
}
