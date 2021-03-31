export default {
    namespaced: true,
    state: {
        accessToken: "",
        refreshToken: "",
        isLogin: false,
    },
    mutations: {
        SET_AUTH(state, options) {
            state.accessToken = options.accessToken
            state.refreshToken = options.refreshToken
            state.isLogin = options.isLogin
        }
    },
    actions: {
        setAuth(context, options) {
            context.commit('SET_AUTH',{
                accessToken: options.accessToken,
                refreshToken: options.refreshToken,
                isLogin: options.isLogin,
            })
        }
    }
}