/**
 * plugins/vuetify.js
 *
 * Framework documentation: https://vuetifyjs.com`
 */

// Styles
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'

// Composables
import { createVuetify } from 'vuetify'
import colors from 'vuetify/util/colors'

// https://vuetifyjs.com/en/introduction/why-vuetify/#feature-guides
export default createVuetify({
  theme: {
    variations: {
      colors: ['primary', 'secondary', 'accent', 'error', 'info', 'success', 'warning', 'surface'],
      lighten: 4,
      darken: 4,
    },
    defaultTheme: 'dark',
    light: {
      colors: {
        primary: colors.grey.base,
        secondary: colors.grey.lighten2,
        background: colors.grey.darken4,
        surface: colors.grey.darken2,
        error: colors.red.base,
        admiral: '#1E52E6',
        commander: '#308CA7',
        lieutenant: '#24AD32',
        specialist: '#DA5C5C',
        technician: '#E69737',
        member: '#FFC900',
        recruit: '#1CFAC0',
        guest: '#929292',
        ally: '#F87847',
      },
    },
  },
})
