<template>
  <div class="bg-surface-dark flex flex-nowrap rounded-xl flex-col sm:flex-row">
    <Aside class="p-5" />
    <div
      class="
      bg-surface-primary p-5

      sm:rounded-br-xl
      sm:rounded-tr-xl
      sm:rounded-bl-none

      rounded-bl-xl
      rounded-br-xl
      rounded-tr-none"
    >
      <div class="text-xl font-bold mb-5">
        <span v-if="$route.name === 'education'">Данные об образовании</span>
        <span v-else-if="$route.name === 'contacts'">Контактные данные</span>
        <span v-else>Личные документы</span>
      </div>

      <router-view />

      <div class="flex my-5 text-[13px] items-center">
        <Checkbox v-model="store.confPolicy" active-color="secondary">
          <template #label>
            <div>
              Согласен с <a
                class="text-secondary inline"
                :href="confPolicyLink"
                target="_blank"
              >политикой конфиденциальности</a>
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
          <Button
            class="rounded-lg w-full p-2.5 bg-secondary-enabled disabled:bg-secondary-disabled"
            :disabled="isButtonDisabled || (route.name === 'docs' && !store.confPolicy)"
            @click="onContinue"
          >
            {{ route.name === 'docs' ? 'Отправить документы' : 'Далее' }}
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
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Aside from '../organisms/Aside.vue'

const confPolicyLink = import.meta.env.VITE_CONF_POLICY_LINK

const router = useRouter()
const route = useRoute()
const store = useFormStore()

const isButtonDisabled = computed(() => {
  const routeName = route.name as string
  if (Object.keys(store).includes(routeName)) {
    return store.findEmptyField(store[routeName as 'education' || 'docs' || 'contacts'])
  }
  return false
})

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
