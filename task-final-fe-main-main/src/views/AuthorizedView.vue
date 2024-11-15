<template>
  <v-container>
    <v-row>
      <v-col cols="12" md="6">
        <v-text-field
          v-model="dateRange.from"
          label="From Date"
          type="date"
          outlined
          dense
        />
      </v-col>
      <v-col cols="12" md="6">
        <v-text-field
          v-model="dateRange.to"
          label="To Date"
          type="date"
          outlined
          dense
        />
      </v-col>
    </v-row>

    <v-btn @click="fetchTransactions" color="primary" class="mt-3">Fetch Transactions</v-btn>

    <v-divider class="my-4"></v-divider>

    <v-data-table
      v-if="transactions.length"
      :headers="headers"
      :items="transactions"
      item-value="date"
      class="elevation-1"
    >
      <!-- Use full v-slot syntax -->
      <template v-slot:[`item.date`]="{ item }">
        <span>{{ new Date(item.date).toLocaleDateString() }}</span>
      </template>
      <template v-slot:[`item.amount`]="{ item }">
        <span>{{ item.amount.toLocaleString() }}</span>
      </template>
    </v-data-table>

    <v-alert v-else type="info">No transactions found for the selected date range.</v-alert>
  </v-container>
</template>

<script setup>
import { ref } from 'vue'
import { useUserStore } from '@/stores/user.js'

const userStore = useUserStore()
const dateRange = ref({
  from: '',
  to: ''
})

const transactions = ref([])

const headers = [
  { text: 'Date', align: 'start', key: 'date', value: 'date' },
  { text: 'Description', value: 'description' },
  { text: 'Amount', value: 'amount' },
]

async function fetchTransactions() {
  // Mock fetching data from an API or database
  // You would replace this with actual API call to fetch transaction history
  const response = await fetch('/account/mutation', {
    method: 'POST',
    body: JSON.stringify(dateRange.value),
    headers: {
      'Content-Type': 'application/json',
      'Authorization': userStore.token,
    },
  })

  if (response.ok) {
    const data = await response.json()
    transactions.value = data.transactions
  } else {
    transactions.value = []
  }
}
</script>
