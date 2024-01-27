<template>
    <div>
      <h1>Customers Data</h1>
      <div v-if="loading">Loading...</div>
      <div v-else>
        <div v-if="error">{{ error }}</div>
        <div v-else>
          <div v-for="customer in customers" :key="customer.id">
            <h2>{{ customer.customer_id }}</h2>
            <p>Adress: {{ customer.address }}</p>
            <p>State: {{ customer.state }}</p>
            <!-- Add other customer data fields as needed -->
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        loading: false,
        error: '',
        customers: [],
      };
    },
    mounted() {
      this.fetchcustomerData();
    },
    methods: {
      async fetchcustomerData() {
        this.loading = true;
        try {
          const response = await axios.get('http://localhost:8080/customers');
          this.customers = response.data;
        } catch (error) {
          this.error = 'Error fetching customer data';
          console.error(error);
        } finally {
          this.loading = false;
        }
      },
    },
  };
  </script>