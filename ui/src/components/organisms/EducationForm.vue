<template>
  <form class="flex flex-col gap-3" @submit.prevent>
    <template
      v-for="key in Object.keys(education)"
      :key="key"
    >
      <Input
        v-if="key !== 'dpo'"
        v-model="education[key].value"
        :mask="education[key].mask"
        :label="`${education[key].field} *`"
        class="w-full"
        clearable
      />
      <Dropdown
        v-else
        v-model="education[key].value"
        v-model:search="searchModel"
        :options="['1', '2', '3', '4', '5', '6', '7', '8', '9']"
        :label="`${education[key].field} *`"
        class="w-full"
        clearable
      />
    </template>
  </form>
</template>

<script setup lang="ts">
import Dropdown from '@/components/atoms/Dropdown.vue'
import { useFormStore } from '@/store/formStore.ts'
import { storeToRefs } from 'pinia'
import { ref } from 'vue'
import Input from '../atoms/Input.vue'

const { education } = storeToRefs(useFormStore())

const searchModel = ref('')
</script>

<style scoped>
</style>
