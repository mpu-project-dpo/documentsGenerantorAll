import { defineStore } from 'pinia'
import { reactive, ref } from 'vue'
import { useApiStore } from './apiStore.ts'

export const useFormStore = defineStore('form', () => {
  const confPolicy = ref(false)
  const apiStore = useApiStore()

  const education = reactive({
    university: { field: 'Университет', value: '' },
    group: { field: 'Группа', value: '' },
    dpo: { field: 'ДПО', value: '' },
  })

  const contacts = reactive({
    name: { field: 'Имя', value: '' },
    surname: { field: 'Фамилия', value: '' },
    patronymic: { field: 'Отчетсво', value: '' },
    dateOfBirth: { field: 'Дата рождения', value: '', mask: '##.##.####' },
    phone: { field: 'Номер телефона', value: '', mask: '+7 (###) - ### - ## - ##' },
    email: { field: 'Почта', value: '' },
  })

  const docs = reactive({
    passport: { field: 'Серия и номер паспорта', value: '', mask: 'XXXX ########' },
    dateOfIssue: { field: 'Дата выдачи', value: '', mask: '##.##.####' },
    issuedBy: { field: 'Кем выдан', value: '' },
    snils: { field: 'Снилс', value: '', mask: '###-###-### ##' },
  })

  const findEmptyField = (obj: Record<string, { field: string, value: string }>) => {
    if (!obj)
      return
    return !!Object.values(obj).find(item => !item.value)
  }

  const sendForm = async () => {
    const rawJson = { ...education, ...contacts, ...docs }
    const jsonToSend = Object.entries(rawJson).reduce((acc, [key, item]) => ({ ...acc, [key]: item.value }), {})
    try {
      await apiStore.api.post(import.meta.env.VITE_SEND_FORM_PATH, { ...jsonToSend, confPolicy: confPolicy.value })
    }
    catch (e) {
      console.error(e)
    }
  }

  return {
    confPolicy,
    education,
    contacts,
    docs,
    findEmptyField,
    sendForm,
  }
})
