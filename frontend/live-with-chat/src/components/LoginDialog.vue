<template>
  <div class="text-center">
    <v-row class="container">
      <v-col class="col"> 
        <v-card
         class="signin"
         width="600"
        >
          <v-card-title>
            <h2>LOGIN</h2>
          </v-card-title>
          
          <validation-observer ref="observer">
            <v-form @submit.prevent="onSubmit">
              <v-card-text>
                <validation-provider v-slot="{ errors }" name="Email" rules="required|email">
                  <v-text-field
                    v-model="email"
                    :error-messages="errors"
                    label="Email"
                    required
                    outlined
                    filled
                    dense
                  ></v-text-field>
                </validation-provider>
                <validation-provider v-slot="{ errors }" name="Password" rules="required">
                  <v-text-field
                    v-model="password"
                    :error-messages="errors"
                    label="Password"
                    :append-icon="showPass ? 'mdi-eye' : 'mdi-eye-off'"
                    @click:append="showPass = !showPass"
                    required
                    outlined
                    dense
                    filled
                    :type="showPass ? 'text' : 'password'"
                  ></v-text-field>
                </validation-provider>
              </v-card-text>
              <v-card-actions class="text-center">
                <v-btn class="signin-btn" type="submit" rounded color="blue-grey" dark>
                  Sign in
                </v-btn>
              </v-card-actions>
            </v-form>
          </validation-observer>
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>

<script>
import { required, email } from 'vee-validate/dist/rules'
import { extend, ValidationProvider, setInteractionMode, ValidationObserver } from 'vee-validate'

setInteractionMode('eager')

extend('required', {
  ...required,
  message: '{_field_} can not be empty'
})

extend('email', {
  ...email,
  message: 'Email must be valid'
})

export default {
  name: 'LoginDialog',
  components: {
    ValidationProvider,
    ValidationObserver
  },
  data: () => ({
    email: '',
    password: null,
    showPass: false
  }),
  methods: {
    async onSubmit() {
      const valid = await this.$refs.observer.validate()
      if (valid) {
        var loginForm = new FormData()
        loginForm.append("email", this.email)
        loginForm.append("password", this.password)

        this.$api.auth.login(
          loginForm
        ).then((resp) => {
          let data = resp.data
          this.$store.dispatch('auth/setAuth', {
            "accessToken": data.AccessToken,
            "refreshToken": data.RefreshToken,
            "isLogin": true,
          })
          this.clear()
          this.$emit('closeLogin')
          // this.$router.push('/')
        }).catch((error) => {
          console.log(error)
        }) // action to login
      }
    },
    clear() {
      // you can use this method to clear login form
      this.email = ''
      this.password = null
      this.$refs.observer.reset()
    }
  }
}
</script>

<style lang="scss" scoped>
.container {
  padding: 0;
  background: rgb(255, 255, 255);
  width: 100%;
  box-shadow: 0 0 1px 1px rgba($color: #000000, $alpha: 0.1);
  box-sizing: border-box;
  .col {
    padding: 0;
    max-width: 1000px;
    box-shadow: 0 0 1px 1px rgba($color: #000000, $alpha: 0.1);
    .signin {
      padding: 0;
      box-sizing: border-box;
      background: #ffffff;
      color: rgb(5, 5, 5);
      h2 {
        text-align: center;
        margin: 30px 0;
      }
      .v-btn {
        width: 100%;
        color: #ffffff;
      }
    }
  }
}

</style>