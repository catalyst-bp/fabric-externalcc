# fabric-externalcc

Repository contains simple Hyperledger Fabric external chaincode example named "Hello world".

Two environment variables `CHAINCODE_ID` and `CHAINCODE_ADDRESS` are necessary to run this chaincode on Catalyst Blockchain Platform.

To check work of deployed chaincode you can use listed below invoke and query functions:
```
{ "function":"invoke", "args":["John"] }
{ "function":"query", "args":["John"] }
```

To build docker image run the following command:
```
docker build . -t fabric-externalcc:latest
```