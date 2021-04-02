import { Message } from 'element-ui'
import router from '../router/index'

export const tip = msg => {
    Message.error(msg)
}

export const toLogin = () => {
    router.replace({
        name: 'Login',
        query: {
            redirect: router.currentRoute.fullPath
        }
    })
}

export const to403Page = () => {
    router.replace({
        name: '403',
    })
}