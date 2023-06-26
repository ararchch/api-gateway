namespace go multiplication.management

struct MultiplicationRequest {
    1: required i32 FirstNum;
    2: required i32 SecondNum;
}

struct MultiplicationResponse {
    1: i32 Product;
}

service MultiplicationManagement {
    MultiplicationResponse multiplyNumbers(1: MultiplicationRequest req);
}
