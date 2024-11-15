<script setup>
import { onMounted } from 'vue';
import { useUserStore } from '@/stores/userStore';

const userStore = useUserStore();

onMounted(() => {
  if (!userStore.user) {
    userStore.fetchUser(); // Pastikan data pengguna diambil saat komponen dimuat
  }
});
</script>

<template>
  <div class="home-container">
    <header class="home-header">
      <h1>Selamat Datang, {{ userStore.user?.name }}</h1>
      <p>Berikut adalah informasi akun Anda:</p>
    </header>

    <section class="account-info">
      <div class="account-detail">
        <label>Nama Akun:</label>
        <span>{{ userStore.user?.name }}</span>
      </div>
      <div class="account-detail">
        <label>Saldo:</label>
        <span>Rp {{ userStore.user?.balance.toLocaleString('id-ID') }}</span>
      </div>
    </section>


  </div>
</template>

<style scoped>
.home-container {
  max-width: 800px; /* Lebar kontainer sedikit lebih besar untuk tampilan yang lebih luas */
  margin: 0 auto;
  padding: 20px;
  background-color: #fff;
  border-radius: 10px;
  box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);
}

.home-header {
  margin-bottom: 30px;
  text-align: center;
}

.home-header h1 {
  font-size: 2em;
  color: #333;
  font-weight: bold;
}

.home-header p {
  font-size: 1.1em;
  color: #666;
}

.account-info {
  margin-top: 20px;
}

.account-detail {
  display: flex;
  justify-content: space-between;
  padding: 10px 0;
  border-bottom: 1px solid #eee;
}

.account-detail label {
  font-weight: 600;
  color: #444;
}

.account-detail span {
  color: #333;
  font-weight: 500;
}

.logout-button {
  margin-top: 30px;
  padding: 12px 25px;
  background-color: #ff5a5a;
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  transition: background-color 0.3s;
  width: 100%; /* Tombol logout menjadi lebar penuh */
}

.logout-button:hover {
  background-color: #e04a4a;
}

/* Responsif untuk layar kecil */
@media (max-width: 600px) {
  .home-container {
    padding: 15px;
  }

  .home-header h1 {
    font-size: 1.5em;
  }

  .account-info {
    margin-top: 15px;
  }

  .account-detail {
    flex-direction: column;
    align-items: flex-start;
  }

  .account-detail label {
    margin-bottom: 5px;
  }
}
</style>
