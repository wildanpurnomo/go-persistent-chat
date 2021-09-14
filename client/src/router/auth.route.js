export const authRouter = {
    path: '/',
    component: () => import('@/pages/Auth.vue'),
    children: [
        {
            path: '/login',
            name: 'Login',
            component: () => import('@/partial-views/auth/Login.vue')
        },
        {
            path: '/register',
            name: 'Register',
            component: () => import('@/partial-views/auth/Register.vue')
        }
    ]
}