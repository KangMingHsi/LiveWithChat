export default {
    namespaced: true,
    state: {
        title: "",
        description: "",
        id: "",
    },
    mutations: {
        SET_WATCH(state, options) {
            state.id = options.id
            state.title = options.title
            state.description = options.description
        },
    },
    actions: {
        setWatch(context, options) {
            context.commit('SET_WATCH', {
                title: options.title,
                description: options.description,
                id: options.id,
            })
        },
    }
}