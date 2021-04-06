<template>
    <div class="text-center">
        <v-row class="container">
            <v-card
             class="signup"
             width="600"
            >
                <v-card-title>
                    <h2>REGISTER</h2>
                </v-card-title>

                <validation-observer ref="observer">
                    <v-form @submit.prevent="onSubmit">
                        <v-card-text>
                            <v-col>
                                <validation-provider v-slot="{ errors }" name="Nickname" rules="required">
                                    <v-text-field
                                    v-model="nickname"
                                    :error-messages="errors"
                                    label="Nickname"
                                    required
                                    outlined
                                    filled
                                    dense
                                    ></v-text-field>
                                </validation-provider>

                                <v-radio-group
                                v-model="gender"
                                mandatory
                                label="Gender: "
                                dense
                                row
                                >
                                    <v-radio
                                    label="Male"
                                    value="male"
                                    ></v-radio>
                                    <v-radio
                                    label="Female"
                                    value="female"
                                    ></v-radio>
                                </v-radio-group>
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
                                <validation-provider v-slot="{ errors }" name="Password" rules="required" vid="password">
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
                                <validation-provider v-slot="{ errors }" name="rePassword" rules="required|confirmed:password">
                                    <v-text-field
                                        v-model="rePassword"
                                        :error-messages="errors"
                                        label="Retype Password"
                                        required
                                        outlined
                                        dense
                                        filled
                                        type="password"
                                    ></v-text-field>
                                </validation-provider>
                            </v-col>
                        </v-card-text>
                        <v-card-actions class="text-center">
                            <v-btn class="signup-btn" type="submit" rounded color="blue-grey" dark>
                                Sign up
                            </v-btn>
                    </v-card-actions>
                    </v-form>
                </validation-observer>
            </v-card>
        </v-row>
    </div>
</template>

<script>
import { required, email, confirmed } from 'vee-validate/dist/rules'
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

extend('confirmed', {
  ...confirmed,
  message: 'Password must match'
})

export default {
    name: 'RegisterDialog',
    components: {
        ValidationProvider,
        ValidationObserver
    },

    data () {
        return {
            nickname: "",
            gender: "",
            email: "",
            password: "",
            rePassword: "",
            showPass: false
        }
    },

    methods: {
        async onSubmit() {
            const valid = await this.$refs.observer.validate()
            if (valid) {
                var loginForm = new FormData()
                loginForm.append("email", this.email)
                loginForm.append("password", this.password)
                loginForm.append("nickname", this.nickname)
                loginForm.append("gender", this.gender)

                this.$api.auth.signUp(
                    loginForm
                ).then(() => {
                    this.clear()
                    this.$emit('closeRegister')
                })
            }
        },
        clear() {
            // you can use this method to clear login form
            this.nickname = ""
            this.gender = ""
            this.email = ""
            this.password = ""
            this.rePassword = ""
            this.showPass = false
            this.$refs.observer.reset()
        }
    },
}
</script>

<style lang="scss" scoped>

.container {
  padding: 0;
  background: rgb(255, 255, 255);
  width: 100%;
  box-shadow: 0 0 1px 1px rgba($color: #000000, $alpha: 0.1);
  box-sizing: border-box;
}

.v-btn {
    width: 100%;
    color: #ffffff;
}

</style>