<script setup>
import { RouterLink, RouterView } from 'vue-router'
import HelloWorld from './components/HelloWorld.vue'
import Button from 'primevue/button';
import 'primevue/resources/themes/aura-dark-green/theme.css'


</script>

<template>
  <header>
    <div class="wrapper">
      <HelloWorld msg="You did it!" />
      <nav>
        <RouterLink title="Home" to="/"> <Button class="button" label="Home"></Button></RouterLink>
        <RouterLink title="About" to="/about"><Button class="button" label="About"></Button></RouterLink>
        <RouterLink title="Customers" to="/customers"><Button class="button" label="Customers"></Button></RouterLink>
        <RouterLink title="Historicalpo" to="/historicalpo"><Button class="button" label="Historicalpo"></Button></RouterLink>
      </nav>
      <br />
      <hr />
    </div>
  </header>


  <!-- content is here -->
  <RouterView />


</template>

<style scoped>
  button {
    margin: 0 0.5rem;
  }
</style>

