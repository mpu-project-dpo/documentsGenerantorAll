<template>
  <button class="flex flex-nowrap items-center">
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

</script>

<style scoped>

</style>
