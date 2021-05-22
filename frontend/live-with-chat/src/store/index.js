import Vue from 'vue'
import Vuex from 'vuex'
import createPersistedState from "vuex-persistedstate"

import moduleAuth from './modules/auth'
import moduleWatch from './modules/watch'

Vue.use(Vuex)

export default new Vuex.Store({
    modules: {
        watch: moduleWatch,
        auth: moduleAuth,
    },
    plugins: [createPersistedState({
        storage: window.localStorage,
        reducer(val) {
            return {
                watch: val.watch,
                auth: val.auth,
            }
        }
    })]
})