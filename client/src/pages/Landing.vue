<template>
  <v-container fill-height>
    <v-row class="d-flex align-center">
      <v-col class="pb-10">
        <v-img src="@/assets/landing.png" max-width="600"></v-img>
      </v-col>
      <v-col>
        <v-card-title class="font py-5">
          Wappita
        </v-card-title>
        <v-card-text class="ma-2">
          Bagikan dan temukan pengetahuan baru. Perluas pengetahuan anda dengan membaca berbagai konten yang menarik.
        </v-card-text>
        <v-row class="pl-8 pt-6">
          <div class="mr-8">
            <v-btn
              class="white--text red accent-2 rounded-xl"
              outlined
              width="100"
              :to="{ name: 'Login' }"
            >Masuk</v-btn>
          </div>
          <div>
            <v-btn
              class="rounded-xl"
              :color="colorTheme"
              outlined
              :to="{ name: 'Register' }"
            >Register</v-btn>
          </div>
        </v-row>
      </v-col>
    </v-row>
  </v-container>
</template>
<script>
export default {
  name: "Landing",

  components: {},

  data: () => ({
    colorTheme: "#394867",
  }),
};
</script>
<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Lobster+Two:wght@700&display=swap");
.font {
  font-family: Lobster Two;
  font-size: 50px;
  color: #394867
}
</style>