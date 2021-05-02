import req from '../https.js'

const stream = {
    list(params) {
        return req('get', '/v1/stream/videos', params)
    },
    upload(params) {
        return req('post', '/v1/stream/videos', params)
    },
    update(params) {
        return req('patch', '/v1/stream/videos', params)
    },
    delete(params) {
        return req('delete', '/v1/stream/videos', params)
    },
}

export default stream