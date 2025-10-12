import { fetchWrapper, router } from '@/helpers';
import { defineStore } from 'pinia';

const baseUrl = `${import.meta.env.VITE_API_URL}/api/auth`;

export const useAuthStore = defineStore({
	id: 'auth',
	state: () => ({
		access_token: localStorage.getItem('access_token') || null,
		refresh_token: localStorage.getItem('refresh_token') || null,
		returnUrl: null,
	}),
	actions: {
		async login(login, password) {
			try {
				const tokens = await fetchWrapper.post(`${baseUrl}/login`, {
					login,
					password,
				});

				this.access_token = tokens.access_token;
				this.refresh_token = tokens.refresh_token;

				localStorage.setItem('access_token', tokens.access_token);
				localStorage.setItem('refresh_token', tokens.refresh_token);

				router.push(this.returnUrl || '/books');
			} catch (err) {
				console.error('Login failed:', err);
				throw err;
			}
		},

		logout() {
			this.access_token = null;
			this.refresh_token = null;

			localStorage.removeItem('access_token');
			localStorage.removeItem('refresh_token');

			router.push('/login');
		},

		async refreshAccessToken() {
			if (!this.refresh_token) throw new Error('No refresh token');

			const tokens = await fetchWrapper.post(`${baseUrl}/refresh`, {
				refresh_token: this.refresh_token,
			});

			this.access_token = tokens.access_token;
			this.refresh_token = tokens.refresh_token;

			localStorage.setItem('access_token', tokens.access_token);
			localStorage.setItem('refresh_token', tokens.refresh_token);
		},

		isLoggedIn() {
			return !!this.access_token;
		},
	},
});
