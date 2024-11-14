<template>
  <div class="relative w-full">
    <div class="cursor-pointer w-full">
      <Input
        v-model="searchModel"
        :label="props.label"
        :placeholder="props.placeholder"
        :clearable="props.clearable"
        @focus="toggleDropdown"
        @click="isOpen = true"
      />

      <transition name="fade">
        <div
          v-if="isOpen"
          ref="target"
          class="absolute mt-2 w-full bg-surface-secondary rounded-lg shadow-lg z-10 overflow-y-auto max-h-44"
        >
          <ul>
            <li
              v-for="option in filteredOptions"
              :key="option"
              class="px-4 py-2 hover:bg-secondary cursor-pointer font-bold"
              @click="selectOption(option)"
            >
              {{ option }}
            </li>
          </ul>
        </div>
      </transition>
    </div>
  </div>
</template>

<script setup lang="ts">
import Input from '@/components/atoms/Input.vue'
import { onClickOutside } from '@vueuse/core'
import { computed, onMounted, onUnmounted, ref } from 'vue'

interface Props {
  label?: string
  placeholder?: string
  clearable?: boolean
  options: string[]
}

const props = defineProps<Props>()

const searchModel = defineModel<string>('search', { required: true })

const filteredOptions = computed(
  () => props.options.filter(item => item.toLowerCase().includes(searchModel.value.toLowerCase())),
)

const modelValue = defineModel<string>({ required: true })
const target = ref<HTMLDivElement | null>(null)
const isOpen = ref(false)

function toggleDropdown() {
  isOpen.value = !isOpen.value
}

function selectOption(option: string) {
  onClose()
  modelValue.value = option
  searchModel.value = option
}

function onClose() {
  searchModel.value = modelValue.value
  isOpen.value = false
}

onClickOutside(target, onClose)

onUnmounted(() => onClose)
onMounted(() => searchModel.value = modelValue.value)
</script>

<style scoped>
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.3s;
}
.fade-enter, .fade-leave-to {
  opacity: 0;
}
</style>
