import req from './https.js'

const auth = {
    signUp(params) {
        return req('post', '/auth/register', params)
    },
    login(params) {
        return req('post', '/auth/login', params)
    },
    logout() {
        return req('post', '/auth/logout', null)
    },
    refresh() {
        return req('post', '/auth/refresh', null)
    },
}

export default auth