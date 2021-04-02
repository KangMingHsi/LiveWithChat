import axios from 'axios'
import router from '../router/index'
import { tip, toLogin, to403Page } from './utils'

const errorHandle = (status, msg) => {
    switch(status) {
        case 400:
            tip(msg)
            break
        case 401:
            if (router.currentRoute != 'Login') {
                this.$store.dispatch('auth/setAuth', {
                    "token": '',
                    "toLogin": false
                })

                tip(msg)
                setTimeout(() => {
                    toLogin()
                }, 1000)
            }
            break
        case 403:
            to403Page()
            break
        case 404:
            tip(msg)
            break
        default:
            console.log('no handle to this error: ' + msg)
    }
}

var instance = axios.create({
    baseURL: '/api'
})

instance.interceptors.request.use((config) => {
    const token = this.$store.state.auth.token
    token && (config.headers.Authorization = 'Bearer ' + token)
    return config
}, (error) => {
    return Promise.reject(error)
})

instance.interceptors.response.use((response) => {
    return response
}, (error) => {
    const { response } = error
    if (response) {
        errorHandle(response.status, response.data.error)
        return Promise.reject(error)
    } else {
        if(!window.navigator.onLine) {
            tip('please connect to network and refresh the page')
        } else {
            return Promise.reject(error)
        }
    }
})

export default function (method, url, data=null) {
    method = method.toLowerCase()
    if (method == 'get') {
        return instance.get(url, {params: data})
    } else if (method == 'post') {
        return instance.post(url, data)
    } else if (method == 'put') {
        return instance.put(url, data)
    } else if (method == 'delete') {
        return instance.delete(url, {params: data})
    } else {
        console.error('unknown method: ' + method)
        return false
    }
}