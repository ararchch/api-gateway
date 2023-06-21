namespace go addition.management

struct AdditionRequest {
    1: required i32 FirstNum;
    2: required i32 SecondNum;
}

struct AdditionResponse {
    1: i32 Sum;
}

service AdditionManagement {
    AdditionResponse addNumbers(1: AdditionRequest req);
}

