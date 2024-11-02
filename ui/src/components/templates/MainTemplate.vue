<template>
  <div class="bg-surface-dark flex rounded-xl overflow-hidden">
    <Aside class="p-5" />
    <div class="bg-surface-primary p-5">
      <div class="text-xl font-bold mb-5">
        <span v-if="$route.name === 'education'">Данные об образовании</span>
        <span v-else-if="$route.name === 'contacts'">Контактные данные</span>
        <span v-else>Личные документы</span>
      </div>

      <router-view />

      <div class="flex my-5 text-[13px] items-center">
        <Checkbox v-model="checkbox" active-color="secondary">
          <template #label>
            <div>
              Согласен с <a class="text-secondary inline" @click.stop.prevent>политикой конфиденциальности</a>
            </div>
          </template>
        </Checkbox>
      </div>
      <div class="actions flex flex-col gap-2 font-bold text-[13px]">
        <div class="flex items-center">
          <div v-if="route.name !== 'education'" class="">
            <Button @click="onBack">
              <Back class="w-5 h-5 mr-3.5" />
            </Button>
          </div>
          <Button class="rounded-lg w-full p-2.5 bg-secondary-disabled" @click="onContinue">
            Далее
          </Button>
        </div>
        <Button v-if="route.name === 'education'" class="rounded-lg w-full bg-surface-tertiary p-2.5">
          Я не студент Мосполитеха
        </Button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import Back from '@/assets/icons/back.vue'
import Button from '@/components/atoms/Button.vue'
import Checkbox from '@/components/atoms/Checkbox.vue'
import { useFormStore } from '@/store/formStore.ts'
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Aside from '../organisms/Aside.vue'

const router = useRouter()
const route = useRoute()
const store = useFormStore()

const checkbox = ref(false)

function onContinue() {
  if (route.name === 'education') {
    router.push('/contacts')
  }
  if (route.name === 'contacts') {
    router.push('/docs')
  }
  if (route.name === 'docs') {
    store.sendForm()
  }
}

function onBack() {
  if (route.name === 'education') {
    return
  }
  if (route.name === 'contacts') {
    router.push('/')
  }
  if (route.name === 'docs') {
    router.push('/contacts')
  }
}
</script>
