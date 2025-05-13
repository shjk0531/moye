// src/features/calendar/index.ts

// API Functions (Re-export from entities)
export { fetchCalendars } from '@/entities/calendar';

// Composables
export * from './composables/useCalendar';

// UI Components
export { default as CalendarList } from './components/CalendarList.vue';
export { default as CalendarView } from './components/CalendarView.vue';
