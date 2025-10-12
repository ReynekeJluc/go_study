import { useAuthStore } from '@/stores';
import HomeView from '@/views/HomeView.vue';
import LoginView from '@/views/LoginView.vue';
import { createRouter, createWebHistory } from 'vue-router';
import BookDetails from '../views/BookDetails.vue';

export const router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	routes: [
		{ path: '/books', component: HomeView },
		{ path: '/books/:BookId', component: BookDetails },
		{ path: '/login', component: LoginView },
		{ path: '/:pathMatch(.*)*', redirect: '/login' },
	],
});

router.beforeEach(async to => {
	const publicPages = ['/login'];
	const authRequired = !publicPages.includes(to.path);
	const auth = useAuthStore();

	const isLoggedIn =
		!!auth.access_token || !!localStorage.getItem('access_token');

	if (authRequired && !isLoggedIn) {
		auth.returnUrl = to.fullPath;
		return '/login';
	}
});
