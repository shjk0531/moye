<template>
    <div class="flex flex-row w-full justify-between items-center">
        <div class="flex flex-row flex-1 items-center gap-2">
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

<script lang="ts">
import {
    ChannelButton,
    CalendarButton,
    PostButton,
    NoteButton,
    MemberButton,
} from '@/shared/ui/Button';

export default {
    name: 'ChatNotice',
    components: {
        ChannelButton,
        CalendarButton,
        PostButton,
        NoteButton,
        MemberButton,
    },

    methods: {
        handleIconClick(item: { type: string }) {
            console.log(
                `${item.type} 아이콘이 클릭되었지만 아직 처리 기능이 구현되지 않았습니다.`,
            );
            // 필요한 경우 추가 처리 로직 구현
        },
    },
};
</script>

<style scoped></style>
