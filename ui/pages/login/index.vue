<template>
  <v-row align="center" justify="center">
    <v-col cols="12" sm="8" md="4">
      <v-card class="elevation-12">
        <v-toolbar color="primary" dark flat>
          <v-toolbar-title>Login form</v-toolbar-title>
          <v-spacer></v-spacer>
        </v-toolbar>
        <v-card-text>
          <v-form >
            <v-text-field
              label="Login"
              name="login"
              v-model="userName"
              prepend-icon="mdi-account"
              type="text"
            ></v-text-field>

            <v-text-field
              id="password"
              label="Password"
              name="password"
              prepend-icon="mdi-lock"
              type="password"
              v-model="password"
            ></v-text-field>
          </v-form>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="primary" @click="login">Login</v-btn>
        </v-card-actions>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
export default {
  data() {
    return {
      userName: "defaultUserName",
      password: "defaultPassword"
    };
  },
  methods: {
    login: function() {
      console.log("Thi sis the loginmethod");
      console.log("User Name: " + this.userName);
      console.log("Password: " + this.password);
      this.$store.dispatch('auth/login', {"userName": this.userName, "password": this.password})
    }
  }
};
</script>