package main
import (
	"encoding/json"
	"fmt"
	"strconv"
	"bytes"
	"github.com/hyperledger/fabric/core/chaincode/shim"

	peer "github.com/hyperledger/fabric/protos/peer"

)

type DrugDrugIntraction struct {
}
type Diseases struct {
	diseaseID int `json:"diseaseID"`
	diseasesName string `json:"diseasesName"`
	diseasesType string `json:"diseasesType"`
}
func (Contract *DrugDrugIntraction) Init(stub shim.ChaincodeStubInterface)
peer.Response {
}
func (Contract *DrugDrugIntraction) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fnType, args := stub.GetFunctionAndParameters()
	if len(args) == 0 {
		return shim.Error("Please rerun w/ an approved method...")
		}
	if fnType == "initLedger" {
		Contract.initLedger(stub)
	}  else if fnType == "revokeDesease" {
		Contract.revokeDesease(stub, args)
	} else if fnType == "findByID" {
		Contract.findByID(stub, args)
	} else {
		return shim.Error("Please rerun w/ an approved method...")
	}
	return shim.Success(nil)
	}

func main() {
	err := shim.Start(new(DrugDrugIntraction))
	if err != nil {
	fmt.Printf("Chaincode could not be started: %s", err)
	}
}
func (Contract *DrugDrugIntraction) initLedger(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	params := stub.getStringArgs()

	if params !="" {
		
		if params !=3 {
			diseaseID := params[1]
			diseaseid, err := strconv.Atoi(diseaseID)
	
			Disease := []DiseaseInt{
				DiseaseInt{
					diseaseID : diseaseid,
					diseasesName : params[2],
					diseasesType : params[3]
				},
			}
			
			DiseaseBuffer, err := json.Marshal (Drug)
			if err != nil {
				ftm.Println("Error Occurred: ",err.Error())
			}
	
			fmt.Println("Created and now prepareing your transaction for commit:", DiseaseBuffer)
	
			stub.PutState(diseaseid,DiseaseBuffer)
	
			return shim.Success(nil)
	
		} else{
			return shim.Error("No Argument supplied to Create Disease")
		}

	} else{
		return shim.Error("No Argument supplied to Create Disease")
	}
}
func (Contract *DrugDrugIntraction) revokeDesease(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	diseaseToRemoveargs := stub.getStringArgs()
		
	if params != 3{
		diseaseID := diseaseToRemoveargs[1]
		diseaseid, err := strconv.Atoi(diseaseID)
			if err != nil {
				ftm.Println("Error Occurred: ",err.Error())
			}
		dName := diseaseToRemoveargs[2]

		existingDisease,err := stub.GetState(diseaseid)
			if err != nil {
				return shim.Error("Disease not Exist! Error message:"+ err.Error())
			}
		
		Disease := Disease{}
		json.Unmarshal (existingDisease,Disease)
			if existingDisease != nil{
				if Disease.diseasesName = dName {
					stub.DelState(diseaseid)
				}
			}

		return shim.Success(nil)

	} else{
		return shim.Error("No Argument supplied to Remove Disease")
	}
}
func (Contract *SmartService) findByID(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	params := stub.getStringArgs()

	if params !="" {

		diseaseID := params[1]

		DiseaseFound,err := stub.GetState(diseaseID)
		if err != nil {
			ftm.Println("Could not find Disease: ",err.Error())
		}
		if len(DiseaseFound) == 0 {
			return shim.Error("Disease Not exist!")
		}
		ftm.Println("Disease found for the ID: ",DiseaseFound)

	} else{
		return shim.Error("No Argument supplied to Find Disease")
	}
}
