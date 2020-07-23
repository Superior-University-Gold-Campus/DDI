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
type Drug struct {
	DrugId int `json:"DrugID"`
	DrugName string `json:"DrugName"`
	AdverseEffecType string `json:"AdverseEffecType"`
	AgeGroupID string `json:"AgeGroupID"`
	DDIMechanism string `json:"DDIMechanism"`
	ReactionDuration string `json:"ReactionDuration"`
	DrugAdministrationMethod string `json:"DrugAdministrationMethod"`
	AdversEffects string `json:"AdversEffects"`
	Cautitions string `json:"Cautitions"`
	DDIType string `json:"DDIType"`
	Diseases string `json:"Diseases"`
	DosageForm string `json:"DosageForm"`
	DrugreactionFrequency string `json:"DrugreactionFrequency"`
	Ingredients string `json:"Ingredients"`
	InteractionLevel string `json:"InteractionLevel"`
	PharmaCompany string `json:"PharmaCompany"`
	PhysologicalEffects string `json:"PhysologicalEffects"`
	SideEffect string `json:"SideEffect"`
	UndesirabeEffects string `json:"UndesirabeEffects"`
	Warning string `json:"Warning"`
	Intractwith string `json:"Intractwith"`
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
	}  else if fnType == "hasAdministrationMethod" {
		Contract.hasAdministrationMethod(stub, args)
	} else if fnType == "hasAdversEffect" {
		Contract.hasAdverseEffect(stub, args)
	} else if fnType == "hasDDIType" {
		Contract.hasDDIType(stub, args)
	} else if fnType == "hasDosageForm" {
		Contract.hasDosageForm(stub,args)
	} else if fnType == "hasDrugMechanism" {
		Contract.hasDrugMechanism(stub, args)
	} else if fnType == "hasInteractionLevel" {
		Contract.hasInteractionLevel(stub, args)
	} else if fnType == "hasReactionDuration" {
		Contract.hasreactionDuration(stub, args)
	} else if fnType == "hasReactionFrequency" {
		Contract.hasReactionFrequency(stub, args)
	} else if fnType == "hasSideEffects" {
		Contract.hasSideEffects(stub, args)
	}else {
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

	if params != 21 {
		DrugId := params[1]
		drugid, err := strconv.Atoi(DrugId)
		AgeGroupID := params[4]
		ageGroup, err := strconv.Atoi(AgeGroupID)

		Drug := []DrugIntraction{
			DrugIntraction{
				DrugId : drugid,
				DrugName : params[2],
				AdverseEffecType : params[3],
				AgeGroupID : ageGroup,
				DDIMechanism : params[5],
				ReactionDuration : params[6],
				DrugAdministrationMethod : params[7],
				AdversEffects : params[8],
				Cautitions : params[9],
				DDIType : params[10],
				Diseases : params[11],
				DosageForm : params[12],
				DrugreactionFrequency : params[13],
				Ingredients : params[14],
				InteractionLevel : params[15],
				PharmaCompany : params[16],
				PhysologicalEffects : params[17],
				SideEffect : params[18],
				UndesirabeEffects : params[19],
				Warning : params[20],
				Intractwith : params[21]
			},
		}
		
		DrugBuffer, err := json.Marshal (Drug)
		if err != nil {
			ftm.Println("Error Occurred: ",err.Error())
		}

		fmt.Println("Created and now prepareing your transaction for commit:", Drug)

		stub.PutState(drugid,DrugBuffer)

		return shim.Success(nil)

	} else{
		return shim.Error("No Argument supplied to Create Drug")
	}
}
func (Contract *DrugDrugIntraction) hasAdministrationMethod(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	params := stub.getStringArgs()

	if params != "" {

		DrugId := params[1]
		drugid, err := strconv.Atoi(DrugId)
			
			Method, err := stub.GetState(drugid)
			if err  !=  nil {
				ftm.Println("Couldn't find the drug, Please try again",err.Error())
			}
			ftm.Println("Administration Method for the Drug is:",Method)

	} else{
		return shim.Error("No Argument supplied to Fetch Administration Method")
	}
}
func (Contract *DrugDrugIntraction) hasAdverseEffect(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	params := stub.getStringArgs()

	if params != "" {

		DrugId := params[1]
		drugid, err := strconv.Atoi(DrugId)

		Effect, err := stub.GetState(drugid)
		if err  !=  nil {
			ftm.Println("Couldn't find the drug, Please try again",err.Error())
		}
		ftm.Println("Effect for the Drug is:",Effect)

	} else{
		return shim.Error("No Argument supplied to Fetch Effect")
	}
}
func (Contract *DrugDrugIntraction) hasDDIType(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	params := stub.getStringArgs()

	if params != "" {

		DrugId := params[1]
		drugid, err := strconv.Atoi(DrugId)

		DDIType, err := stub.GetState(drugid)
		if err  !=  nil {
			ftm.Println("Couldn't find the drug, Please try again",err.Error())
		}
		ftm.Println("DDI type for the Drug is:",DDIType)

	} else{
		return shim.Error("No Argument supplied to Fetch DDI type")
	}
}
func (Contract *DrugDrugIntraction) hasDosageForm(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	params := stub.getStringArgs()

	if params != "" {

		DrugId := params[1]
		drugid, err := strconv.Atoi(DrugId)

		DosageForm, err := stub.GetState(drugid)
		if err  !=  nil {
			ftm.Println("Couldn't find the drug, Please try again",err.Error())
		}
		ftm.Println("Dosage Form for the Drug is:",DosageForm)

	} else{
		return shim.Error("No Argument supplied to Fetch Dosage Form")
	}
}
func (Contract *DrugDrugIntraction) hasDrugMechanism(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	params := stub.getStringArgs()

	if params != "" {

		DrugId := params[1]
		drugid, err := strconv.Atoi(DrugId)

		DrugMechanism, err := stub.GetState(drugid)
		if err  !=  nil {
			ftm.Println("Couldn't find the drug, Please try again",err.Error())
		}
		ftm.Println("Drug Mechanism for the Drug is:",DrugMechanism)

	} else{
		return shim.Error("No Argument supplied to Fetch Drug Mechanism")
	}
}
func (Contract *DrugDrugIntraction) hasInteractionLevel(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	params := stub.getStringArgs()

	if params != "" {

		DrugId := params[1]
		drugid, err := strconv.Atoi(DrugId)

		InteractionLevel, err := stub.GetState(drugid)
		if err  !=  nil {
			ftm.Println("Couldn't find the drug, Please try again",err.Error())
		}
		ftm.Println("Interaction Level for the Drug is:",InteractionLevel)

	} else{
		return shim.Error("No Argument supplied to Fetch Interaction Level")
	}
}
func (Contract *DrugDrugIntraction) hasReactionDuration(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	params := stub.getStringArgs()

	if params != "" {

		DrugId := params[1]
		drugid, err := strconv.Atoi(DrugId)

		ReactionDuration, err := stub.GetState(drugid)
		if err  !=  nil {
			ftm.Println("Couldn't find the drug, Please try again",err.Error())
		}
		ftm.Println("Reaction Duration for the Drug is:",ReactionDuration)

	} else{
		return shim.Error("No Argument supplied to Fetch Reaction Duration")
	}
}
func (Contract *DrugDrugIntraction) hasReactionFrequency(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	params := stub.getStringArgs()

	if params != "" {

		DrugId := params[1]
		drugid, err := strconv.Atoi(DrugId)

		ReactionFrequency, err := stub.GetState(drugid)
		if err  !=  nil {
			ftm.Println("Couldn't find the drug, Please try again",err.Error())
		}
		ftm.Println("Reaction Frequency for the Drug is:",ReactionFrequency)

	} else{
		return shim.Error("No Argument supplied to Fetch Reaction Frequency")
	}
}