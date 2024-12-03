<template>
  <form class="flex flex-col gap-3" @submit.prevent>
    <Input
      v-for="key in Object.keys(contacts)"
      :key="key"
      v-model="contacts[key].value"
      :mask="contacts[key].mask"
      :label="`${contacts[key].field} *`"
      class="w-full"
      clearable
    />
  </form>
</template>

<script setup lang="ts">
import { useFormStore } from '@/store/formStore.ts'
import { storeToRefs } from 'pinia'
import Input from '../atoms/Input.vue'

const { contacts } = storeToRefs(useFormStore())
</script>

<style scoped>
</style>
