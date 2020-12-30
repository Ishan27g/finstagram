import Signup from '../components/SignUp.vue';
import Login from '../components/Login.vue';

const routes = [
  {
    path: '/',
    component: () => import('layouts/EntryLayout.vue'),
    children: [
      { path: '/signup', name: "Signup", component: Signup },
      { path: '/login', name: "Login", component: Login },
    ]
  },
  {
    path: '/app',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '/app/home/:userId', name: 'Home', component: () => import('pages/HomePage.vue') },
      { path: '/app/addImage/:userId', name: 'AddImage', component: () => import('pages/AddImage.vue') },
    ]
  },
  {
    path: '*',
    component: () => import('pages/Error404.vue')
  }
]

export default routes
