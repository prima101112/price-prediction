<template>
    <div>
      <h1>Historicalpo Data</h1>
      <div v-if="loading">Loading...</div>
      <div v-else>
        <div v-if="error">{{ error }}</div>
        <div v-else>
          <DataTable :value="hpos" tableStyle="min-width: 50rem">
            <Column field="id" sortable header="ID"></Column>
            <Column field="sku_id" sortable header="SKUID">
              <template #body="{ data }">
                <button @click="navigateToHistoricalPO(data.sku_id)">{{ data.sku_id }}</button>
              </template>
            </Column>
            <Column field="unit_selling_price_formatted" sortable header="Unit Selling Price" ></Column>
          </DataTable>
        </div>
      </div>
    </div>
  </template>

<script setup>
  import DataTable from 'primevue/datatable';
  import Column from 'primevue/column';
</script>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      loading: false,
      error: '',
      hpos: [],
    };
  },
  setup() {
    console.log('Setup');
  },
  mounted() {
    // Get query parameters when the component is mounted
    const queryParams = this.$route.query;
    console.log('Query Parameters:', queryParams);

    // Fetch data using query parameters
    this.fetchhistoricalpoData(queryParams);
  },
  watch: {
    // Watch for changes in the route query parameters
    '$route.query': {
      handler: function (newQueryParams) {
        console.log('Query Parameters:', newQueryParams);
        this.fetchhistoricalpoData(newQueryParams);
      },
      deep: true,
    },
  },
  methods: {
      navigateToHistoricalPO(skuId) {
      this.$router.push({
        name: 'historicalpo',
        query: {
          sku_id: skuId,
        },
      });
    },
    async fetchhistoricalpoData(queryParams) {
      this.loading = true;
      try {
        const response = await axios.get('http://localhost:8080/historicalpo', {
          params: queryParams,
        });
        // Format unit selling price in the fetched data
        this.hpos = response.data.map(item => ({
          ...item,
          unit_selling_price_formatted: `Rp ${item.unit_selling_price}`
        }));
      } catch (error) {
        this.error = 'Error fetching historicalpo data';
        console.error(error);
      } finally {
        this.loading = false;
      }
    },
  },
};


</script>
  
  