<template>
  <form class="flex flex-col gap-3" @submit.prevent>
    <Input
      v-for="key in Object.keys(docs)"
      :key="key"
      v-model="docs[key].value"
      :mask="docs[key].mask"
      :label="`${docs[key].field} *`"
      class="w-full"
      clearable
    />
  </form>
</template>

<script setup lang="ts">
import { useFormStore } from '@/store/formStore.ts'
import { storeToRefs } from 'pinia'
import Input from '../atoms/Input.vue'

const { docs } = storeToRefs(useFormStore())
</script>

<style scoped>
</style>
