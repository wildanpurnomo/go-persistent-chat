<template>
  <nav>
    <v-toolbar :color="colorTheme" dark>
      <v-toolbar-title class="hidden-md-and-up pl-1">{{
        pageTitle
      }}</v-toolbar-title>
      <v-toolbar-title
        class="hidden-sm-and-down fontLobsterTwo pl-3"
        @click="$router.push({ name: `ChatList` })"
        >Wappita</v-toolbar-title
      >
      <v-spacer></v-spacer>
      <v-btn @click="logout" text class="pr-1">
        <span class="hidden-sm-and-down">Logout</span>
        <v-icon>mdi-logout-variant</v-icon>
      </v-btn>
    </v-toolbar>
  </nav>
</template>
<script>
import { EventBus } from "@/bus";

export default {
  name: "Navbar",

  data: () => ({
    isDrawer: false,
    colorTheme: "#394867",
    pageTitle: "",
  }),

  methods: {
    logout() {
      EventBus.$emit("onLogout", false);
    },
  },
};
</script>
<style scoped>
.fontLobsterTwo {
  font-size: 25px;
  color: white;
  cursor: pointer;
}
</style>