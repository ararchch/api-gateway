namespace go hertz-api

struct AdditionRequest {
    1: required i32 FirstNum (api.body="firstnum");
    2: required i32 SecondNum (api.body="secondnum")
}

struct AdditionResponse {
    1: string Sum;
}

service StudentApi {
   AdditionResponse addNumbers(1: AdditionRequest req) (api.post="/add");
}