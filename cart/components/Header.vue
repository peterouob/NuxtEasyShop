<template>
  <v-app-bar>
    <v-toolbar-title @click="$router.push('/')">
      shopping website
    </v-toolbar-title>
    <v-spacer></v-spacer>
    <v-btn @click="$router.push('/Cart')" class="text-none" stacked="">
      <v-badge :content="cartStore.productsTotal" color="error">
        <v-icon>
          mdi-cart-outline
        </v-icon>
      </v-badge>
    </v-btn>
<!--    register-->
    <v-btn
        @click="dialogLogin = !dialog"
        stacked
        class="mb-5"
    >
      <v-icon>
        mdi-login
      </v-icon>
      <v-row justify="center">
        <v-dialog
            v-model="dialog"
            persistent
            width="1024"
        >
          <v-card>
            <v-card-title>
              <span class="text-h5">User Register</span>
            </v-card-title>
            <v-card-text>
              <v-container>
                <v-row>
                  <v-col
                      cols="12"
                      sm="6"
                      md="4"
                  >
                    <v-text-field
                        label="Legal first name*"
                        v-model="user.firstname"
                        required
                    ></v-text-field>
                  </v-col>
                  <v-col
                      cols="12"
                      sm="6"
                      md="4"
                  >
                    <v-text-field
                        label="Legal last name*"
                        v-model="user.lastname"
                        hint="example of persistent helper text"
                        persistent-hint
                        required
                    ></v-text-field>
                  </v-col>
                  <v-col cols="12">
                    <v-text-field
                        label="Phone*"
                        v-model="user.phone"
                        required
                    ></v-text-field>
                  </v-col>
                  <v-col cols="12">
                    <v-text-field
                        label="Email*"
                        v-model="user.email"
                        required
                    ></v-text-field>
                  </v-col>
                  <v-col cols="12">
                    <v-text-field
                        label="username*"
                        v-model="user.username"
                        required
                    ></v-text-field>
                  </v-col>
                  <v-col cols="12">
                    <v-text-field
                        label="Password*"
                        v-model="user.password"
                        type="password"
                        required
                    ></v-text-field>
                  </v-col>

                </v-row>
              </v-container>
              <v-btn
                  color="blue-darken-1"
                  variant="text"
                  @click="dialogLogin = true;dialog = false"
              >
                I have an account
              </v-btn>
            </v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn
                  color="blue-darken-1"
                  variant="text"
                  @click="dialog = false"
              >
                Close
              </v-btn>
              <v-btn
                  color="blue-darken-1"
                  variant="text"
                  @click="register"
              >
                Register
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </v-row>
<!--      login-->
        <v-row justify="center">
          <v-dialog
              v-model="dialogLogin"
              persistent
              width="1024"
          >
            <v-card>
              <v-card-title>
                <span class="text-h5">User Login</span>
              </v-card-title>
              <v-card-text>
                <v-container>
                  <v-row>
                    <v-col cols="12">
                      <v-text-field
                          name="username"
                          v-model="user.username"
                          label="Username*"
                          required
                      ></v-text-field>
                    </v-col>
                    <v-col cols="12">
                      <v-text-field
                          name="password"
                          v-model="user.password"
                          label="Password*"
                          type="password"
                          required
                      ></v-text-field>
                    </v-col>
                  </v-row>
                </v-container>
                <v-btn
                    color="blue-darken-1"
                    variant="text"
                    @click="dialogLogin = false;dialog = true"

                >
                  I don't have an account
                </v-btn>
              </v-card-text>
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn
                    color="blue-darken-1"
                    variant="text"
                    @click="dialogLogin = false"
                >
                  Close
                </v-btn>
                <v-btn
                    color="blue-darken-1"
                    variant="text"
                    @click="login"
                >
                  Login
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-dialog>
        </v-row>
    </v-btn>
    <v-btn
    :prepend-icon="
    useCartStore().theme === 'light' ? 'mdi-weather-sunny' : 'mdi-weather-night'
      "
    @click="cartStore.toggleTheme()"
    >
      Toggle Theme
    </v-btn>
  </v-app-bar>
</template>

<script setup>
import {useCartStore,useUserTestStore} from "~/stores/index.js";
import {useAuthStore} from "~/stores/auth.store.js";

const cartStore = useCartStore()
const userStore = useUserTestStore()
const authStore = useAuthStore()

let user = reactive({
  firstname:"",
  lastname:"",
  phone:"",
  email:"",
  username:"",
  password:""
})

let dialog = ref(false)
let dialogLogin = ref(false)

function login(){
  authStore.login(user.username,user.password)
      .catch(error=>console.log(error))
}
function register(){
  authStore.register(user.firstname,user.lastname,user.phone,user.email,user.username,user.password)
      .catch(error=>console.log(error))

}

</script>

<style lang="scss" scoped>
</style>