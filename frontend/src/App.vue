<script setup lang="ts">
import {RouterView} from 'vue-router'
import HeaderComponent from "@/components/HeaderComponent.vue";
import {useUserStore} from "@/stores/user";
import router from "@/router";
import {onMounted} from "vue";

const userStore = useUserStore();

onMounted(() => {
  fetch('http://localhost:8080/authenticated', {
    method: 'GET',
  }).then(res => res.json()).then(data => {
    if (data.authenticated) {
      userStore.user = 'demo';
    } else {
      router.push('/login');
    }
  });
});

</script>

<template>
  <HeaderComponent></HeaderComponent>
  <main class="container mx-auto mt-5">
    <RouterView />
  </main>
</template>

<style scoped>
</style>
