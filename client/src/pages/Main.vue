<template>
  <div>
    <MainNavbar :username="me ? me.username : ''" />
    <router-view></router-view>
    <Snackbar :duration="3000" />
  </div>
</template>
<script>
import AUTH_QUERY from "@/gql/queries/Authenticate.gql";
import LOGOUT_MUTATION from "@/gql/mutations/Logout.gql";
import Snackbar from "@/components/Snackbar";
import MainNavbar from "@/components/Navbar.vue";
import { EventBus } from "@/bus";
import UserModel from "@/models/user.model";

export default {
  name: "Main",

  components: {
    MainNavbar,
    Snackbar,
  },

  data: () => ({
    me: new UserModel(),
  }),

  apollo: {
    me: {
      query: AUTH_QUERY,
      error() {
        this.logout(true);
      },
    },
  },

  methods: {
    async logout(isForceLogout = false) {
      let result = await this.$apollo.mutate({
        mutation: LOGOUT_MUTATION,
        update: (store, { data: { logout } }) => {
          if (logout) {
            let data = store.readQuery({
              query: AUTH_QUERY,
            });
            data.me = null;
            store.writeQuery({
              query: AUTH_QUERY,
              data,
            });
          }
        },
      });

      if (result) {
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
  },
};
</script>