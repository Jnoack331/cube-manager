<script setup lang="ts">
import PrimaryButtonComponent from "@/components/buttons/PrimaryButtonComponent.vue";
import {ref} from "vue";
import {useUserStore} from "@/stores/user";
import router from "@/router";

const username = ref('');
const password = ref('');

const userStore = useUserStore();


function login(event: any): void {
  event.preventDefault();
  const formData = new FormData(event.target);
  fetch('http://localhost:8080/login', {
    method: 'POST',
    body: formData
  }).then(res => res.json()).then(data => {
    if (data.authenticated) {
      userStore.user = 'demo';
      router.push('/');
    }
  });
}
</script>

<template>
    <form class="bg-primary-1000 p-10 rounded absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2" @submit="login">
      <div class="grid gap-3">
        <div>
          <h1 class="text-3xl mb-5 text-secondary">Login</h1>
          <div class="flex flex-col">
            <input type="text" name="username" v-model="username" class="mb-2 text-secondary font-semibold rounded bg-primary-900">
            <input type="password" name="password" v-model="password" class="rounded text-secondary bg-primary-900 font-semibold">
            <PrimaryButtonComponent>Login</PrimaryButtonComponent>
          </div>
        </div>
      </div>
    </form>
</template>
