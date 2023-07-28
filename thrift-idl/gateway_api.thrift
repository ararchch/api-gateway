namespace go api

struct AdditionRequest {
1: required string FirstNum (api.body="FirstNum");
2: required string SecondNum (api.body="SecondNum")
}

struct AdditionResponse {
1: string Sum;
}

struct MultiplicationRequest {
1: required string FirstNum (api.body="FirstNum");
2: required string SecondNum (api.body="SecondNum")
}

struct MultiplicationResponse {
1: string Product;
}

// request structure that the user sends to the gateway
struct DivisionRequest { 
1: required string FirstNum (api.body="FirstNum");
2: required string SecondNum (api.body="SecondNum")
}

// response structure sent by gateway to the user
struct DivisionResponse {
1: string Quotient;
}

service Gateway {
AdditionResponse addNumbers(1: AdditionRequest req) (api.post="/add");
MultiplicationResponse multiplyNumbers(1: MultiplicationRequest req) (api.post="/multiply");
DivisionResponse divideNumbers(1: DivisionRequest req) (api.post="/divide"); // endpoint and method details
}