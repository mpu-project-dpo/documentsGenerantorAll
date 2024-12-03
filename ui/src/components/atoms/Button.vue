<template>
  <button
    class="flex flex-nowrap items-center"
    :disabled="props.disabled"
  >
    <div v-if="$slots.prepend">
      <slot name="prepend" />
    </div>
    <div class="grow">
      <slot />
    </div>
    <div v-if="$slots.append">
      <slot name="append" />
    </div>
  </button>
</template>

<script setup lang="ts">
interface Props {
  disabledClass?: string
  disabled?: boolean
}

const props = defineProps<Props>()
</script>

<style scoped>

</style>
