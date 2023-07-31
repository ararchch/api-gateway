namespace go name.management

struct NameRequest {
    1: required string Name;
}

struct NameResponse {
    1: string HelloResp;
}

service NameManagement {
    NameResponse helloName(1: NameRequest req);
}

