import req from './https.js'

const auth = {
    signUp(params) {
        return req('post', '/signup', params)
    },
    login(params) {
        return req('post', '/login', params)
    }
}

export default auth