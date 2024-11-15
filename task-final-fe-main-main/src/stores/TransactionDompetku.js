import { defineStore } from 'pinia';

export const useTransactionDompetkuStore = defineStore('transactionDompetku', {
  state: () => ({
    // Data transaksi disimpan dalam array
    records: []
  }),

  actions: {
    // Memuat data dari localStorage ketika store diinisialisasi
    loadData() {
      const savedData = localStorage.getItem('transactions');
      if (savedData) {
        this.records = JSON.parse(savedData);
      }
    },

    // Menambahkan transaksi baru dan menyimpan ke localStorage
    addRecord(description, amount, category) {
      this.records.push({ description, amount, category });
      this.saveData();
    },

    // Menghapus transaksi berdasarkan index dan memperbarui localStorage
    deleteRecord(index) {
      this.records.splice(index, 1);
      this.saveData();
    },

    // Menghapus semua transaksi dan memperbarui localStorage
    clearAll() {
      this.records = [];
      this.saveData();
    },

    // Menyimpan data transaksi ke localStorage
    saveData() {
      localStorage.setItem('transactions', JSON.stringify(this.records));
    }
  }
});
