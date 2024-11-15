<!-- src/views/HomeView.vue -->
<template>
    <div>
      <!-- <h1>Welcome to Home</h1> -->
      <AccountListView /> <!-- Tampilkan komponen AccountListView di sini -->
    </div>
  </template>
  
  <script>
  import AccountListView from './AccountListView.vue';
  
  export default {
    name: 'AccountView',
    components: {
      AccountListView
    }
  };
  </script>