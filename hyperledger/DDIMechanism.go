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
type DDIMechanism struct {
	MechanismID int `json:"MechanismID"`
	MechanismName string `json:"MechanismName"`
	MechanismType string `json:"MechanismType"`
	MechanismCategoryType string `json:"MechanismCategoryType"`	
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
		Contract.initLedger(stub, args)
	}  else if fnType == "createMechanism" {
		Contract.createMechanism(stub, args)
	} else if fnType == "revokeMechanism" {
		Contract.revokeMechanism(stub, args)
	} else if fnType == "findByID" {
		Contract.findByID(stub, args)
	}  else {
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
			MechanismID := params[1]
			mechanismid, err := strconv.Atoi(MechanismID)
	
			Mechanism := []MechanismInt{
				MechanismInt{
					MechanismID : mechanismid,
					MechanismName : params[2],
					MechanismType : params[3]
					MechanismCategoryType : params[4]
				},
			}
			
			MechanismBuffer, err := json.Marshal (Mechanism)
			if err != nil {
				ftm.Println("Error Occurred: ",err.Error())
			}
	
			fmt.Println("Created and now prepareing your transaction for commit:", MechanismBuffer)
	
			stub.PutState(mechanismid,MechanismBuffer)
	
			return shim.Success(nil)
	
		} else{
			return shim.Error("No Argument supplied to Create Mechanism")
		}

	} else{
		return shim.Error("No Argument supplied to Create Mechanism")
	}
}
func (Contract *DrugDrugIntraction) revokeMechanism(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	MechanismToRemoveargs := stub.getStringArgs()
		
	if params != 3{
		MechanismID := MechanismToRemoveargs[1]
		Mechanismid, err := strconv.Atoi(MechanismID)
			if err != nil {
				ftm.Println("Error Occurred: ",err.Error())
			}
		mName := MechanismToRemoveargs[2]

		existingMechanism,err := stub.GetState(Mechanismid)
			if err != nil {
				return shim.Error("Mechanism not Exist! Error message:"+ err.Error())
			}
		
		Mechanism := Mechanism{}
		json.Unmarshal (existingMechanism,Mechanism)
			if existingMechanism != nil{
				if Mechanism.MechanismName = mName {
					stub.DelState(Mechanismid)
				}
			}

		return shim.Success(nil)

	} else{
		return shim.Error("No Argument supplied to Remove Mechanism")
	}
}
func (Contract *SmartService) findByID(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	params := stub.getStringArgs()

	if params !="" {

		MechanismID := params[1]

		MechanismFound,err := stub.GetState(MechanismID)
		if err != nil {
			ftm.Println("Could not find Mechanism: ",err.Error())
		}
		if len(MechanismFound) == 0 {
			return shim.Error("Mechanism Not exist!")
		}
		ftm.Println("Mechanism found for the ID: ",MechanismFound)

	} else{
		return shim.Error("No Argument supplied to Find Mechanism")
	}
}
