<template>
  <form class="flex flex-col gap-4">
    <Input label="Университет *" class="w-full" />
    <Input label="Группа *" class="w-full" />
    <Input label="ДПО *" class="w-full">
      <template #append>
        <div>123</div>
      </template>
    </Input>
  </form>
</template>

<script setup lang="ts">
import Input from '../atoms/Input.vue'
</script>

<style scoped>
</style>
