import defaultColors from 'tailwindcss/colors'

/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{js,ts,vue}'],
  safelist: ['secondary-disabled', 'secondary-enabled'],
  theme: {
    colors: {
      'bg-primary': 'var(--bg-color)',
      'primary': 'var(--primary-color)',

      'secondary': 'var(--secondary-color)',
      'secondary-disabled': 'var(--secondary-disabled-color)',
      'secondary-enabled': 'var(--secondary-enabled-color)',

      'surface': {
        dark: 'var(--surface-dark)',
        primary: 'var(--surface-primary)',
        secondary: 'var(--surface-secondary)',
        tertiary: 'var(--surface-tertiary)',
      },

      'text-alert': 'var(--text-alert)',
      'placeholder-color': 'var(--placeholder-color)',

      ...defaultColors,
    },
  },
  plugins: [],
}
