import {Client, StatusOK, Stream} from 'k6/net/grpc';
import {check, sleep} from 'k6';
import {crypto} from 'k6/experimental/webcrypto';
import {makeParams} from "../common/common.js";


const client = new Client();
client.load(['../../api/mail'], 'mail.proto');
client.load(['../../api/auth'], 'auth.proto');

const GRPC_ADDR = __ENV.SERVER_HOST || '127.0.0.1:8081';
let uid = "";
export default function () {
    client.connect(GRPC_ADDR, {
        plaintext: true
    });
    const token = login();
    watch(token);
    sendMail(token);
    updateMail(token)
    client.close();
    sleep(1);
}

function login() {
    const data = {
        app_id: 'test',
        id: crypto.randomUUID(),
        auth: 1,
        data: 'test'
    };
    let response = client.invoke('auth.pb.AuthService/Authenticate', data);
    check(response, {
        'status is OK': (r) => r && r.status === StatusOK,
    });
    console.log(response)
    uid = response["message"]["id"];
    return response["message"]["accessToken"];
}

function watch(token) {
    const params = makeParams(token);
    // const data = {
    //     channel: "",
    //     language: "en"
    // };
    const stream = new Stream(client, 'mail.v1.MailService/Watch', params);
    stream.on('data', (data) => {
        console.log(data);
    });

    stream.on('error', (err) => {
        console.log('error', err);
    });

    stream.on('end', () => {
        client.close();
        console.log('stream ended');
    });
}

function sendMail(token) {
    const request = {
        platform_id: "",
        send_type: 1,
        mail: {
            id: 1,
            title: {"en": "hello"},
            body: {"en": "body"},
            date: Math.floor(Date.now() / 1000),
            from: "sender",
            template_id: 1,
            template_args: ["1","2","3"]
        }
    }
    let response = client.invoke('mail.v1.MailPrivateService/SendMail', request);
    check(response, {
        'status is OK': (r) => r && r.status === StatusOK,
    });
    console.log(response);
}

function updateMail(token) {
    const params = makeParams(token);
    const req = {
        updates: {
            1: 1
        }
    };
    let response = client.invoke('mail.v1.MailService/UpdateMail', req, params);
    check(response, {
        'status is OK': (r) => r && r.status === StatusOK,
    });
    console.log(response);
}