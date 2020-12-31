
const routes = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('components/Login.vue') },
      { path: '/signup', component: () => import('components/Signup.vue') },
      { path: '/app/home/:id', name: 'Home', props: true, component: () => import('pages/HomePage.vue') },
      { path: '/app/addImage/:id', name: 'AddImage', props: true, component: () => import('pages/AddImage.vue') },
    ]
  },
  {
    path: '*',
    component: () => import('pages/Error404.vue')
  }
]

export default routes
