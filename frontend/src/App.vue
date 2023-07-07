<script setup>
import { RouterLink, RouterView } from 'vue-router'
import Header from './components/Header.vue'
</script>

<script>
export default {
  data() {
    return {
      authenticated: false,
    }
  },
  methods: {
    checkAuthentication: function () {
      this.$axios.get('/authenticated')
          .then(response => {
            if (response.data.authenticated !== true) {
              this.$router.push('/login');
            }
            this.authenticated = response.data.authenticated;
          })
          .catch(error => {
            this.$router.push('/login');
            this.authenticated = false;
          });
    }
  },
  mounted() {
    this.checkAuthentication();
  }
}
</script>

<template>
  <Header v-if="authenticated"></Header>

  <RouterView v-on:login-succeeded="checkAuthentication"/>
</template>

<style scoped>
</style>
