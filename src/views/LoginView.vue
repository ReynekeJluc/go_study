<script setup>
import { useAuthStore } from '@/stores';
import { Field, Form } from 'vee-validate';
import * as Yup from 'yup';

const schema = Yup.object().shape({
	login: Yup.string().required('Login is required'),
	password: Yup.string().required('Password is required'),
});

async function onSubmit(values, { setErrors }) {
	const authStore = useAuthStore();
	const { login, password } = values;

	try {
		await authStore.login(login, password);
	} catch (error) {
		setErrors({ apiError: error.message });
	}
}
</script>

<template>
	<div>
		<h2>Login</h2>
		<Form
			@submit="onSubmit"
			:validation-schema="schema"
			v-slot="{ errors, isSubmitting }"
		>
			<div class="form-group">
				<label>Login</label>
				<Field
					name="login"
					type="text"
					class="form-control"
					:class="{ 'is-invalid': errors.login }"
				/>
				<div class="invalid-feedback">{{ errors.login }}</div>
			</div>
			<div class="form-group">
				<label>Password</label>
				<Field
					name="password"
					type="password"
					class="form-control"
					:class="{ 'is-invalid': errors.password }"
				/>
				<div class="invalid-feedback">{{ errors.password }}</div>
			</div>
			<div class="form-group">
				<button class="btn btn-primary" :disabled="isSubmitting">
					<span
						v-show="isSubmitting"
						class="spinner-border spinner-border-sm mr-1"
					></span>
					Login
				</button>
			</div>
			<div v-if="errors.apiError" class="alert alert-danger mt-3 mb-0">
				{{ errors.apiError }}
			</div>
		</Form>
	</div>
</template>
