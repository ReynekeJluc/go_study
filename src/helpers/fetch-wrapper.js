import { useAuthStore } from '@/stores';

export const fetchWrapper = {
	get: request('GET'),
	post: request('POST'),
	put: request('PUT'),
	delete: request('DELETE'),
};

function request(method) {
	return async (url, body) => {
		const authStore = useAuthStore();
		const isApiUrl = url.startsWith(import.meta.env.VITE_API_URL);

		let headers =
			isApiUrl && authStore.access_token
				? { Authorization: `Bearer ${authStore.access_token}` }
				: {};

		if (body) headers['Content-Type'] = 'application/json';

		let response = await fetch(url, {
			method,
			headers,
			body: body ? JSON.stringify(body) : undefined,
		});

		// === если токен истёк, пробуем обновить ===
		if (response.status === 401 && authStore.refresh_token) {
			try {
				await authStore.refreshAccessToken();

				// перезапрос с новым access_token
				headers = {
					...headers,
					Authorization: `Bearer ${authStore.access_token}`,
				};

				response = await fetch(url, {
					method,
					headers,
					body: body ? JSON.stringify(body) : undefined,
				});
			} catch (err) {
				authStore.logout();
				throw new Error('Session expired. Please login again.');
			}
		}

		const text = await response.text();
		const data = text ? JSON.parse(text) : null;

		if (!response.ok) {
			if ([401, 403].includes(response.status) && authStore.isLoggedIn()) {
				authStore.logout();
			}
			const error = (data && data.message) || response.statusText;
			throw new Error(error);
		}

		return data;
	};
}
