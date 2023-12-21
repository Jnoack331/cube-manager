<script setup lang="ts">
import ButtonComponent from "@/components/buttons/ButtonComponent.vue";
</script>

<template>
  <ButtonComponent class="bg-secondary text-primary-900 hover:bg-primary-800 hover:text-secondary ">
    <slot/>
  </ButtonComponent>
</template>
