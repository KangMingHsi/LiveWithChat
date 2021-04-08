export default {
    namespaced: true,
    state: {
        token: "",
        isLogin: false,
    },
    mutations: {
        SET_AUTH(state, options) {
            state.token = options.token
            state.isLogin = options.isLogin
        },
    },
    actions: {
        setAuth(context, options) {
            context.commit('SET_AUTH', {
                token: options.token,
                isLogin: options.isLogin,
            })
        },
    }
}