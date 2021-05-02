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
    <div class="navbar-end">
      <LoginDialog
        v-if="!isLogin"
      />
      <AccountMenu
        v-else
      />
    </div>
  </div>
</nav>
</template>
<script>
  import LoginDialog from '../LoginDialog'
  import AccountMenu from '../AccountMenu'

  export default {
    name: 'Nav',
    components: {
      LoginDialog,
      AccountMenu,
    },
    data() {
      return {
      }
    },
    computed: {
      isLogin() {
        return this.$store.state.auth.isLogin
      }
    },
    watch: {
    },
    methods: {
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