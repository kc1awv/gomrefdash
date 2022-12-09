const routes = [
  {
    path: "/",
    component: () => import("layouts/MainLayout.vue"),
    children: [{ path: "", component: () => import("pages/IndexPage.vue") }],
  },
  {
    path: "/links",
    component: () => import("layouts/MainLayout.vue"),
    children: [{ path: "", component: () => import("pages/LinksPage.vue") }],
  },
  {
    path: "/peers",
    component: () => import("layouts/MainLayout.vue"),
    children: [{ path: "", component: () => import("pages/PeersPage.vue") }],
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: "/:catchAll(.*)*",
    component: () => import("pages/ErrorNotFound.vue"),
  },
];

export default routes;
