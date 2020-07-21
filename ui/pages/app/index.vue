<template>
  <div>
    <div>{{users}}</div>
    <div v-if="$auth.loggedIn">{{ $auth.user.email }}</div>
    <div v-else>Hello Guest</div>
  </div>
</template>

<script>
import Logo from "~/components/Logo.vue";
import VuetifyLogo from "~/components/VuetifyLogo.vue";

export default {
  components: {
    Logo,
    VuetifyLogo
  },
  data() {
    return {
      users: []
    }
  },
  methods: {
  },
  computed: {
  },
  async fetch() {
    // console.log("Context: " + JSON.parse(context))
    var response = await this.$axios.get("/user");
    console.log("User List: " + JSON.stringify(response.data.users));
    this.users = response.data.users;
  }
};
</script>
