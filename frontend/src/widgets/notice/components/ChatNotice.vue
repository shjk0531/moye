<template>
    <div class="flex flex-row w-full justify-between items-center">
        <div class="flex flex-row flex-1 items-center gap-2">
            <h2 class="text-lg font-semibold text-gray-50">Notice</h2>
        </div>
        <div class="notice-icons flex flex-row gap-4 text-2xl">
            <!-- defineAsyncComponent으로 동적으로 로드 -->
            <ChannelButton @click="() => handleIconClick('channel')" />
            <CalendarButton @click="() => handleIconClick('calendar')" />
            <PostButton @click="() => handleIconClick('post')" />
            <NoteButton @click="() => handleIconClick('note')" />
            <MemberButton @click="() => handleIconClick('member')" />
        </div>
    </div>
</template>

<script setup lang="ts">
import { defineAsyncComponent } from 'vue';
import { noticeImporters } from '@/shared';

const ChannelButton = defineAsyncComponent(noticeImporters.ChannelButton);
const CalendarButton = defineAsyncComponent(noticeImporters.CalendarButton);
const PostButton = defineAsyncComponent(noticeImporters.PostButton);
const NoteButton = defineAsyncComponent(noticeImporters.NoteButton);
const MemberButton = defineAsyncComponent(noticeImporters.MemberButton);

function handleIconClick(type: string) {
    console.log(
        `${type} 아이콘이 클릭되었지만 아직 처리 기능이 구현되지 않았습니다.`,
    );
}
</script>
