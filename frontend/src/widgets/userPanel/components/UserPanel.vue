<template>
    <div
        class="user-panel flex items-center p-2 bg-gray-750 text-white rounded-lg"
    >
        <!-- 사용자 아이콘 -->
        <img
            class="user-icon w-8 h-8 rounded-full mr-2"
            :src="user.icon"
            alt="User Icon"
        />

        <!-- 사용자 닉네임 및 태그 -->
        <div class="user-info flex-grow flex flex-col items-start">
            <p class="user-name font-bold m-0">
                {{ user.nickname }}
            </p>
            <p class="user-discriminator text-xs text-gray-400 m-0">
                #{{ user.discriminator }}
            </p>
        </div>

        <!-- 액션 버튼: 마이크, 헤드셋, 설정 -->
        <div class="user-actions flex space-x-2">
            <button
                @click="toggleMic"
                class="action-button p-1 hover:bg-gray-700 rounded"
            >
                <i :class="micClass" class="pi pi-microphone"></i>
            </button>
            <button
                @click="toggleHeadset"
                class="action-button p-1 hover:bg-gray-700 rounded"
            >
                <i :class="headsetClass" class="pi pi-headphones"></i>
            </button>
            <button
                @click="openSettings"
                class="action-button p-1 hover:bg-gray-700 rounded"
            >
                <i class="pi pi-cog"></i>
            </button>
        </div>
    </div>
</template>

<script>
export default {
    name: 'UserPanel',
    data() {
        return {
            // 임시 사용자 데이터 (실제 구현에서는 props 또는 전역 상태 관리로 대체)
            user: {
                icon: 'https://picsum.photos/200/200?random=5',
                nickname: 'Username',
                discriminator: '1234',
                micOn: true,
                headsetOn: true,
            },
        };
    },
    computed: {
        // Font Awesome 아이콘 클래스를 사용하여 상태에 따른 아이콘 변경
        micClass() {
            return this.user.micOn ? 'text-gray-50' : 'text-red-500';
        },
        headsetClass() {
            return this.user.headsetOn ? 'text-gray-50' : 'text-red-500';
        },
    },
    methods: {
        toggleMic() {
            this.user.micOn = !this.user.micOn;
        },
        toggleHeadset() {
            this.user.headsetOn = !this.user.headsetOn;
        },
        openSettings() {
            // 사용자 설정 모달 열기 등의 처리를 여기에 구현합니다.
            console.log('User settings opened');
        },
    },
};
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
