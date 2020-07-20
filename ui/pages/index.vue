<template>
  <div>
    <div v-if="$auth.loggedIn">
      {{ $auth.user.email }}
    </div>
    <div v-else>
      Hello Guest
    </div>
  </div>
</template>

<script>
import Logo from '~/components/Logo.vue'
import VuetifyLogo from '~/components/VuetifyLogo.vue'

export default {
  components: {
    Logo,
    VuetifyLogo
  }
}
</script>
