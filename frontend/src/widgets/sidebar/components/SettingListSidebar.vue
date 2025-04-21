<template>
    <div class="flex flex-col w-full p-10">
        <div v-for="group in settingsList" :key="group.title">
            <h3 class="px-4 py-1 font-semibold text-gray-300 text-xs">
                {{ group.title }}
            </h3>
            <ul>
                <li
                    v-for="item in group.items"
                    :key="item.route"
                    class="px-4 py-1 hover:bg-gray-750 hover:text-gray-50 text-gray-200 cursor-pointer mb-1 rounded"
                >
                    <!-- full path 사용 -->
                    <router-link
                        :to="`/settings/${item.route}`"
                        class="block w-full h-full"
                    >
                        {{ item.name }}
                    </router-link>
                </li>
            </ul>
            <div class="px-4">
                <div class="border-t border-gray-700 my-2" />
            </div>
        </div>
        <div
            class="flex flex-raw justify-between px-4 py-1 hover:bg-gray-750 hover:text-gray-50 text-gray-200 cursor-pointer mb-1 rounded"
            @click="handleLogout"
        >
            <p class="">로그아웃</p>
            <span class="mdi mdi-logout"></span>
        </div>
    </div>
</template>

<script setup lang="ts">
import { settingsList } from '@/widgets/sidebar';
import { useLogout } from '@/features/auth';

const emit = defineEmits(['success']);
const { logout } = useLogout((event) => emit(event));

const handleLogout = () => {
    logout();
};
</script>
