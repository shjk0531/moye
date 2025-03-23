import { createRouter, createWebHistory } from "vue-router";
import { AuthPage } from "@/pages";
import { AppLayout } from "@/shared/layout";

const routes = [
  {
    path: "/",
    component: AppLayout,
    children: [{ path: "/login", component: AuthPage }],
  },
  // HomePage는 필요에 따라 추가하세요.
  // { path: '/', component: () => import('@/pages/auth') },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
