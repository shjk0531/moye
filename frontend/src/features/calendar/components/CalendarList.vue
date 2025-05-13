<!-- src/features/calender/components/CalendarList.vue -->

<template></template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { fetchCalendars } from '@/entities/calendar';

const route = useRoute();
const router = useRouter();

const calendarsGrouped = ref([]);
const calendarsUngrouped = ref([]);
</script>
