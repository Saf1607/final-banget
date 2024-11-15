<template>
  <v-container class="edit-profile" fluid>
    <v-row justify="center">
      <v-col cols="12" sm="8" md="6">
        <v-card>
          <v-card-title class="headline">Edit Profile</v-card-title>
          <v-card-text>
            <v-form ref="form">

              <!-- Username Section -->
              <v-card class="section-card">
                <v-card-subtitle>Username</v-card-subtitle>
                <v-text-field
                  label="Username"
                  v-model="username"
                  :rules="[usernameRules.required]"
                  required
                ></v-text-field>
                <v-btn
                  color="primary"
                  @click="handleUsernameSubmit"
                  class="mt-2"
                  :disabled="!username"
                >
                  Save Username
                </v-btn>
              </v-card>

              <v-divider class="my-4"></v-divider> <!-- Divider -->

              <!-- Password Section -->
              <v-card class="section-card">
                <v-card-subtitle>Password</v-card-subtitle>
                <v-text-field
                  label="Password Baru"
                  type="password"
                  v-model="newPassword"
                  :rules="[passwordRules.required, passwordRules.minLength]"
                  required
                ></v-text-field>

                <v-text-field
                  label="Konfirmasi Password Baru"
                  type="password"
                  v-model="confirmPassword"
                  :rules="[passwordRules.required, passwordRules.match]"
                  required
                ></v-text-field>
                <v-btn
                  color="primary"
                  @click="handlePasswordSubmit"
                  class="mt-2"
                  :disabled="newPassword !== confirmPassword || newPassword.length < 6"
                >
                  Save Password
                </v-btn>
              </v-card>
            </v-form>

            <!-- Alert Messages -->
            <v-alert
              v-if="errorMessage"
              type="error"
              class="mt-4"
              :value="true"
            >
              {{ errorMessage }}
            </v-alert>

            <v-alert
              v-if="successMessage"
              type="success"
              class="mt-4"
              :value="true"
            >
              {{ successMessage }}
            </v-alert>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import axios from 'axios';
import { useUserStore } from '@/stores/user';

export default {
  data() {
    return {
      username: '', // New field for username
      newPassword: '',
      confirmPassword: '',
      errorMessage: '',
      successMessage: '',
      usernameRules: {
        required: (value) => !!value || 'Username is required',
      },
      passwordRules: {
        required: (value) => !!value || 'Password is required',
        minLength: (value) => value.length >= 6 || 'Password must be at least 6 characters',
        match: (value) => value === this.newPassword || 'Passwords do not match',
      },
    };
  },
  methods: {
    async handleUsernameSubmit() {
      const userStore = useUserStore();
      const token = userStore.token || localStorage.getItem('token');

      if (!token) {
        this.errorMessage = 'Token tidak ditemukan. Anda perlu login ulang.';
        return;
      }

      try {
        await axios.patch(
          'http://localhost:8080/account/update',
          { name: this.username },
          {
            headers: {
              Authorization: token,
            },
          }
        );
        this.successMessage = 'Username berhasil diperbarui.';
        this.errorMessage = '';
      } catch {
        this.errorMessage = 'Gagal memperbarui username.';
        this.successMessage = '';
      }
    },

    async handlePasswordSubmit() {
      const userStore = useUserStore();
      const token = userStore.token || localStorage.getItem('token');

      if (!token) {
        this.errorMessage = 'Token tidak ditemukan. Anda perlu login ulang.';
        return;
      }

      try {
        await axios.post(
          'http://localhost:8080/auth/change-password',
          { new_password: this.newPassword,
            confirm_new_password: this.confirmPassword
           },
          {
            headers: {
              Authorization: token,
            },
          }
        );
        this.successMessage = 'Password berhasil diperbarui.';
        this.errorMessage = '';
      } catch {
        this.errorMessage = 'Gagal memperbarui password.';
        this.successMessage = '';
      }
    },
  },
};
</script>

<style scoped>
.edit-profile {
  margin-top: 50px;
}
.section-card {
  padding: 20px;
}
.my-4 {
  margin: 16px 0;
}
</style>
