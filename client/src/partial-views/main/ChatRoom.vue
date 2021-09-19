<template>
  <v-container>
    <h1>{{ chatRoomId }}</h1>
    <v-row>
      <v-col cols="12">
        <v-card class="pa-10 rounded-lg"> </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="10"></v-col>
      <v-col cols="2"></v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "ChatRoom",

  data: () => ({
    chatRoomId: -1,
  }),

  mounted() {
    this.chatRoomId = this.$route.params.chatRoomId;
  },
};
</script>