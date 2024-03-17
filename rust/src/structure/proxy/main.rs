use super::{nginx::NginxServer, server::Server};

fn main() {
    let app_status = "/app/status";
    let create_user = "/create/user";

    let mut nginx = NginxServer::new();

    let (code, body) = nginx.handle_request(app_status, "GET");
    let (code, body) = nginx.handle_request(app_status, "GET");
    let (code, body) = nginx.handle_request(app_status, "GET");
    let (code, body) = nginx.handle_request(create_user, "POST");
    let (code, body) = nginx.handle_request(create_user, "GET");
}
