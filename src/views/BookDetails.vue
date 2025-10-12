<script setup>
import { fetchWrapper } from '@/helpers';
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute();

const book = ref(null);
const loading = ref(true);
const error = ref(null);

onMounted(async () => {
	try {
		const id = route.params.BookId;
		const url = `${import.meta.env.VITE_API_URL}/api/books/${id}`;
		book.value = await fetchWrapper.get(url);
	} catch (err) {
		error.value = err.message || 'Ошибка загрузки данных';
	} finally {
		loading.value = false;
	}
});
</script>

<template>
	<div class="container mt-4">
		<div v-if="loading" class="text-center mt-5">
			<div class="spinner-border" role="status">
				<span class="visually-hidden">Загрузка...</span>
			</div>
		</div>

		<div v-else-if="error" class="alert alert-danger text-center">
			{{ error }}
		</div>

		<div v-else-if="book" class="card shadow-sm p-4">
			<h2 class="mb-3">{{ book.book_name }}</h2>
			<p class="text-muted mb-1">
				<strong>Издательство:</strong> {{ book.book_publisher }}
			</p>
			<p class="text-muted mb-3">
				<strong>Год:</strong> {{ book.book_year_release }}
			</p>

			<p class="mb-4">{{ book.book_desc }}</p>

			<ul class="list-group mb-4">
				<li class="list-group-item">
					<strong>Страниц:</strong> {{ book.book_pages_amount }}
				</li>
				<li class="list-group-item">
					<strong>Количество на складе:</strong> {{ book.book_total_quantity }}
				</li>
				<li class="list-group-item">
					<strong>ISBN:</strong> {{ book.book_isbn }}
				</li>
			</ul>

			<div class="d-flex justify-content-between align-items-center">
				<h4 class="mb-0">Цена: {{ book.book_price }} ₽</h4>
				<button class="btn btn-primary btn-lg">Купить</button>
			</div>
		</div>

		<div v-else class="text-center text-muted mt-4">Книга не найдена.</div>
	</div>
</template>

<style scoped>
.container {
	max-width: 800px;
}
.card {
	border: none;
}
</style>
