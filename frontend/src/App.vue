<template>
  <div class="min-h-screen bg-gradient-to-br from-gray-900 via-gray-800 to-gray-900">
    <header class="bg-black/30 backdrop-blur-md border-b border-gray-700">
      <div class="container mx-auto px-6 py-4">
        <h1 class="text-3xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-blue-400 to-cyan-400">
          Sovereign UI
        </h1>
        <p class="text-gray-400">Command Your Network</p>
      </div>
    </header>

    <main class="container mx-auto px-6 py-8">
      <div v-if="loading" class="text-center text-gray-400">
        <p>Loading network status...</p>
      </div>
      <div v-else-if="error" class="text-center text-red-400">
        <p>Error: {{ error }}</p>
      </div>
      <UserList v-else :users="users" @user-selected="selectUser" />
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import UserList from './components/UserList.vue';

const users = ref([]);
const loading = ref(true);
const error = ref(null);

const selectUser = (user) => {
  // Find the user in the list and toggle a 'selected' state
  users.value.forEach(u => {
    if (u.user === user.user) {
      u.selected = !u.selected;
    } else {
      u.selected = false;
    }
  });
};

onMounted(async () => {
  try {
    const response = await axios.get('/api/v1/nodes');
    users.value = response.data.map(u => ({ ...u, selected: false }));
  } catch (err) {
    error.value = err.message || 'An error occurred while fetching data.';
  } finally {
    loading.value = false;
  }
});
</script>