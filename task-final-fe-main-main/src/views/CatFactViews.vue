<script setup>
import { ref } from 'vue'
import axios from 'axios' // Import axios for API calls

// Define table headers
const headers = [
  { title: 'Fact', key: 'fact', sortable: false },
  { title: 'Actions', key: 'actions', sortable: false }
]

// Define a reactive array to store facts
const facts = ref([])

// Function to load data from the API
async function loadMore() {
  try {
    const response = await axios.get("http://localhost:8080/auth/login")
    if (response.data) {
      facts.value.push({ fact: response.data.fact })
    }
  } catch (error) {
    console.error("Failed to load data:", error)
  }
}

const loadMore = async () => {
  try {
    const response = await axios.get('http://localhost:8080/login')
    console.log(response.data) // Periksa data di console
    facts.value.push(...response.data.facts) // Sesuaikan dengan struktur respons
  } catch (error) {
    console.error('Error loading facts:', error)
  }
}

</script>



<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <div class="text-center">
          <!-- Button to load more data -->
          <v-btn @click="loadMore" color="primary">Load More</v-btn>
        </div>
      </v-col>

      <v-col cols="12">
        <!-- Data table to display facts -->
        <v-data-table :headers="headers" :items="facts" :items-per-page="5">
          <!-- Slot for actions (optional) -->
          <template v-slot:[item.actions]="{ item }">
            <v-btn color="red" @click="facts.splice(facts.indexOf(item), 1)">
              Delete
            </v-btn>
          </template>
        </v-data-table>
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped>
.text-center {
  text-align: center;
  margin: 20px 0;
}
</style>