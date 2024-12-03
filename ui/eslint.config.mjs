import antfu from '@antfu/eslint-config'

export default antfu({
  rules: {
    'vue/block-order': [
      'error',
      {
        order: ['template', 'script', 'style'],
      },
    ],
    'max-len': [
      'warn',
      {
        code: 120,
        comments: 200,
        tabWidth: 2,
      },
    ],
  },
})
