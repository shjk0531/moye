<template>
    <div class="flex flex-row justify-between items-center">
        <div class="notice-header">
            <h2 class="text-lg font-semibold text-gray-50">Notice</h2>
        </div>
        <div class="notice-icons flex flex-row gap-4 text-2xl">
            <ChannelButton />
            <CalendarButton />
            <PostButton />
            <NoteButton />
            <MemberButton />
        </div>
    </div>
</template>

<script setup lang="ts">
import {
    ChannelButton,
    CalendarButton,
    PostButton,
    NoteButton,
    MemberButton,
} from '@/shared/ui/Button';

function handleIconClick(item: { type: string }) {
    console.log(
        `${item.type} 아이콘이 클릭되었지만 아직 처리 기능이 구현되지 않았습니다.`,
    );
}
</script>

<style scoped></style>
