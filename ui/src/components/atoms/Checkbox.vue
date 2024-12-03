<template>
  <label
    class="flex flex-row items-center gap-2.5 cursor-pointer"
  >
    <input id="hr" v-model="modelValue" type="checkbox" class="hidden">
    <div
      class="
        h-5 w-5 flex items-center justify-center rounded-md
      "
      :class="{
        'bg-secondary': modelValue && props.activeColor === 'secondary',
        'bg-surface-secondary': (!modelValue) || modelValue && props.activeColor === 'primary',
        'border-2 border-surface-tertiary': bordered,
      } "
    >
      <svg v-if="modelValue" width="10" height="8" viewBox="0 0 10 8" fill="none" xmlns="http://www.w3.org/2000/svg">
        <path
          d="M9.1001 1.09827L3.6001 6.59827L1.1001 4.09827"
          stroke="white"
          stroke-width="0.9" stroke-linecap="round"
          stroke-linejoin="round"
        />
      </svg>
    </div>
    <div><slot name="label">{{ props.label }}</slot></div>
  </label>
</template>

<script setup lang="ts">
interface Props {
  label?: string
  activeColor?: 'primary' | 'secondary'
  bordered?: boolean
  disabled?: boolean
}
const props = withDefaults(defineProps<Props>(), { activeColor: 'primary' })
const modelValue = defineModel<boolean>({ required: true })
</script>

<style scoped>

</style>
