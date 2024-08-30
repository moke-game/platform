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