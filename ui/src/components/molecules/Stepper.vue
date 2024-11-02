<template>
  <div class="text-[14px]">
    <div class="checkbox" @click.stop.prevent>
      <Checkbox v-model="test1" label="Данные об образовании" />
    </div>
    <div class="checkbox" @click.stop.prevent>
      <Checkbox v-model="test2" label="Контактные данные" />
    </div>
    <div class="checkbox" @click.stop.prevent>
      <Checkbox v-model="test3" label="Личные документы" />
    </div>
  </div>
</template>

<script setup lang="ts">
import Checkbox from '@/components/atoms/Checkbox.vue'
import { ref } from 'vue'

interface Props {
  items: Array<any>
}
const props = defineProps<Props>()

const test1 = ref(false)
const test2 = ref(false)
const test3 = ref(false)
</script>

<style scoped>
.checkbox:after {
  display: block;
  content: ' ';
  height: 48px;
  width: 1px;
  margin: 6px 0 6px 10px;
  background: #6F6E6F;
}
.checkbox:after {
  display: block;
  content: ' ';
  height: 48px;
  width: 1px;
  margin: 6px 0 6px 10px;
  background: #6F6E6F;
}

.checkbox:last-child:after {
  display: none;
}
</style>
