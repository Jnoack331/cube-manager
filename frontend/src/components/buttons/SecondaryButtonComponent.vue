<script setup lang="ts">
import ButtonComponent from "@/components/buttons/ButtonComponent.vue";
</script>

<template>
  <ButtonComponent class="bg-secondary-700 text-primary-text hover:bg-secondary-600">
    <slot/>
  </ButtonComponent>
</template>
