namespace php MicroService
namespace go phalcon.demo.micro.service

struct UserDTO {
    1:i64 user_id,
    2:string name
}

service App {
    string version()
}

service User{
    UserDTO getByUserId(1:i64 user_id)
    bool save(1:i64 user_id, 2:string name)
}