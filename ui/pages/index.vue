<template>
  <div>
    This is Home Page for unauthorized users
  </div>
</template>

<script>
import Logo from '~/components/Logo.vue'
import VuetifyLogo from '~/components/VuetifyLogo.vue'

export default {
  layout: 'unauthz',
  components: {
    Logo,
    VuetifyLogo
  },
}
</script>
