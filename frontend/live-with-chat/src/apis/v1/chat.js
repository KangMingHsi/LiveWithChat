import req from '../https.js'

const chat = {
    list(params) {
        return req('get', '/v1/chat/messages', params)
    },
    send(params) {
        return req('post', '/v1/chat/messages', params)
    },
    update(params) {
        return req('patch', '/v1/chat/messages', params)
    },
    delete(params) {
        return req('delete', '/v1/chat/messages', params)
    },
}

export default chat