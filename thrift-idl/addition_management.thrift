namespace go addition.management

struct AdditionRequest {
    1: required string FirstNum;
    2: required string SecondNum;
}

struct AdditionResponse {
    1: string Sum;
}

service AdditionManagement {
    AdditionResponse addNumbers(1: AdditionRequest req);
}

