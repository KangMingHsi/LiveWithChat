export default {
    namespaced: true,
    state: {
        token: "",
        isLogin: false,
        id: "",
    },
    mutations: {
        SET_AUTH(state, options) {
            state.id = options.id
            state.token = options.token
            state.isLogin = options.isLogin
        },
        SET_TOKEN(state, options) {
            state.token = options.token
        }
    },
    actions: {
        setAuth(context, options) {
            context.commit('SET_AUTH', {
                token: options.token,
                isLogin: options.isLogin,
                id: options.id,
            })
        },
        setToken(context, options) {
            context.commit('SET_AUTH', {
                token: options.token,
            })
        }
    }
}