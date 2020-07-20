<template>
  <v-row align="center" justify="center">
    <v-col cols="12" sm="8" md="4">
      <v-card class="elevation-12">
        <v-toolbar color="primary" dark flat>
          <v-toolbar-title>Register</v-toolbar-title>
          <v-spacer></v-spacer>
        </v-toolbar>
        <v-card-text>
          <v-form>
            <v-text-field
              label="Name"
              name="name"
              v-model="name"
              prepend-icon="mdi-face"
              type="text"
            ></v-text-field>
            <v-text-field
              label="Email"
              name="email"
              v-model="email"
              prepend-icon="mdi-email"
              type="text"
            ></v-text-field>
            <v-text-field
              label="User Name"
              name="userName"
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
          <v-btn color="primary" @click="register">Register</v-btn>
        </v-card-actions>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
export default {
  auth: false,
  data() {
    return {
      name: "Alwin Doss",
      email: "name@email.com",
      userName: "alwindoss",
      password: "password"
    };
  },
  methods: {
    async register() {
      await this.$axios.post("/user", {
        name: this.name,
        userName: this.userName,
        email: this.email,
        password: this.password
      });
      this.$auth.loginWith("local", {
        data: {
          userName: this.userName,
          password: this.password
        }
      });
      this.$router.push("/")
    }
  }
};
</script>