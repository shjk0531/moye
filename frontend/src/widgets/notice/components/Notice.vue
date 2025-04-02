<template>
    <div class="notice-panel">
        <div
            class="flex flex-row justify-between items-center bg-gray-850 px-4 py-2"
        >
            <div class="notice-header">
                <h2 class="text-lg font-semibold text-gray-50">Notice</h2>
            </div>
            <div class="notice-icons flex flex-row gap-4 text-2xl">
                <span
                    v-for="(item, idx) in icons"
                    :key="idx"
                    :class="`mdi ${item.icon} ${item.color} ${item.hover} cursor-pointer`"
                ></span>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
export default {
    name: 'Notice',
    data() {
        return {
            icons: [
                {
                    icon: 'mdi-calendar-check',
                    color: 'text-gray-400',
                    hover: 'hover:text-gray-200',
                },
                {
                    icon: 'mdi-post',
                    color: 'text-gray-400',
                    hover: 'hover:text-gray-200',
                },
                {
                    icon: 'mdi-note',
                    color: 'text-gray-400',
                    hover: 'hover:text-gray-200',
                },
                {
                    icon: 'mdi-account-group',
                    color: 'text-gray-400',
                    hover: 'hover:text-gray-200',
                },
            ],
        };
    },

    methods: {
        // 여기에 필요한 메소드 추가
    },
};
</script>

<style scoped>
.notice-panel {
    grid-area: notice;
}
</style>
