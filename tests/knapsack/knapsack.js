import {Client, StatusOK} from 'k6/net/grpc';
import {check, sleep} from 'k6';
import {makeParams} from "../common/common.js";
import {crypto} from 'k6/experimental/webcrypto';


const client = new Client();
client.load(['../../api/knapsack'], 'knapsacks.proto');
client.load(['../../api/auth'], 'auth.proto');

const GRPC_ADDR = __ENV.SERVER_HOST || '127.0.0.1:8081';

export default function () {
    client.connect(GRPC_ADDR, {
        plaintext: true
    });

    const token = login();
    addKnapsackItems(token)
    addAndRemoveKnapsackItems(token)
    removeKnapsackItems(token)
    getKnapsack(token)
    addAtomicKnapsackItems(token)
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
    return response["message"]["accessToken"];
}


function getKnapsack(token) {
    const params = makeParams(token);
    const request = {};
    let res = client.invoke('knapsack.pb.KnapsackService/GetKnapsack', request, params);
    check(res, {
        'status is OK': (r) => r && r.status === StatusOK,
    });
    console.log("get knapsack", res)
}

function removeKnapsackItems(token) {
    const params = makeParams(token);
    const request = {
        "items": {
            "1": {
                "id": 1,
                "type": 1,
                "num": 1,
                "expire": 0
            },
            "2": {
                "id": 2,
                "type": 2,
                "num": 1,
                "expire": 0
            },
            "3": {
                "id": 3,
                "type": 3,
                "num": 1,
                "expire": 0
            }
        }
    };
    let res = client.invoke('knapsack.pb.KnapsackService/RemoveItem', request, params);
    check(res, {
        'status is OK': (r) => r && r.status === StatusOK,
    });
    console.log("remove items", res)
}

function addKnapsackItems(token) {
    const params = makeParams(token);

    const request = {
        "items": {
            "1": {
                "id": 1,
                "type": 1,
                "num": 1,
                "expire": 0
            },
            "2": {
                "id": 2,
                "type": 2,
                "num": 1,
                "expire": 0
            },
            "3": {
                "id": 3,
                "type": 3,
                "num": 1,
                "expire": 0
            }
        }
    };
    let res = client.invoke('knapsack.pb.KnapsackService/AddItem', request, params);
    check(res, {
        'status is OK': (r) => r && r.status === StatusOK,
    });
    console.log("add items", res)
}

function addAndRemoveKnapsackItems(token) {
    const params = makeParams(token);

    const request = {
        "remove_items": {
            "1": {
                "id": 1,
                "type": 1,
                "num": 1,
                "expire": 0
            },
            "2": {
                "id": 2,
                "type": 2,
                "num": 1,
                "expire": 0
            },
            "3": {
                "id": 3,
                "type": 3,
                "num": 1,
                "expire": 0
            }
        },
        "add_items": {
            "1": {
                "id": 1,
                "type": 1,
                "num": 1,
                "expire": 0
            },
            "2": {
                "id": 2,
                "type": 2,
                "num": 1,
                "expire": 0
            },
            "3": {
                "id": 3,
                "type": 3,
                "num": 1,
                "expire": 0
            }
        }
    };
    let res = client.invoke('knapsack.pb.KnapsackService/RemoveThenAddItem', request, params);
    check(res, {
        'status is OK': (r) => r && r.status === StatusOK,
    });
    console.log("add and remove items", res)
}


function addAtomicKnapsackItems(token) {
    const params = makeParams(token);

    const request = {
        "add_items": {
            "1": {
                "id": 1,
                "type": 1,
                "num": 1,
                "expire": 0
            },
            "2": {
                "id": 2,
                "type": 2,
                "num": 1,
                "expire": 0
            },
            "3": {
                "id": 3,
                "type": 3,
                "num": 1,
                "expire": 0
            }
        },
        "sub_items": {
            "1": {
                "id": 1,
                "type": 1,
                "num": 1,
                "expire": 0
            },
            "2": {
                "id": 2,
                "type": 2,
                "num": 1,
                "expire": 0
            },
            "3": {
                "id": 3,
                "type": 3,
                "num": 1,
                "expire": 0
            }
        },
        "atomic_id": "test",
        "events": [
            {
                "event_id": "test",
                "action": 1,
                "num": 1
            }
        ]
    };
    let res = client.invoke('knapsack.pb.KnapsackService/AtomicUpdateItem', request, params);
    check(res, {
        'status is OK': (r) => r && r.status === StatusOK,
    });
    console.log("add atomic items", res)
}
