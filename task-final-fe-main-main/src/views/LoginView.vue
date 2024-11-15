<script setup>
import { ref } from 'vue'
import { mdiAccountOutline, mdiLockOutline, mdiEye, mdiEyeOff } from '@mdi/js'
import logo from '@/assets/celenganku4.png'
import { useUserStore } from '@/stores/user.js'
import { useRouter } from 'vue-router'

const router = useRouter()
const userStore = useUserStore()

const username = ref('')
const password = ref('')
const submitErr = ref('') // error message

const showPassword = ref(false)
const submitInProgress = ref(false) // submit loading state

async function submit() {
  submitInProgress.value = true
  submitErr.value = ''

  try {
    await userStore.login({
      username: username.value,
      password: password.value,
    })
    // success
    router.push({ name: 'home' })
  } catch (err) {
    if (err.response) {
      submitErr.value = err.response.data.error
    } else {
      submitErr.value = err
    }
  }

  submitInProgress.value = false
}
</script>

<template>
  <div class="login-container d-flex justify-center align-center h-100">
    <v-card class="login-card pa-8" elevation="10" min-width="448" rounded="lg">
      <v-img class="mx-auto my-6" max-width="200" :src="logo" contain></v-img>

      <div class="text-center mb-5">
        <div class="text-subtitle-1 font-weight-bold">Welcome to Celenganku</div>
        <div class="text-caption text-medium-emphasis">Please log in to continue</div>
      </div>

      <v-form>
        <!-- Username Input -->
        <v-text-field
          v-model="username"
          density="compact"
          placeholder="Username"
          :prepend-inner-icon="mdiAccountOutline"
          variant="outlined"
          class="mb-4"
        ></v-text-field>

        <!-- Password Input -->
        <v-text-field
          v-model="password"
          :append-inner-icon="showPassword ? mdiEyeOff : mdiEye"
          :type="showPassword ? 'text' : 'password'"
          density="compact"
          autoComplete="true"
          placeholder="Enter your password"
          :prepend-inner-icon="mdiLockOutline"
          variant="outlined"
          @click:append-inner="showPassword = !showPassword"
          class="mb-4"
        ></v-text-field>

        <!-- Error Message -->
        <v-card class="mb-5" color="error" variant="tonal" v-show="submitErr">
          <v-card-text class="text-medium-emphasis text-caption text-center">
            {{ submitErr }}
          </v-card-text>
        </v-card>

        <!-- Submit Button -->
        <v-btn
          color="primary"
          size="large"
          variant="tonal"
          block
          @click="submit"
          :readonly="submitInProgress"
          class="mb-5"
        >
          Login
        </v-btn>

        <v-card class="mb-5" color="surface-variant" variant="tonal">
          <!-- <v-card-text class="text-medium-emphasis text-caption text-center">
            "Thou shalt not pass!"
          </v-card-text> -->
        </v-card>
      </v-form>
    </v-card>
  </div>
</template>

<style scoped>
.login-container {
  background: linear-gradient(135deg, #c3dfff, #bad1ff);
  height: 100vh;
  padding: 20px;
}

.login-card {
  background-color: white;
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
}

.v-text-field,
.v-btn {
  border-radius: 12px;
}

.v-text-field input {
  font-size: 1rem;
}

.v-btn {
  font-size: 1.1rem;
  font-weight: bold;
}

.v-img {
  border-radius: 12px;
}

.text-center {
  text-align: center;
}
</style>
