<template>
  <div class="flex flex-col">
    <img src="/title.svg" alt="title" class="mb-8">
    <Stepper :items="[1, 2, 3]" />
  </div>
</template>

<script setup lang="ts">
import Stepper from '@/components/molecules/Stepper.vue'
</script>

<style scoped>

</style>
