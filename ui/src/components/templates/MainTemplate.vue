<template>
  <div class="bg-surface-dark flex rounded-xl overflow-hidden">
    <Aside class="p-5" />
    <div class="bg-surface-primary p-5">
      <div class="text-xl font-bold mb-5">
        Данные об образовании
      </div>
      <StudyingForm />
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
        <Button class="w-full p-3 bg-secondary-disabled">
          Далее
        </Button>
        <Button class="w-full bg-surface-tertiary p-3">
          Я не студент Мосполитеха
        </Button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import Button from '@/components/atoms/Button.vue'
import Checkbox from '@/components/atoms/Checkbox.vue'
import { ref } from 'vue'
import Aside from '../organisms/Aside.vue'
import StudyingForm from '../organisms/StudyingForm.vue'

const checkbox = ref(false)
</script>

<style scoped>

</style>
