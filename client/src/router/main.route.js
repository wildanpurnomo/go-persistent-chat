export const mainRouter = {
    path: '/main',
    component: () => import('@/pages/Main.vue'),
    children: [
        {
            path: '/main/chat/list',
            name: 'ChatList',
            component: () => import('@/partial-views/main/ChatList.vue')
        },
    ]
}