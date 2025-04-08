<template>
    <div class="bg-gray-150 h-full">
        <div class="member-header">
            <h2 class="text-lg font-semibold text-gray-50">Members</h2>
        </div>
    </div>
</template>

<script lang="ts">
import { mapState } from 'vuex';

export default {
    name: 'MemberList',
    computed: {
        ...mapState(['isMemberListVisible']),
    },
};
</script>

<style scoped lang="scss"></style>
