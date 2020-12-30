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
      { path: '/app/home', component: () => import('pages/HomePage.vue') },
      { path: '/app/addImage', component: () => import('pages/AddImage.vue') },
    ]
  },
  {
    path: '*',
    component: () => import('pages/Error404.vue')
  }
]

export default routes
