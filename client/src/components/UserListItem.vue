<template>
  <v-container>
    <v-row>
      <v-col cols="2">
        <v-avatar size="48" color="primary">
          <span>{{ initial }}</span>
        </v-avatar>
      </v-col>
      <v-col class="d-flex align-center">
        <span class="text-md-h6 text-sm-subtitle pl-3">{{ username }}</span>
      </v-col>
      <v-col class="d-flex align-center justify-end">
        <v-btn
          class="rounded-xl"
          dark
          depressed
          :color="buttonColor"
          @click="onChatButtonClicked"
        >
          <v-icon class="hidden-md-and-up">mdi-account-multiple-plus</v-icon>
          <span class="hidden-sm-and-down px-5">Chat</span>
        </v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "UserListItem",

  props: {
    userId: Number,
    username: String,
  },

  data: () => ({
    buttonColor: "#394867",
  }),

  computed: {
    initial() {
      let split = this.username.split(" ");
      let initial = "";
      split.forEach((word) => {
        initial += word[0];
      });
      return initial;
    },
  },

  methods: {
    onChatButtonClicked() {
      this.$emit("on-chat-button-click", this.userId);
    },
  },
};
</script>