<!-- src/components/BillingPage.vue -->
<template>
  <div class="billing-page">
    <h2>Bayar Tagihan</h2>

    <!-- Pilihan Jenis Tagihan -->
    <div class="form-group">
      <label for="biller">Pilih Jenis Tagihan:</label>
      <select v-model="selectedBiller" @change="resetBillerInfo">
        <option disabled value="">-- Pilih Jenis Tagihan --</option>
        <option v-for="biller in billers" :key="biller.biller_id" :value="biller.biller_id">
          {{ biller.name }}
        </option>
      </select>
    </div>

    <!-- Input Nomor Akun Tagihan -->
    <div class="form-group">
      <label for="account">Masukkan Nomor Akun:</label>
      <input
        type="text"
        v-model="billingAccount"
        placeholder="Contoh: 123456789"
        @input="resetBillerInfo"
      />
    </div>

    <!-- Tombol Cek Tagihan -->
    <button :disabled="!selectedBiller || !billingAccount" @click="checkBiller">Cek Tagihan</button>

    <!-- Tampilkan Detail Tagihan setelah cek -->
    <div v-if="billerInfo" class="biller-info">
      <h3 class="center-text">Detail Tagihan:</h3>
      <p><strong>Nama:</strong> {{ billerInfo.name }}</p>
      <p><strong>Total Tagihan:</strong> Rp {{ billerInfo.bill_amount }}</p>

      <!-- Input Jumlah Pembayaran -->
      <div class="form-group">
        <label for="amount">Jumlah Pembayaran:</label>
        <input
          type="number"
          v-model="paymentAmount"
          :min="billerInfo.bill_amount"
          placeholder="Masukkan jumlah pembayaran"
        />
      </div>

      <!-- Tombol Bayar -->
      <button :disabled="!isPaymentValid" @click="payBiller">Bayar</button>
    </div>

    <!-- Pesan Status Pembayaran -->
    <div v-if="paymentStatus" class="payment-status">
      <p>{{ paymentStatus }}</p>
    </div>

    <!-- Pesan Error Cek Tagihan -->
    <div v-if="billerError" class="error-message">
      <p>{{ billerError }}</p>
    </div>
  </div>
</template>

<script>
import { getBillerList, checkBillerAccount, payBillerAccount } from '../services/billerService';

export default {
  data() {
    return {
      billers: [],           // Menyimpan daftar biller
      selectedBiller: '',     // Menyimpan jenis tagihan yang dipilih
      billingAccount: '',     // Menyimpan nomor akun yang diinput
      billerInfo: null,       // Menyimpan informasi tagihan setelah cek
      paymentAmount: 0,       // Menyimpan jumlah pembayaran
      paymentStatus: '',      // Menyimpan status hasil pembayaran
      billerError: '',        // Menyimpan pesan error jika terjadi masalah saat cek biller
    };
  },
  computed: {
    isPaymentValid() {
      // Validasi bahwa jumlah pembayaran tidak kurang dari tagihan
      return this.paymentAmount >= (this.billerInfo ? this.billerInfo.bill_amount : 0);
    }
  },
  mounted() {
    this.loadBillers();
  },
  methods: {
    async loadBillers() {
      try {
        const response = await getBillerList();
        this.billers = response.data;
      } catch (error) {
        console.error("Gagal memuat daftar biller:", error);
        this.billerError = "Gagal memuat daftar biller.";
      }
    },
    async checkBiller() {
      this.billerError = ''; // Reset error sebelum cek
      try {
        const response = await checkBillerAccount(this.selectedBiller, this.billingAccount);
        this.billerInfo = response.data;
        this.paymentAmount = this.billerInfo.bill_amount; // Set nilai awal jumlah pembayaran
      } catch (error) {
        console.error("Gagal memeriksa akun biller:", error);
        this.billerInfo = null; // Reset jika terjadi error
        this.billerError = "Nomor akun tidak ditemukan atau terjadi kesalahan.";
      }
    },
    async payBiller() {
  if (!this.isPaymentValid) {
    alert("Jumlah yang dibayar tidak boleh kurang dari total tagihan.");
    return;
  }
  try {
    const payload = {
      biller_id: this.selectedBiller,
      biller_account_id: this.billingAccount,
      amount: this.paymentAmount
    };
    const response = await payBillerAccount(payload);
    this.paymentStatus = `Pembayaran berhasil: ${response.data.message}`;
    
    // Menampilkan alert dan me-refresh halaman setelah berhasil
    alert("Pembayaran berhasil");
    window.location.reload();

    this.resetForm();
  } catch (error) {
    console.error("Gagal melakukan pembayaran:", error);
    this.paymentStatus = "Pembayaran gagal. Silakan coba lagi.";
  }
},
resetForm() {
  this.selectedBiller = '';
  this.billingAccount = '';
  this.billerInfo = null;
  this.paymentAmount = 0;
  this.paymentStatus = '';
  this.billerError = '';
}

  }
};
</script>

<style scoped>
.billing-page {
  max-width: 1240px;
  margin: 0 auto;
  padding: 20px;
  background-color: #f9f9f9;
  border: 1px solid #ddd;
  border-radius: 8px;
}

h2 {
  text-align: center;
  color: #333;
}

.form-group {
  margin-bottom: 15px;
}

label {
  display: block;
  font-weight: bold;
  margin-bottom: 5px;
}

input, select {
  width: 100%;
  padding: 8px;
  box-sizing: border-box;
  border-radius: 4px;
  border: 1px solid #ccc;
}

button {
  background-color: #28a745;
  color: white;
  padding: 10px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  width: 100%;
  margin-top: 10px;
}

button:disabled {
  background-color: #ccc;
}

.biller-info {
  margin-top: 25px;
  padding: 15px;
  background-color: #e9ecef;
  border-radius: 4px;
}

.center-text {
  text-align: center;  /* Hanya center-kan text di h3 */
  color: #333;
  margin-bottom: 10px;
  font-size: 1.4em;
}

.biller-info p {
  font-size: 1em;
  margin: 5px 0;
}

.biller-info p strong {
  font-weight: bold;
}

.payment-status {
  margin-top: 20px;
  font-weight: bold;
  text-align: center;
  color: #28a745;
}

.error-message {
  margin-top: 20px;
  font-weight: bold;
  text-align: center;
  color: red;
}
</style>
