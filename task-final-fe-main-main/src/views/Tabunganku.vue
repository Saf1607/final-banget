<script setup>
import { ref, computed, onMounted } from 'vue';
import { useTransactionDompetkuStore } from '@/stores/TransactionDompetku';

const store = useTransactionDompetkuStore();

// State untuk form input
const description = ref('');
const amount = ref(null);
const category = ref('');
const categories = ['Food', 'Transport', 'Health','Others'];


// Memuat data dari localStorage saat komponen dimuat
onMounted(() => {
  store.loadData();
});

// Computed properties untuk perhitungan total pemasukan, pengeluaran, dan saldo
const totalIncome = computed(() =>
  store.records
//   Ini mengambil hanya transaksi dengan nilai amount positif (uang yang masuk).
    .filter(record => record.amount > 0)
    .reduce((total, record) => total + record.amount, 0)
);

const totalExpense = computed(() =>
  store.records
    .filter(record => record.amount < 0)
    .reduce((total, record) => total + record.amount, 0)
);

const totalBalance = computed(() => totalIncome.value + totalExpense.value);

// Fungsi untuk menambahkan record
const addRecord = () => {
    // memanggil metode addRecord dari objek store, 
    // yang berfungsi untuk menambah transaksi ke dalam daftar catatan (records).
  store.addRecord(description.value, amount.value, category.value);
  // Reset input setelah menambah transaksi
  description.value = '';
  amount.value = null;
  category.value = '';
};

// Fungsi addRecord ini menyimpan transaksi baru ke dalam store berdasarkan input pengguna 
// dan mengosongkan kolom input setelah transaksi berhasil ditambahkan.

// Fungsi untuk menghapus record berdasarkan index
const deleteRecord = (index) => {
  store.deleteRecord(index);
};

// Fungsi untuk menghapus semua record dengan konfirmasi
const confirmClearAll = () => {
  if (confirm("Yakin ingin menghapus semua data?")) {
    store.clearAll();
  }
};
</script>

<template>
  <div class="dompetku">
    <h1>Dompetku</h1>
    <h2>Expense & Income Money Management</h2>

    <!-- Form Input -->
    <form @submit.prevent="addRecord" class="horizontal-form">
        <v-text-field v-model="description" label="Keterangan" outlined required />
      <v-text-field v-model.number="amount" label="Jumlah Uang" type="number" outlined required />

      <!-- Dropdown Kategori -->
      <v-select
        v-model="category"
        :items="categories"
        label="Pilih Kategori"
        outlined
        dense
        required
      />

      <v-btn rounded="xl" size="large" block type="submit">Submit</v-btn>
    </form>



    <!-- Summary Data -->
    <div class="summary">
      <h3>Summary</h3>
      <p>Total Transaksi: {{ store.records.length }}</p>
      <p>Total Pemasukan: Rp{{ totalIncome.toLocaleString() }}</p>
      <p>Total Pengeluaran: Rp{{ totalExpense.toLocaleString() }}</p>
      <p>Total Saldo: Rp{{ totalBalance.toLocaleString() }} ({{ totalBalance >= 0 ? 'Surplus' : 'Minus' }})</p>
    </div>

    <h2>List Transaksi</h2>
    <!-- List Data -->
    <ul>
        <!-- v for ini adalah cara Vue untuk melakukan perulangan. Artinya,  mengulang elemen <li> untuk setiap transaksi yang ada di store.records. 
        store.records adalah kumpulan data transaksi yang disimpan. -->
      <li v-for="(record, index) in store.records" :key="index">
        <span>{{ record.category }} | {{ record.description }} - Rp{{ record.amount.toLocaleString() }}</span>
        <button @click="deleteRecord(index)" class="delete-btn">hapus</button>
     
      </li>
    </ul>

    

    <!-- Tombol Hapus Semua -->
    <!-- <button @click="confirmClearAll" class="clear-all">Hapus Semua Data</button> -->
    <v-btn rounded="xl" size="x-large" block button @click="confirmClearAll"class="clear-all">Hapus Semua Data</v-btn>
  </div>
</template>

<style scoped>
.dompetku {
  max-width: 1000px;
  margin: 1em auto;
  padding: 1em;
  border: 1.5px solid #ddd;
  border-radius: 8px;
  text-align: center;
}

form {
  display: flex;
  align-items: flex-start;
  gap: 1em;
  flex-wrap: wrap;
  margin-bottom: 1em;
  margin-top: 2em;
}

.horizontal-form > * {
  flex: 1;
  min-width: 200px;
}

input, select {
  margin: 0.5em 0;
  padding: 0.5em;
  width: 100%;
  box-sizing: border-box;
}

button {
  margin-top: 0.5em;
  padding: 0.5em;
  width: 100%;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.summary {
  margin-top: 1em;
  padding: 1em;
  border: 1px solid #ddd;
  border-radius: 8px;
  background-color: #f9f9f9;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  margin: 0.5em 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.5em;
  border: 1px solid #ddd;
  border-radius: 4px;
  background-color: #f9f9f9;
}

.delete-btn {
  padding: 0.3em 0.5em;
  background-color: #f44336;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.9rem;
  width: auto;
}

.clear-all {
  margin-top: 1em;
  background-color: #f44336;
  color: white;
  border: none;
  border-radius: 4px;
  padding: 0.7em 1em;
  cursor: pointer;
  font-size: 1rem;
}

.clear-all:hover, .delete-btn:hover {
  background-color: #d32f2f;
}
</style>
