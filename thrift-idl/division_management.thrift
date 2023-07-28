// tells thrift to generate scaffolding code in ‘Go’ in the   subdirectory kitex_gen/division/api
namespace go division.management// directories generated will be named following this convention

// creates a new struct ‘Division request’ with the above implementation
struct DivisionRequest {
1: required string FirstNum;
2: required string SecondNum;
} 

// creates a new struct ‘Division Response’ with the above implementation
struct DivisionResponse {
1: string Quotient;
} 

// Defines the service name and the method signature of the divideNumbers method offerred by the service
service DivisionManagement {
DivisionResponse divideNumbers(1: DivisionRequest req);
} 