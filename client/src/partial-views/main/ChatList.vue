<template>
  <v-container>
    <h3>Chat List</h3>
    <EmptyState
      caption="Halo niece and nephew its uncle roger"
      :on-button-clicked="toSearchUsersPage"
      button-text="Search users"
    />
  </v-container>
</template>

<script>
import EmptyState from "@/components/EmptyState.vue";

export default {
  name: "ChatList",

  components: { EmptyState },

  data: () => ({}),

  methods: {
    toSearchUsersPage() {
      this.$router.push({ name: "SearchUsers" });
    },
  },
};
</script>