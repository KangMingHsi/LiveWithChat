<template>
  <v-row name="row">
    <v-menu
      offset-y
      transition="slide-y-transition"
      bottom
    >
      <template v-slot:activator="{ on, attrs }">
        <v-avatar
          color="black"
          v-bind="attrs"
          v-on="on"
        >
          <v-icon dark>
            mdi-account-circle
          </v-icon>
        </v-avatar>
      </template>
      <v-list
        nav
        dense
        subheader
      >
        <v-subheader>Account</v-subheader>
        <v-list-item
          v-for="item in accountItems"
          :key="item.title"
          link
        >
          <v-list-item-content @click="item.action">
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-divider></v-divider>
        <v-list-item
          v-for="item in videoItems"
          :key="item.title"
          link
        >
          <v-list-item-content @click="item.action">
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-menu>
  </v-row>
</template>

<script>

  export default {
    name: 'AccountMenu',
    data() {
      return {
        accountItems: [
          {
            title: 'Log out',
            action: this.logout,
          },
        ],
        videoItems: [
          {
            title: 'My Channel',
            action: () => {
              this.$router.push('/channel/' + this.$store.state.auth.id)
            },
          },
        ],
      }
    },
    methods: {
      logout() {
        this.$api.auth.logout().finally(() => {
          this.$store.dispatch('auth/setAuth', {
            "token": "",
            "isLogin": false,
            "id": "",
          })
        })
      },
    }
  }
</script>