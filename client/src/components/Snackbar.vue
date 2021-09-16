<template>
  <v-snackbar v-model="isSnackbarShown" :bottom="true" :timeout="duration">
    {{ message }}
    <template v-slot:action="{ attrs }">
      <v-btn dark text v-bind="attrs" @click="isSnackbarShown = false"
        >Tutup</v-btn
      >
    </template>
  </v-snackbar>
</template>
<script>
import { EventBus } from "@/bus.js";

export default {
  name: "Snackbar",
  props: {
    duration: Number,
  },
  data: () => ({
    isSnackbarShown: false,
    message: "",
  }),
  methods: {
    showSnackbar(message) {
      this.message = message;
      this.isSnackbarShown = true;
    },
  },
  mounted() {
    EventBus.$on("onShowSnackbar", (message) => {
      this.showSnackbar(message);
    });
  },
};
</script>