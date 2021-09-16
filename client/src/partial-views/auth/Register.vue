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
                v-model="userData.username"
                :rules="this.usernameRules"
                label="Username"
                type="text"
                required
                dark
              ></v-text-field>
              <v-text-field
                v-model="userData.password"
                :rules="this.passwordRules"
                label="Password"
                :append-icon="isPasswordShown ? 'mdi-eye' : 'mdi-eye-off'"
                :type="isPasswordShown ? 'text' : 'password'"
                @click:append="isPasswordShown = !isPasswordShown"
                required
                dark
              ></v-text-field>
              <v-text-field
                v-model="userData.passwordConfirmation"
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
                  @click.prevent="initiateRegister"
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
    <Snackbar :duration="3000" />
  </v-container>
</template>
<script>
import AUTH_QUERY from "@/gql/queries/Authenticate.gql";
import REGISTER_MUTATION from "@/gql/mutations/Register.gql";
import UserModel from "@/models/user.model";
import formMixin from "@/mixins/form.mixin";
import Snackbar from "@/components/Snackbar.vue";
import { EventBus } from "@/bus.js";

export default {
  name: "Register",

  components: {
    Snackbar,
  },

  mixins: [formMixin],

  data: () => ({
    userData: new UserModel(),
    passwordConfirmation: "",
    colorTheme: "#4F4F68",
  }),

  computed: {
    passwordConfirmationRules() {
      return [
        (v) =>
          (!!v && v) === this.userData.password ||
          "Masukkan password yang sama.",
      ];
    },
  },

  methods: {
    async register() {
      try {
        let result = await this.$apollo.mutate({
          mutation: REGISTER_MUTATION,
          variables: {
            username: this.userData.username,
            password: this.userData.password,
          },
          update: (store, { data: { register } }) => {
            if (register.user_id) {
              let data = store.readQuery({
                query: AUTH_QUERY,
              });
              data.me = register;
              store.writeQuery({
                query: AUTH_QUERY,
                data,
              });
            }
          },
          optimisticResponse: {
            __typename: "Mutation",
            login: {
              __typename: "UserType",
              user_id: -1,
              username: "",
              created_at: "",
              updated_at: "",
            },
          },
        });

        if (result !== null) {
          EventBus.$emit("onAuthSuccess", true);
        } else {
          EventBus.$emit(
            "onShowSnackbar",
            "An unexpected error happened. Please try again"
          );
        }
      } catch (error) {
        EventBus.$emit("onShowSnackbar", "Wrong credentials or user not found");
      }
    },
  },
};
</script>