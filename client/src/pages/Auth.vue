<template>
  <router-view></router-view>
</template>
<script>
import { EventBus } from "@/bus";
import AUTH_QUERY from "@/gql/queries/Authenticate.gql";

export default {
  name: "Auth",

  data: () => ({
    me: {},
  }),

  apollo: {
    me: {
      query: AUTH_QUERY,
      update() {
        this.logUserIn();
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