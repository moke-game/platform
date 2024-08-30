import {Client, StatusOK} from 'k6/net/grpc';
import {check, sleep} from 'k6';


const client = new Client();
client.load(['../../api/'], './auth/auth.proto');

const GRPC_ADDR = __ENV.SERVER_HOST || '127.0.0.1:8081';
export default function () {
    client.connect(GRPC_ADDR, {
        plaintext: true
    });
    const data = {
        app_id: 'test',
        id: 'test',
    };

    let response = client.invoke('auth.v1.AuthService/Authenticate', data);
    check(response, {
        'status is OK': (r) => r && r.status === StatusOK,
    });
    response = client.invoke('auth.v1.AuthService/RefreshToken', {"refresh_token": response["message"]["refreshToken"]});
    check(response, {
        'status is OK': (r) => r && r.status === StatusOK,
    });

    response = client.invoke('auth.v1.AuthService/ValidateToken', {"access_token": response["message"]["accessToken"]});
    check(response, {
        'status is OK': (r) => r && r.status === StatusOK,
    });

    response = client.invoke('auth.v1.AuthService/AddBlocked', {
        "uid": "test",
        "is_block": true,
        "duration": 100
    });
    check(response, {
        'status is OK': (r) => r && r.status === StatusOK,
    });


    client.close();
    sleep(1);
}


