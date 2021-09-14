<template>
  <v-container fill-height>
    <v-row>
      <v-col cols="12">
        <v-row justify="center">
          <v-card class="pa-10 rounded-lg" width="500" :color="colorTheme">
            <v-row justify="center">
              <div class="mb-15 text-h3 white--text">Register</div>
            </v-row>
            <v-form ref="registerForm" v-model="isFormValid">
              <v-text-field
                v-model="user.username"
                :rules="this.usernameRules"
                label="Username"
                type="text"
                required
                dark
              ></v-text-field>
              <v-text-field
                v-model="user.email"
                :rules="this.emailRules"
                label="Email"
                type="email"
                required
                dark
              ></v-text-field>
              <v-text-field
                v-model="user.password"
                :rules="this.passwordRules"
                label="Password"
                :append-icon="isPasswordShown ? 'mdi-eye' : 'mdi-eye-off'"
                :type="isPasswordShown ? 'text' : 'password'"
                @click:append="isPasswordShown = !isPasswordShown"
                required
                dark
              ></v-text-field>
              <v-text-field
                v-model="user.passwordConfirmation"
                label="Konfirmasi Password"
                :append-icon="isPasswordShown ? 'mdi-eye' : 'mdi-eye-off'"
                :type="isPasswordShown ? 'text' : 'password'"
                required
                dark
                :rules="passwordConfirmationRules"
                @click:append="isPasswordShown = !isPasswordShown"
              ></v-text-field>
              <div class="text-center ma-8">
                <v-btn
                  type="submit"
                  class="white--text red accent-2 rounded-xl"
                  width="70%"
                  :disabled="!isFormValid || isFormLoading"
                  :loading="isFormLoading"
                  @click.prevent="register"
                  dark
                  >Register</v-btn
                >
              </div>
              <div class="text-center white--text">
                Sudah memiliki akun?
                <router-link :to="{ name: 'Login' }"
                  ><b>Kembali ke login</b></router-link
                >
              </div>
              <div class="h6 red--text" :hidden="errorMessage.length === 0">
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
  name: "Register",
  data: () => ({
    user: new UserModel(),
    errorMessage: "",
    passwordConfirmation: "",
    isPasswordShown: false,
    colorTheme: "#4F4F68",
  }),
  computed: {
    passwordConfirmationRules() {
      return [
        (v) =>
          (!!v && v) === this.user.password || "Masukkan password yang sama.",
      ];
    },
  },
  methods: {}
};
</script>