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
                v-model="userData.username"
                label="Username"
                type="text"
                append-icon="mdi-account"
                required
                dark
              ></v-text-field>
              <v-text-field
                v-model="userData.password"
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
    <Snackbar :duration="3000" />
  </v-container>
</template>
<script>
import AUTH_QUERY from "@/gql/queries/Authenticate.gql";
import LOGIN_MUTATION from "@/gql/mutations/Login.gql";
import UserModel from "@/models/user.model";
import Snackbar from "@/components/Snackbar.vue";
import { EventBus } from "@/bus";
import formMixin from "@/mixins/form.mixin";

export default {
  name: "Login",

  components: {
    Snackbar,
  },

  mixins: [formMixin],

  data: () => ({
    userData: new UserModel(),
    colorTheme: "#4F4F68",
  }),

  methods: {
    async login() {
      try {
        let result = await this.$apollo.mutate({
          mutation: LOGIN_MUTATION,
          variables: {
            username: this.userData.username,
            password: this.userData.password,
          },
          update: (store, { data: { login } }) => {
            if (login.user_id) {
              let data = store.readQuery({
                query: AUTH_QUERY,
              });
              data.me = login;
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

  mounted() {
    let params = this.$route.params;
    if (params.snackbarMessage) {
      EventBus.$emit("onShowSnackbar", params.snackbarMessage);
    }
  },
};
</script>