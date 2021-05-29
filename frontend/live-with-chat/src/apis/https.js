import axios from 'axios'
import api from '../apis/index'
import store from '../store/index'
import { to403Page } from './utils'

async function errorHandle(status, msg, error) {
    switch(status) {
        case 400:
            alert(msg)
            break
        case 401:
            alert(msg)
            var refresh = api.auth.refresh().then((resp) => {
                let data = resp.data
                store.dispatch('auth/setToken', {
                    "token": data.token,
                })
            }).catch(() => {
                store.dispatch('auth/setAuth', {
                    "token": "",
                    "isLogin": false,
                    "id": "",
                })
            })
            await refresh
            error.config.headers['Authorization'] = 'Bearer ' + store.state.auth.token;
            return axios.request(error.config)
        case 403:
            to403Page()
            break
        case 404:
            alert(msg)
            break
        default:
            console.log('no handle to this error: ' + msg)
    }
    return Promise.reject(error)
}

const instance = axios.create({
    baseURL: '/api'
})

instance.interceptors.request.use((config) => {
    const token = store.state.auth.token
    token && (config.headers.Authorization = 'Bearer ' + token)
    return config
}, (error) => {
    return Promise.reject(error)
})

instance.interceptors.response.use((response) => {
    return response
}, async (error) => {
    const { response } = error
    if (response) {
        return await errorHandle(response.status, response.data.error, error)
        // if (response.status == 401) {
        //     error.config.headers['Authorization'] = 'Bearer ' + store.state.auth.token;
        //     return axios.request(error.config)
        // }
        // return Promise.reject(error)
    } else {
        if(!window.navigator.onLine) {
            alert('please connect to network and refresh the page')
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
    } else if (method == 'patch') {
        return instance.patch(url, data)
    } else {
        console.error('unknown method: ' + method)
        return false
    }
}