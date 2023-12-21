<script setup lang="ts">
import { RouterLink } from "vue-router";
</script>

<template>
  <header class="h-16 bg-primary-1000">
    <nav class="container m-auto h-full flex">
      <div class="text-white my-auto flex flex-row">
        <div class="h-10 w-10 bg-[url('creeper.png')] bg-contain">
        </div>
        <div class="my-auto ml-5">
          <RouterLink to="/">Home</RouterLink>
        </div>
      </div>
    </nav>
  </header>
</template>

<style scoped>
</style>
