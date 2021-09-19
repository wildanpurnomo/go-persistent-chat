<template>
  <v-container>
    <v-row class="d-flex justify-center">
      <v-col cols="10" class="pl-1">
        <v-text-field
          outlined
          dense
          label="Search by username"
          append-icon="mdi-magnify"
          rounded
          @input="onSearchQueryChanged($event)"
        ></v-text-field>
      </v-col>
      <v-col cols="2" class="d-flex justify-end pr-1">
        <v-btn dark rounded :color="colorTheme" :to="{ name: 'ChatList' }">
          <v-icon class="hidden-lg-and-up">mdi-account-multiple</v-icon>
          <span class="hidden-md-and-down">Back to chat list</span>
        </v-btn>
      </v-col>
      <v-col
        cols="12"
        v-for="(item, index) in search_users ? search_users.search_result : []"
        :key="index"
        class="mb-3 px-5 py-1 white rounded-lg"
      >
        <UserListItem
          :username="item.username"
          :userId="item.user_id"
          @on-chat-button-click="createChatRoom($event)"
        />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import UserListItem from "@/components/UserListItem.vue";
import SEARCH_USERS_QUERY from "@/gql/queries/SearchUsers.gql";
import CREATE_CHAT_ROOM_MUTATION from "@/gql/mutations/CreateChatRoom.gql";
import { EventBus } from "@/bus";

export default {
  name: "SearchUsers",

  components: {
    UserListItem,
  },

  data: () => ({
    colorTheme: "#394867",
    searchQuery: "",
    searchLimit: 10,
    searchOffset: 0,
    lastSetTimeoutId: null,
  }),

  apollo: {
    search_users: {
      query: SEARCH_USERS_QUERY,
      variables() {
        return {
          searchQuery: this.searchQueryInLowerCase,
          limit: this.searchLimit,
          offset: this.searchOffset,
        };
      },
    },
  },

  computed: {
    searchQueryInLowerCase() {
      return this.searchQuery.toLowerCase();
    },
  },

  methods: {
    onSearchQueryChanged(query) {
      if (this.lastSetTimeoutId !== null) {
        clearTimeout(this.lastSetTimeoutId);
      }
      this.lastSetTimeoutId = setTimeout(() => {
        this.searchQuery = query;
      }, 2000);
    },

    async createChatRoom(recipientId) {
      try {
        let result = await this.$apollo.mutate({
          mutation: CREATE_CHAT_ROOM_MUTATION,
          variables: {
            memberIds: [recipientId],
          },
          optimisticResponse: {
            __typename: "Mutation",
            create_chat_room: {
              __typename: "ChatRoomType",
              chat_room_id: -1,
              created_at: "",
              updated_at: "",
              deleted_at: "",
            },
          },
        });

        if (result !== null) {
          this.$router.push({
            name: "ChatRoom",
            params: {
              chatRoomId: result.data.create_chat_room.chat_room_id,
            },
          });
        } else {
          EventBus.$emit(
            "onShowSnackbar",
            "Something's wrong. Please try again."
          );
        }
      } catch (error) {
        EventBus.$emit(
          "onShowSnackbar",
          "Something's wrong. Please try again."
        );
      }
    },
  },

  beforeDestroy() {
    clearTimeout(this.lastSetTimeoutId);
  },
};
</script>