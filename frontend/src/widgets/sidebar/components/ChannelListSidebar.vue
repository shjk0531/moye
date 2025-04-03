src/widgets/sidebar/components/ChannelListSidebar.vue

<template>
    <div class="channel-notice text-white p-2">
        <p>공지사항</p>
    </div>

    <ChannelList />
</template>

<script lang="ts">
import { ChannelList } from '@/entities/channel';

export default {
    name: 'ChannelListSidebar',
    components: {
        ChannelList,
    },
};
</script>

<style scoped lang="scss"></style>
