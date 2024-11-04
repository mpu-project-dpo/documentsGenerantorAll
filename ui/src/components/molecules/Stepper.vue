<template>
  <div class="text-[14px] grid grid-cols-3 sm:grid-cols-1">
    <div
      class="col-span-2 sm:col-span-1"
    >
      <div
        v-if="windowWidth > 640 || isDropdownOpened || (!isDropdownOpened && $route.meta.pageNumber === 1)"
        class="checkbox"
        @click.stop.prevent="goToPage('education', 'education')"
      >
        <Checkbox
          v-model="education"
          label="Данные об образовании"
          :bordered="$route.name === 'education'"
        />
      </div>
      <div
        v-if="windowWidth > 640 || isDropdownOpened || (!isDropdownOpened && $route.meta.pageNumber === 2)"
        class="checkbox"
        @click.stop.prevent="goToPage('contacts', 'education')"
      >
        <Checkbox
          v-model="contacts"
          label="Контактные данные"
          :bordered="$route.name === 'contacts'"
        />
      </div>
      <div
        v-if="windowWidth > 640 || isDropdownOpened || (!isDropdownOpened && $route.meta.pageNumber === 3)"
        class="checkbox"
        @click.stop.prevent="goToPage('docs', 'contacts')"
      >
        <Checkbox
          v-model="docs"
          label="Личные документы"
          :bordered="$route.name === 'docs'"
        />
      </div>
    </div>
    <div class="justify-self-end sm:hidden col-span-1">
      <div class="flex flex-nowrap items-center" @click="isDropdownOpened = !isDropdownOpened">
        <div class="text-placeholder-color font-bold mr-2">
          {{ $route.meta.pageNumber }}/3
        </div>
        <Dropdown :class="{ 'rotate-180': !isDropdownOpened }" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import Dropdown from '@/assets/icons/dropdown.vue'
import Checkbox from '@/components/atoms/Checkbox.vue'
import { useFormStore } from '@/store/formStore.ts'
import { ref, watch } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const store = useFormStore()

const education = ref(false)
const contacts = ref(false)
const docs = ref(false)

const isDropdownOpened = ref(false)

const windowWidth = ref(window.innerWidth)

watch(store.education, () => {
  education.value = !store.findEmptyField(store.education)
})

watch(store.contacts, () => {
  contacts.value = !store.findEmptyField(store.contacts)
})

watch(store.docs, () => {
  docs.value = !store.findEmptyField(store.docs)
})

function goToPage(page: 'education' | 'contacts' | 'docs', dependsOn: typeof page) {
  if (!store.findEmptyField(store[dependsOn])) {
    router.push({ name: page })
  }
}
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

@media (max-width: 640px) {
  .checkbox:after {
    height: 8px
  }
}
</style>
