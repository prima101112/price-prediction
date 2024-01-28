<template>
    <div>
      <h1>Historicalpo Data</h1>
      <div>
        sku_id : <InputText type="text" v-model="sku_id_form" style="margin-right: 5px;" /><Button @click="suggesBySKU" style="margin-right: 5px;"> suggest me </Button> | using default qty 10
      </div>
      <div class="result">
        <h2>Result {{ sku_id_form }}</h2>
        <p>lowest price : {{ lowest_price }} | higest price : {{ higest_price }} | linear regression price {{ lr_price }}</p>
        <p>for selling this item we could use middle price between avg lr and median {{ middle_price }} with profit per unit {{ middle_lowest }},</p>
        <p>you could adjust based on customer connection and qty of the unit</p>
      </div>
      <br>
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
  import InputText from 'primevue/inputtext';
  import Button from 'primevue/button';

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

    suggesBySKU() {
      this.$router.push({
        name: 'suggestion',
        query: {
          sku_id: this.sku_id_form,
        },
      });
    },
    navigateToHistoricalPO(skuId) {
    this.$router.push({
        name: 'suggestion',
        query: {
          sku_id: skuId,
        },
      });
    },
    formatCurrency(amount) {
    // Convert the amount to a fixed number with 2 decimal places
    const fixedAmount = amount.toFixed(2);

    // Use toLocaleString() to add comma separators and format as currency
    const currencyFormatted = parseFloat(fixedAmount).toLocaleString('id-ID', {
        style: 'currency',
        currency: 'IDR'
    });

    return currencyFormatted;
},
    findMiddleValue(a, b, c) {
      // Find the middle value by comparing the three values
      if ((a <= b && b <= c) || (c <= b && b <= a)) {
          return b;
      } else if ((b <= a && a <= c) || (c <= a && a <= b)) {
          return a;
      } else {
          return c;
      }
    },
    async fetchhistoricalpoData(queryParams) {
      this.sku_id = queryParams.sku_id;
      this.sku_id_form = queryParams.sku_id;
      this.loading = true;
      try {
        const response = await axios.get('http://localhost:8080/suggestion', {
          params: queryParams,
        });
        console.log(response.data);
        // Format unit selling price in the fetched data
        this.hpos = response.data.historical_po_data.map(item => ({
          ...item,
          unit_selling_price_formatted: `Rp ${item.unit_selling_price}`
        }));
        const mid = this.findMiddleValue(response.data.linear_regression_price, response.data.median_price, response.data.average_price)
        this.lowest_price = this.formatCurrency(response.data.lowest_price);
        this.higest_price = this.formatCurrency(response.data.higest_price);
        this.lr_price = this.formatCurrency(response.data.linear_regression_price);
        this.middle_price = this.formatCurrency(mid);
        this.middle_lowest = this.formatCurrency(mid - response.data.lowest_price);
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
  
  