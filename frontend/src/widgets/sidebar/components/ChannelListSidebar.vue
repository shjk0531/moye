src/widgets/sidebar/components/ChannelListSidebar.vue

<template>
    <div class="flex flex-1 flex-col">
        <div
            class="flex channel-notice text-white p-2 h-(--custom-notice-bar-height) items-center justify-center"
        >
            <p
                class="text-lg font-bold w-full justify-center items-center text-center"
            >
                공지사항
            </p>
        </div>

        <ChannelList @selectChannel="onSelectChannel" />
    </div>
</template>

<script setup lang="ts">
import { ChannelList } from '@/entities/channel';
import { useRouter, useRoute } from 'vue-router';

const router = useRouter();
const route = useRoute();
const loungeId = route.params.loungeId as string;

function onSelectChannel(channelId: string) {
    router.push({
        name: 'chat',
        params: {
            loungeId,
            channelId,
        },
    });
}
</script>

<style scoped lang="scss"></style>
