<!-- src/widgets/userPanel/components/UserPanel.vue -->

<template>
    <div class="pb-4 px-2 w-(--custom-side-bar-left-width)">
        <div
            class="user-panel flex items-center p-2 bg-gray-750 text-white rounded-lg h-(--custom-bottom-pannel-height)"
        >
            <!-- 사용자 아이콘 -->
            <img
                class="user-icon w-8 h-8 rounded-full mr-2"
                :src="userProfile || 'https://picsum.photos/200/200?random=5'"
                alt="User Icon"
            />

            <!-- 사용자 닉네임 및 태그 -->
            <div class="user-info flex-grow flex flex-col items-start">
                <p class="user-name font-bold m-0">
                    {{ userNickname }}
                </p>
                <p class="user-discriminator text-xs text-gray-400 m-0">
                    #{{ userEmail }}
                </p>
            </div>

            <!-- 액션 버튼: 마이크, 헤드셋, 설정 -->
            <div class="user-actions flex space-x-0.5">
                <button
                    @click="toggleMic"
                    class="action-button p-1 hover:bg-gray-700 rounded"
                >
                    <!-- Font Awesome 아이콘 사용 -->
                    <!-- 마이크 아이콘: 상태에 따라 아이콘, 색상 변경 -->
                    <i v-if="micOn" class="mdi mdi-microphone"></i>
                    <i v-else class="mdi mdi-microphone-off text-red-500"></i>
                </button>
                <button
                    @click="toggleHeadset"
                    class="action-button p-1 hover:bg-gray-700 rounded"
                >
                    <i v-if="headsetOn" class="mdi mdi-headphones"></i>
                    <i v-else class="mdi mdi-headphones-off text-red-500"></i>
                </button>
                <button
                    @click="openSettings"
                    class="action-button p-1 hover:bg-gray-700 rounded"
                >
                    <i class="mdi mdi-cog"></i>
                </button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useUserStore, useAppStore } from '@/store';

// 임시 사용자 데이터 (실제 구현에서는 props 또는 전역 상태 관리로 대체)
const userStore = useUserStore();
const appStore = useAppStore();

const userNickname = computed(() => userStore.userNickname);
const userEmail = computed(() => userStore.userEmail);
const userProfile = computed(() => userStore.userProfile);

const micOn = computed(() => appStore.isMicOn);
const headsetOn = computed(() => appStore.isHeadsetOn);

function toggleMic() {
    appStore.toggleMic();
}

function toggleHeadset() {
    appStore.toggleHeadset();
}

function openSettings() {
    // 사용자 설정 모달 열기 등의 처리를 여기에 구현합니다.
    console.log('User settings opened');
}
</script>

<style scoped>
.user-info {
    flex-grow: 1;
}

.user-name {
    margin: 0;
    font-weight: bold;
}

.user-discriminator {
    margin: 0;
    font-size: 0.75rem;
    color: #b9bbbe;
}

.user-actions button {
    background: transparent;
    border: none;
    padding: 4px;
    cursor: pointer;
}

.user-actions button:hover {
    background-color: rgba(79, 84, 92, 0.32);
    border-radius: 4px;
}
</style>
