<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <v-img src="@/assets/emptyState.png" max-width="500"></v-img>
      </v-col>
      <v-col cols="12">
        <div class="text-caption">{{ caption }}</div>
      </v-col>
      <v-col cols="12">
        <v-btn
          class="rounded-xl"
          :color="buttonColor"
          outlined
          @click="onButtonClicked"
          v-if="buttonText"
          >{{ buttonText }}</v-btn
        >
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "EmptyState",

  props: {
    caption: String,
    buttonText: String,
    onButtonClicked: Function,
  },

  data: () => ({
    buttonColor: "#394867",
  }),
};
</script>