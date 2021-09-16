<template>
  <div>
    <MainNavbar />
    <router-view></router-view>
    <Snackbar :duration="3000" />
  </div>
</template>
<script>
import gql from "graphql-tag";
import Snackbar from "@/components/Snackbar";
import MainNavbar from "@/components/Navbar.vue";
import { EventBus } from "@/bus";

export default {
  name: "Main",

  components: {
    MainNavbar,
    Snackbar,
  },

  data: () => ({}),

  methods: {
    async authenticate() {},
    async logout(isForceLogout = false) {
      let result = await this.$apollo.mutate({
        mutation: gql`
          mutation {
            logout
          }
        `,
      });

      if (result !== null) {
        let logoutMessage = isForceLogout
          ? "Session has expired. Please re-login"
          : "Successfully logged out";
        this.redirectLogin(logoutMessage);
      } else {
        EventBus.$emit(
          "onShowSnackbar",
          "An unexpected error happened. Please try again"
        );
      }
    },
    redirectLogin(snackbarMessage) {
      this.$router.push({
        name: "Login",
        params: {
          snackbarMessage: snackbarMessage,
        },
      });
    },
  },

  mounted() {
    EventBus.$on("onLogout", (isForceLogout) => {
      this.logout(isForceLogout);
    });
    EventBus.$on("onAuthenticate", () => {
      this.authenticate();
    });
  },
};
</script>