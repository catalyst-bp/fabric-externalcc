package main

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"os"
)

type HelloWorld struct {}

func (h *HelloWorld) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (h *HelloWorld) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	
	fmt.Printf("Invoking %s with args: %v", function, args)

	if function == "invoke" {
		return h.invoke(stub, args)
	} else if function == "query" {
		return h.query(stub, args)
	}

	return shim.Error("Invalid function call")
}

func (h *HelloWorld) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of args")
	}
	
	fmt.Printf("Querying %s with args: %v", function, args)

	key := args[0]

	if err := stub.PutState(key, []byte("Hello World!")); err != nil {
		return shim.Error(err.Error())
	}

	if err := stub.SetEvent("notification", []byte(fmt.Sprintf("key %s successfully saved", key))); err != nil {
		return shim.Error("error happened emitting event: " + err.Error())
	}

	return shim.Success(nil)
}

func (h *HelloWorld) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of args")
	}

	key := args[0]

	value, err := stub.GetState(key)
	if err != nil {
		return shim.Error("Error getting state: " + err.Error())
	}

	response := "{\"Greetings for " + key + "\":\"" + string(value) + "\"}"
	return shim.Success([]byte(response))
}

func main() {
	ccid := os.Getenv("CHAINCODE_ID")
	address := os.Getenv("CHAINCODE_ADDRESS")

	server := &shim.ChaincodeServer{
		CCID:     ccid,
		Address:  address,
		CC:       &HelloWorld{},
		TLSProps: shim.TLSProperties{
			Disabled: true,
		},
	}
	
	fmt.Printf("Started ccid %s on %s", ccid, address)

	if err := server.Start(); err != nil {
		fmt.Printf("Error starting HelloWorld chaincode: %s", err)
	}
}
