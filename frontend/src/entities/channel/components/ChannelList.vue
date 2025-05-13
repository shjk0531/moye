<!-- src/entities/channel/components/ChannelList.vue -->
<template>
    <div>
        <h1>Channel List</h1>
        <div v-for="item in response?.items" :key="item.channel?.id">
            {{ item.channel?.name }}
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { fetchChannels } from '@/entities/channel';
import type { StudyChannelResponse } from '@/entities/channel';

const route = useRoute();
const router = useRouter();

const response = ref<StudyChannelResponse>();

onMounted(async () => {
    response.value = await fetchChannels(route.params.studyId as string);
});
</script>
