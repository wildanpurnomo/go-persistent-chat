export const mainRouter = {
  path: "/main",
  component: () => import("@/pages/Main.vue"),
  children: [
    {
      path: "/main/chat/list",
      name: "ChatList",
      component: () => import("@/partial-views/main/ChatList.vue"),
    },
    {
      path: "/main/search/users",
      name: "SearchUsers",
      component: () => import("@/partial-views/main/SearchUsers.vue"),
    },
    {
      path: "/main/chat/room/:chatRoomId",
      name: "ChatRoom",
      component: () => import("@/partial-views/main/ChatRoom.vue"),
    },
  ],
};
