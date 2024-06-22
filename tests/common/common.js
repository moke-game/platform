export const makeParams = (token) => {
    return {
        metadata: {
            'x-my-header': 'k6test',
            'x-my-header-bin': new Uint8Array([1, 2, 3]),
            'authorization': "bearer " + token,
        },
        tags: {k6test: 'yes'},
    };
}


// export const login = () =>{
//     const data = {
//         app_id: 'test',
//         id: crypto.randomUUID(),
//         auth: 1,
//         data: 'test'
//     };
//     let response = client.invoke('auth.pb.AuthService/Authenticate', data);
//     check(response, {
//         'status is OK': (r) => r && r.status === StatusOK,
//     });
//     console.log(response)
//     return response["message"]["accessToken"];
// }

export const login = (client) => {
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