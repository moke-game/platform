import {Client, StatusOK} from 'k6/net/grpc';
import {check, sleep} from 'k6';
import {crypto} from 'k6/experimental/webcrypto';


const client = new Client();
client.load(['../../api/knapsack'], 'knapsacks.proto');
client.load(['../../api/auth'], 'auth.proto');

const GRPC_ADDR = __ENV.SERVER_HOST || '127.0.0.1:8081';
const uid = crypto.randomUUID();
export default function () {
    client.connect(GRPC_ADDR, {
        plaintext: true
    });

    addItem();
    getItemById();
    removeItem();
    getKnapsack()
    client.close();
    sleep(1);
}

function addItem() {
    const data = {
        uid: uid,
        items: {1: {id: 1, type: 1, num: 1, expire: 0}}
    };
    let response = client.invoke('knapsack.pb.KnapsackPrivateService/AddItem', data);
    check(response, {
        'status is OK': (r) => r && r.status === StatusOK,
    });
    console.log(response);
}

function removeItem() {
    const data = {
        uid: uid,
        items: {1: {id: 1, type: 1, num: 1, expire: 0}}
    };
    let response = client.invoke('knapsack.pb.KnapsackPrivateService/RemoveItem', data);
    check(response, {
        'status is OK': (r) => r && r.status === StatusOK,
    });
    console.log(response);
}

function getItemById() {
    const data = {
        uid: uid,
        item_id: 1
    };
    let response = client.invoke('knapsack.pb.KnapsackPrivateService/GetItemById', data);
    check(response, {
        'status is OK': (r) => r && r.status === StatusOK,
    });
    console.log(response);

}

function getKnapsack() {
    const data = {
        uid:uid
    };
    let response = client.invoke('knapsack.pb.KnapsackPrivateService/GetKnapsack', data);
    check(response, {
        'status is OK': (r) => r && r.status === StatusOK,
    });
    console.log(response);
}