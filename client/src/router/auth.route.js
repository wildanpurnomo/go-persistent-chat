export const authRouter = {
    path: '/auth',
    component: () => import('@/pages/Auth.vue'),
    children: [
        {
            path: '/auth/login',
            name: 'Login',
            component: () => import('@/partial-views/auth/Login.vue')
        },
        {
            path: '/auth/register',
            name: 'Register',
            component: () => import('@/partial-views/auth/Register.vue')
        }
    ]
}