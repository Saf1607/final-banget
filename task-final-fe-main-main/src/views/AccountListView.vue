<script>
import { getAccountList, topUpBalance } from '@/services/accountService';

export default {
  name: 'AccountListView',
  data() {
    return {
      accounts: [],  // Untuk menyimpan daftar akun
      totalAccounts: 0,  // Untuk menyimpan total akun
      totalBalance: 0,   // Untuk menyimpan total saldo
      averageBalance: 0, // Untuk menyimpan rata-rata saldo
    };
  },
  async created() {
    try {
      // Ambil data akun dari backend menggunakan getAccountList
      this.accounts = await getAccountList();

      // Hitung total akun, total saldo dan rata-rata saldo
      this.calculateSummary();
    } catch (error) {
      console.error("Failed to load accounts:", error);
    }
  },
  methods: {
    formatBalance(balance) {
      // Format saldo menjadi format mata uang IDR
      return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(balance);
    },
    async topUp(accountId) {
      // Meminta jumlah top-up dari pengguna
      const amount = prompt('Masukkan jumlah top-up:');
      if (amount && !isNaN(amount) && parseFloat(amount) > 0) {
        try {
          // Panggil API untuk top-up saldo akun
          const response = await topUpBalance(accountId, parseFloat(amount));

          // Jika berhasil, perbarui saldo akun di frontend
          const accountIndex = this.accounts.findIndex(account => account.account_id === accountId);
          if (accountIndex !== -1) {
            this.accounts[accountIndex].balance += parseFloat(amount);
          }

          // Perbarui informasi total saldo dan rata-rata saldo setelah top-up
          this.calculateSummary();

          alert('Top-up berhasil!');
        } catch (error) {
          console.error('Gagal melakukan top-up:', error);
          alert('Gagal melakukan top-up!');
        }
      } else {
        alert('Jumlah top-up tidak valid!');
      }
    },
    calculateSummary() {
      // Hitung total akun, total saldo dan rata-rata saldo
      this.totalAccounts = this.accounts.length;
      this.totalBalance = this.accounts.reduce((sum, account) => sum + account.balance, 0);
      this.averageBalance = this.totalAccounts > 0 ? this.totalBalance / this.totalAccounts : 0;
    }
  }
};
</script>

<template>
  <div class="account-list-container">
    <h2 class="title">Informasi Akun</h2>

    <!-- Card untuk menampilkan informasi summary -->
    <v-card class="summary-card" outlined>
      <v-card-title class="summary-card-title">
      </v-card-title>
      <v-card-subtitle class="summary-card-content">
        <p>Total Akun: {{ totalAccounts }}</p>
        <p>Total Saldo: {{ formatBalance(totalBalance) }}</p>
        <p>Rata-rata Saldo: {{ formatBalance(averageBalance) }}</p>
      </v-card-subtitle>
    </v-card>

    <!-- Menampilkan daftar akun -->
    <v-card v-if="accounts.length" class="account-list-card" outlined>
      <v-simple-table>
        <thead>
          <tr>
            <th class="column-title">Nama Akun</th>
            <th class="column-title">Saldo</th>
            <th class="column-title">Aksi</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="account in accounts" :key="account.account_id">
            <td class="account-name">{{ account.name }}</td>
            <td class="account-balance">{{ formatBalance(account.balance) }}</td>
            <td>
              <v-btn color="primary" @click="topUp(account.account_id)" class="topup-btn">
                Top-up Saldo
              </v-btn>
            </td>
          </tr>
        </tbody>
      </v-simple-table>
    </v-card>

    <p v-else class="no-accounts-text">Tidak ada akun yang ditemukan.</p>
  </div>
</template>

<style scoped>
/* Styling yang sudah ada tetap digunakan untuk tampilan */
.account-list-container {
  padding: 20px;
  max-width: 1024px;
  margin: 0 auto;
  background-color: #f9f9f9;
  border-radius: 10px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.title {
  font-size: 1.2rem;
  font-weight: 600;
  color: #0e185d;
  margin-bottom: 10px;
  text-align: left;
  padding-left: 10px;
}

/* Styling untuk card ringkasan akun */
.summary-card {
  margin-bottom: 20px;
  padding: 20px;
  background-color: #ffffff;
}

.summary-card-title {
  font-weight: 600;
  color: #0e185d;
}

.summary-card-content p {
  font-size: 1rem;
  font-weight: 500;
  color: #000000;
  margin-bottom: 10px;
}

.v-simple-table {
  width: 100%;
  border-collapse: collapse;
  table-layout: fixed;
}

.column-title {
  font-weight: 600;
  color: #34495e;
  padding-left: 10px;
}

.account-name {
  padding-left: 30px;
}

.account-balance {
  padding-left: 300px;
  padding-right: 300px;
}

.v-simple-table td {
  font-size: 1rem;
  color: #2c3e50;
}

.v-simple-table tr:nth-child(even) {
  background-color: #f9f9f9;
}

.v-simple-table tr:hover {
  background-color: #f1f1f1;
}

.v-btn.topup-btn {
  font-weight: 500;
  padding: 6px 12px;
  border-radius: 5px;
  margin-bottom: 10px;
}

.no-accounts-text {
  font-size: 1.1rem;
  color: #7f8c8d;
  text-align: center;
  margin-top: 20px;
}
</style>
