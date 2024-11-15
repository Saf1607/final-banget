<script>
import { createAccount, updateAccount } from '@/services/accountService'; // Import service untuk create dan update

export default {
  name: 'AccountForm',
  data() {
    return {
      accountId: null,  // ID akun (untuk edit)
      name: '',         // Nama akun
      balance: 0,       // Saldo akun
      isEditMode: false, // Menandakan apakah sedang dalam mode edit
    };
  },
  async created() {
    // Jika route memiliki parameter 'accountId', maka ini untuk edit akun
    const accountId = this.$route.params.accountId;
    if (accountId) {
      this.isEditMode = true;
      this.accountId = accountId;
      await this.loadAccountData(accountId);
    }
  },
  methods: {
    async loadAccountData(accountId) {
      // Ambil data akun berdasarkan accountId
      try {
        const account = await this.getAccount(accountId);
        this.name = account.name;
        this.balance = account.balance;
      } catch (error) {
        console.error("Failed to load account data:", error);
      }
    },
    async handleSave() {
      // Menyimpan data baru atau mengupdate akun yang ada
      const accountData = {
        name: this.name,
        balance: this.balance,
      };

      try {
        if (this.isEditMode) {
          await updateAccount(this.accountId, accountData);  // Mengupdate akun
        } else {
          await createAccount(accountData);  // Membuat akun baru
        }
        this.$router.push('/');  // Kembali ke halaman sebelumnya (misalnya halaman daftar akun)
        alert('Akun berhasil disimpan!');
      } catch (error) {
        console.error("Error saving account:", error);
        alert('Gagal menyimpan akun!');
      }
    },
    // Fungsi untuk memformat saldo dengan format IDR
    formatBalance(balance) {
      return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(balance);
    },
    // Fungsi untuk mendapatkan akun berdasarkan ID (misalnya melalui API)
    async getAccount(accountId) {
      // Contoh: panggil API atau service untuk mengambil data akun
      const response = await fetch(`http://localhost:8080/account/${accountId}`);
      return response.json();
    }
  }
};
</script>

<template>
  <div class="account-form-container">
    <h2 class="title">{{ isEditMode ? 'Edit Akun' : 'Buat Akun Baru' }}</h2>

    <v-card class="account-form-card">
      <v-form @submit.prevent="handleSave">
        <v-text-field
          v-model="name"
          label="Nama Akun"
          required
          :disabled="isEditMode"
        ></v-text-field>
        
        <v-text-field
          v-model="balance"
          label="Saldo Awal"
          required
          type="number"
          min="0"
        ></v-text-field>
        
        <v-btn color="primary" @click="handleSave" :disabled="!name || balance <= 0">
          Simpan
        </v-btn>
      </v-form>
    </v-card>
  </div>
</template>

<style scoped>
.account-form-container {
  padding: 20px;
  max-width: 600px;
  margin: 0 auto;
  background-color: #f9f9f9;
  border-radius: 10px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.title {
  font-size: 1.5rem;
  font-weight: 600;
  color: #0e185d;
  margin-bottom: 20px;
  text-align: center;
}

.account-form-card {
  padding: 20px;
  background-color: #ffffff;
  border-radius: 10px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.v-text-field {
  margin-bottom: 20px;
}

.v-btn {
  display: block;
  width: 100%;
}
</style>
