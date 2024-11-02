<template>
  <div class="text-[14px]">
    <div class="checkbox" @click.stop.prevent="$router.push('/')">
      <Checkbox v-model="education" label="Данные об образовании" :bordered="$route.name === 'education'" />
    </div>
    <div class="checkbox" @click.stop.prevent="$router.push('/contacts')">
      <Checkbox v-model="contacts" label="Контактные данные" :bordered="$route.name === 'contacts'" />
    </div>
    <div
      class="checkbox"

      @click.stop.prevent="$router.push('/docs')"
    >
      <Checkbox v-model="docs" label="Личные документы" :bordered="$route.name === 'docs'">
        />
      </checkbox>
    </div>
  </div>
</template>

<script setup lang="ts">
import Checkbox from '@/components/atoms/Checkbox.vue'
import { useFormStore } from '@/store/formStore.ts'
import { ref, watch } from 'vue'

const store = useFormStore()

const education = ref(false)
const contacts = ref(false)
const docs = ref(false)

watch(store.education, () => {
  education.value = !store.checkFields(store.education)
  console.log(store.education)
})
watch(store.contacts, () => {
  contacts.value = !store.checkFields(store.contacts)
})
watch(store.docs, () => {
  docs.value = !store.checkFields(store.docs)
})
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
