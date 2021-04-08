<template>
<nav class="navbar container" role="navigation" aria-label="main navigation">
  <div class="navbar-brand">
    <a class="navbar-item" href="/">
      <strong class="is-size-4">Live With Chat</strong>
    </a>
    <a role="button" class="navbar-burger burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
    </a>
  </div>
  <div id="navbar" class="navbar-menu">
    <div class="navbar-start">
      <router-link to="/" class="navbar-item">Home</router-link>
      <router-link to="/about" class="navbar-item">About</router-link>
    </div>
    <div class="navbar-end-login"
     v-if="!isLogin"
    >
      <v-dialog
       v-model="showLogin"
       width="500"
      >
        <template v-slot:activator="{ on, attrs }">
          <div class="buttons"
           v-bind="attrs"
           v-on="on"
          >
            <a class="button is-dark">
              <strong>Sign In</strong>
            </a>
          </div>
        </template>
        <LoginDialog
          ref="loginDialog"
          @closeLogin="closeLogin"
        />
      </v-dialog>
    </div>
    <div class="navbar-end-logout"
     v-if="isLogin"
    >
      <a class="button is-dark" @click="logout">
        <strong>Sign Out</strong>
      </a>
    </div>
  </div>
</nav>
</template>
<script>
import LoginDialog from '../LoginDialog'

export default {
  name: 'Nav',
  components:{
    LoginDialog
  },
  data() {
    return {
      showLogin: false,
    }
  },
  computed: {
    isLogin() {
      return this.$store.state.auth.isLogin
    }
  },
  watch: {
    showLogin: function(val) {
      val || this.$refs.loginDialog.clear()
    },
  },
  methods: {
    closeLogin() {
      this.showLogin = false
    },
    logout() {
      this.$api.auth.logout().finally(() => {
        this.$store.dispatch('auth/setAuth', {
          "token": "",
          "isLogin": false,
        })
      })
    }
  },
}
</script>
<style lang="scss" scoped>
  nav {
    margin-top: 25px;
    margin-bottom: 30px;
    a {
      font-weight: bold;
      color: #2c3e50;
      &.router-link-exact-active {
        color: #d88d00;
      }
    }  
  } 
</style>