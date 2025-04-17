// src/features/channel/index.ts

// API Functions (Re-export from entities)
export {
    fetchChannelsGrouped,
    fetchChannelsUngrouped,
    findFirstChannel,
    type Channel,
    type ChannelGroup,
} from '@/entities/channel';

// Composables
export * from './composables/useChannel';

// UI Components
export { default as ChannelList } from './components/ChannelList.vue';
