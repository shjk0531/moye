<template>
    <div class="bg-gray-150 h-full">
        <div class="member-header">
            <h2 class="text-lg font-semibold text-gray-50">Members</h2>
        </div>
    </div>
</template>

<script lang="ts">
import { useAppStore } from '@/store';

export default {
    name: 'MemberList',
    setup() {
        const appStore = useAppStore();
        return { appStore };
    },
};
</script>

<style scoped lang="scss"></style>
