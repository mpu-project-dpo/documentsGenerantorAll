<template>
  <form class="flex flex-col gap-3" @submit.prevent>
    <Input
      v-for="key in Object.keys(education)"
      :key="key"
      v-model="education[key].value"
      :mask="education[key].mask"
      :label="`${education[key].field} *`"
      class="w-full"
      clearable
    />
  </form>
</template>

<script setup lang="ts">
import { useFormStore } from '@/store/formStore.ts'
import { storeToRefs } from 'pinia'
import Input from '../atoms/Input.vue'

const { education } = storeToRefs(useFormStore())
</script>

<style scoped>
</style>
