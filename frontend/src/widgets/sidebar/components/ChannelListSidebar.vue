src/widgets/sidebar/components/ChannelListSidebar.vue

<template>
    <div class="flex flex-col">
        <div
            class="flex channel-notice text-white p-2 h-(--custom-notice-bar-height) items-center justify-center"
        >
            <p
                class="text-lg font-bold w-full justify-center items-center text-center"
            >
                공지사항
            </p>
        </div>

        <ChannelList />
    </div>
</template>

<script setup lang="ts">
import { ChannelList } from '@/features/channel';
</script>

<style scoped lang="scss"></style>
