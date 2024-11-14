<template>
  <div class="flex flex-col justify-between">
    <div>
      <div>
        <img src="/title.svg" alt="title" class="mb-8">
      </div>
      <Stepper />
    </div>
    <Button class="text-text-alert font-bold no-wrap sm:flex hidden">
      <template #prepend>
        <img src="/face.svg" alt="face icon" class="mr-2">
      </template>
      <div class="text-left">
        Не получается войти
      </div>
    </Button>
  </div>
</template>

<script setup lang="ts">
import Button from '@/components/atoms/Button.vue'
import Stepper from '@/components/molecules/Stepper.vue'
</script>

<style scoped>

</style>
