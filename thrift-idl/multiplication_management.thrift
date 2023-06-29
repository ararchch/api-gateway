namespace go multiplication.management

struct MultiplicationRequest {
    1: required string FirstNum;
    2: required string SecondNum;
}

struct MultiplicationResponse {
    1: string Product;
}

service MultiplicationManagement {
    MultiplicationResponse multiplyNumbers(1: MultiplicationRequest req);
}