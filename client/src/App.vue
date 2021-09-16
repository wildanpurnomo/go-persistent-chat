<template>
  <v-app>
    <v-main class="blue-grey lighten-5">
      <router-view></router-view>
    </v-main>
  </v-app>
</template>

<script>
export default {
  name: "App",

  data: () => ({
    //
  }),
};
</script>

<style>
@import url("https://fonts.googleapis.com/css2?family=Lobster+Two:wght@700&display=swap");
.fontLobsterTwo {
  font-family: Lobster Two;
}
</style>
