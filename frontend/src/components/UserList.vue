<template>
  <div class="space-y-4">
    <div v-for="user in users" :key="user.user" class="bg-gray-800/50 backdrop-blur-sm rounded-lg border border-gray-700 overflow-hidden">
      <button
        @click="$emit('user-selected', user)"
        class="w-full px-6 py-4 text-left flex justify-between items-center hover:bg-gray-700/50 transition-colors duration-200"
      >
        <h2 class="text-xl font-semibold text-cyan-400">{{ user.user }}</h2>
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" :class="['w-6 h-6 text-gray-400 transition-transform duration-200', user.selected && 'rotate-180']">
          <path stroke-linecap="round" stroke-linejoin="round" d="m19.5 8.25-7.5 7.5-7.5-7.5" />
        </svg>
      </button>
      <div v-if="user.selected" class="px-6 pb-4">
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          <NodeCard v-for="node in user.nodes" :key="node.id" :node="node" />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import NodeCard from './NodeCard.vue';

defineProps({
  users: {
    type: Array,
    required: true
  }
});

defineEmits(['user-selected']);
</script>