<template>
  <label class="font-bold text-[13px]">
    <div v-if="props.label" class="label mb-2">{{ props.label }}</div>
    <div
      class="
        bg-surface-secondary
        focus-within:ring-2
        focus-within:ring-surface-tertiary
        px-2
        py-1.5
        rounded-lg
        border-2
        border-surface-tertiary
        flex
        flex-nowrap
        justify-between
      "
      tabindex="0"
      @click="inputRef?.focus()"
    >
      <input
        ref="inputRef"
        v-model="modelValue"

        type="text"
        class="bg-transparent text-white outline-none placeholder:text-placeholder-color placeholder:font-thin grow"
        :placeholder="props.placeholder"
      >
      <div class="ml-1"><slot name="append" /></div>
    </div>
  </label>
</template>

<script setup lang="ts">
import { ref } from 'vue'

interface Props {
  label?: string
  placeholder?: string
}

const props = defineProps<Props>()
const inputRef = ref<HTMLInputElement | null>(null)
const modelValue = defineModel<string>({ required: true })
</script>

<style scoped>
.label {
  font-size: 14px;
}
</style>
