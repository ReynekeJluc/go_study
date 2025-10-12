<script setup>
import { fetchWrapper } from '@/helpers';
import { computed, onMounted, ref } from 'vue';

const books = ref([]);
const loading = ref(false);
const error = ref(null);
const searchQuery = ref('');

onMounted(async () => {
	try {
		loading.value = true;
		error.value = null;

		const res = await fetchWrapper.get(
			`${import.meta.env.VITE_API_URL}/api/books`
		);
		books.value = res;
	} catch (err) {
		error.value = err.message || 'Ошибка загрузки книг';
	} finally {
		loading.value = false;
	}
});

// фильтруем книги по поисковому запросу
const filteredBooks = computed(() => {
	const q = searchQuery.value.toLowerCase().trim();
	if (!q) return books.value;
	return books.value.filter(b => b.book_name.toLowerCase().includes(q));
});
</script>

<template>
	<div class="container mt-4">
		<h1 class="mb-4">Каталог книг</h1>

		<!-- Строка поиска -->
		<div class="form-group mb-3">
			<input
				v-model="searchQuery"
				type="text"
				class="form-control"
				placeholder="Поиск по названию книги..."
			/>
		</div>

		<!-- Загрузка -->
		<div v-if="loading" class="spinner-border spinner-border-sm"></div>
		<!-- Ошибка -->
		<div v-if="error" class="text-danger">{{ error }}</div>

		<!-- Список книг -->
		<div v-if="filteredBooks.length" class="book-list">
			<div
				v-for="book in filteredBooks"
				:key="book.book_id"
				class="book-item border p-3 mb-3 rounded"
			>
				<h3>
					<router-link :to="`/books/${book.book_id}`">
						{{ book.book_name }}
					</router-link>
				</h3>
				<p><strong>Издатель:</strong> {{ book.book_publisher }}</p>
				<p><strong>Год выпуска:</strong> {{ book.book_year_release }}</p>
				<p><strong>Страниц:</strong> {{ book.book_pages_amount }}</p>
				<p><strong>Цена:</strong> {{ book.book_price }} ₽</p>
				<p class="text-muted small">{{ book.book_desc }}</p>
			</div>
		</div>

		<div v-else-if="!loading && !error" class="text-muted">
			Нет книг, подходящих под запрос.
		</div>
	</div>
</template>
