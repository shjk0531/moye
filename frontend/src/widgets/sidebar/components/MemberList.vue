<template>
    <div class="member-area bg-gray-150 w-50">
        <div class="member-header">
            <h2 class="text-lg font-semibold text-gray-50">Members</h2>
        </div>
    </div>
</template>

<script lang="ts">
export default {
    name: 'MemberList',
};
</script>

<style scoped lang="scss">
.member-area {
    grid-area: memberList;
}
</style>
