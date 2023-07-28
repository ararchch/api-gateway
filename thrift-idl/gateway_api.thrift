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

service Gateway {
AdditionResponse addNumbers(1: AdditionRequest req) (api.post="/add");
MultiplicationResponse multiplyNumbers(1: MultiplicationRequest req) (api.post="/multiply");
}