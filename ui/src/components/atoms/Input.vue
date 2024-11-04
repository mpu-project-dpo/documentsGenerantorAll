<template>
  <label class="font-bold text-[13px]">
    <div v-if="props.label" class="label mb-1">{{ props.label }}</div>
    <div
      class="
        bg-surface-secondary
        focus-within:ring-2
        focus-within:ring-surface-tertiary
        p-1.5
        rounded-lg
        border-2
        border-surface-tertiary
        flex
        flex-nowrap
        items-center
        justify-between
      "
      tabindex="0"
      @click="inputRef?.focus()"
    >
      <input
        v-if="props.mask"
        ref="inputRef"
        v-model="modelValue"
        v-mask="props.mask"
        type="text"
        class="bg-transparent min-h-6 text-white outline-none placeholder:text-placeholder-color placeholder:font-thin grow"
        :placeholder="props.placeholder"
        @focus="emits('focus')"
      >
      <input
        v-else
        ref="inputRef"
        v-model="modelValue"
        type="text"
        class="bg-transparent min-h-6 text-white outline-none placeholder:text-placeholder-color placeholder:font-thin grow"
        :placeholder="props.placeholder"
        @focus="emits('focus')"
      >
      <div class="ml-1"><slot name="append" /></div>
      <div v-if="props.clearable && modelValue" class="w-6 h-6 ml-1">
        <Button v-if="modelValue" class="bg-surface-dark rounded" @click.prevent="modelValue = ''">
          <Close class="m-0.5" />
        </Button>
      </div>
    </div>
  </label>
</template>

<script setup lang="ts">
import Close from '@/assets/icons/close.vue'
import Button from '@/components/atoms/Button.vue'
import { ref } from 'vue'

interface Props {
  label?: string
  placeholder?: string
  clearable?: boolean
  mask?: string
}

interface Emits {
  (name: 'focus'): void
}

const props = defineProps<Props>()
const emits = defineEmits<Emits>()

const inputRef = ref<HTMLInputElement | null>(null)
const modelValue = defineModel<string>({ required: true })
</script>

<style scoped>
.label {
  font-size: 14px;
}
</style>
