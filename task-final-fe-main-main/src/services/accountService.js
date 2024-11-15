import axios from 'axios';

const API_URL = 'http://localhost:8080/account';

// Fungsi untuk mengambil daftar akun
export const getAccountList = async () => {
  try {
    const response = await axios.get(`${API_URL}/list`);
    return response.data.data; // Mengembalikan array account dari respons
  } catch (error) {
    console.error("Error fetching account list:", error);
    throw error;
  }
};

// Fungsi untuk melakukan top-up saldo
export const topUpBalance = async (accountID, amount) => {
  try {
    // Payload yang dikirimkan dalam request
    const payload = {
      account_id: accountID,
      amount: amount,
    };

    // Mengirimkan request POST untuk top-up saldo
    const response = await axios.post(`${API_URL}/topup`, payload);

    if (response.status === 200) {
      return response.data; // Mengembalikan data yang diterima dari API
    }
  } catch (error) {
    console.error("Error during top-up:", error);
    throw error;
  }
};

export const createAccount = async (accountData) => {
  try {
    console.log("Sending account data:", accountData); // Log to check the payload
    const response = await axios.post(`${API_URL}/create`, accountData);
    return response.data;
  } catch (error) {
    console.error("Error creating account:", error.response?.data || error.message); // Log detailed error
    throw error;
  }
};
export const updateAccount = async (accountId, accountData) => {
  try {
    const response = await axios.put(`${API_URL}/update/${accountId}`, accountData);
    return response.data;
  } catch (error) {
    console.error("Error updating account:", error);
    throw error;
  }
};
