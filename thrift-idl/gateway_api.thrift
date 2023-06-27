namespace go api

struct AdditionRequest {
    1: required string FirstNum (api.body="firstnum");
    2: required string SecondNum (api.body="secondnum")
}

struct AdditionResponse {
    1: string Sum;
}

service AdditionApi {
   AdditionResponse addNumbers(1: AdditionRequest req) (api.post="/add");
}