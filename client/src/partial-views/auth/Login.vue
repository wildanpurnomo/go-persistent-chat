<template>
  <v-container fill-height>
    <v-row>
      <v-col cols="12">
        <v-row justify="center">
          <v-card class="pa-10 rounded-lg" width="500" :color="colorTheme">
            <v-row justify="center">
              <div class="mb-15 text-h3 white--text">Login</div>
            </v-row>
            <v-form ref="loginForm">
              <v-text-field
                v-model="user.username"
                label="Username"
                type="text"
                append-icon="mdi-account"
                required
                dark
              ></v-text-field>
              <v-text-field
                v-model="user.password"
                label="Password"
                :append-icon="isPasswordShown ? 'mdi-eye' : 'mdi-eye-off'"
                :type="isPasswordShown ? 'text' : 'password'"
                required
                dark
                ref="passwordField"
                @click:append="isPasswordShown = !isPasswordShown"
              ></v-text-field>

              <div class="text-center ma-8">
                <v-btn
                  width="70%"
                  type="submit"
                  class="white--text red accent-2 rounded-xl"
                  :disabled="isFormLoading"
                  :loading="isFormLoading"
                  @click.prevent="login"
                  >Masuk</v-btn
                >
              </div>
              <div class="text-center white--text">
                Belum memiliki akun?
                <router-link :to="{ name: 'Register' }"
                  ><b>Daftar</b></router-link
                >
              </div>
              <div class="red--text mt-8" :hidden="errorMessage.length === 0">
                {{ errorMessage }}
              </div>
            </v-form>
          </v-card>
        </v-row>
      </v-col>
    </v-row>
  </v-container>
</template>
<script>
import UserModel from "@/models/user.model";

export default {
  name: "Login",
  data: () => ({
    user: new UserModel(),
    isPasswordShown: false,
    errorMessage: "",
    colorTheme: "#4F4F68",
  }),
  methods: {},
  mounted() {
    console.log(this.$refs);
  },
};
</script>