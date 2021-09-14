export const fallbackRouter = {
    path: '*',
    name: 'Fallback',
    component: () => import('@/pages/404.vue')
}