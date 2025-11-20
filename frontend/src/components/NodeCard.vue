<template>
  <div class="bg-gray-900/50 rounded-lg p-4 border border-gray-600">
    <div class="flex items-center justify-between mb-2">
      <h3 class="font-bold text-lg">{{ node.givenName || node.name }}</h3>
      <span :class="statusClass" class="w-3 h-3 rounded-full"></span>
    </div>
    <p class="text-sm text-gray-400 mb-1">{{ node.ipAddresses.join(', ') }}</p>
    <p class="text-xs text-gray-500">Last Seen: {{ formattedLastSeen }}</p>
  </div>
</template>

<script setup>
import { computed } from 'vue';

const props = defineProps({
  node: {
    type: Object,
    required: true
  }
});

const statusClass = computed(() => {
  return props.node.online ? 'bg-green-500' : 'bg-gray-500';
});

const formattedLastSeen = computed(() => {
  if (!props.node.lastSeen) return 'Never';
  const date = new Date(props.node.lastSeen);
  return date.toLocaleString();
});
</script>