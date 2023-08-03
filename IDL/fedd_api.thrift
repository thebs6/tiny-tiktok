namespace go douyin.core



service HelloService {
    HelloResp HelloMethod(1: HelloReq request) (api.get="/hello");
}
struct HelloReq {}
struct HelloResp {}