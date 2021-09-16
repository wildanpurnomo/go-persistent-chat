<template>
  <router-view></router-view>
</template>
<script>
import { EventBus } from "@/bus";
import AUTH_QUERY from "@/gql/queries/Authenticate.gql";
import UserModel from "@/models/user.model";

export default {
  name: "Auth",

  data: () => ({
    me: new UserModel(),
  }),

  apollo: {
    me: {
      query: AUTH_QUERY,
      update(data) {
        if (data.me) this.logUserIn();
      },
    },
  },

  methods: {
    logUserIn() {
      this.$router.push({ name: "ChatList" });
    },
  },

  mounted() {
    EventBus.$on("onAuthSuccess", () => {
      this.logUserIn();
    });
  },
};
</script>