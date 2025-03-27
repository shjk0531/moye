<template>
    <div class="sidebar">
        <div class="server-list">
            <!-- 예시: 서버 아이콘 목록 -->
            <div class="server-item" v-for="server in servers" :key="server.id">
                <img :src="server.icon" alt="server icon" />
            </div>
        </div>
        <div class="channel-list">
            <!-- PrimeVue PanelMenu로 채널 목록 구성 -->
            <PanelMenu :model="channels" class="channel-panel" />
        </div>
    </div>
</template>

<script>
export default {
    name: 'Sidebar',
    data() {
        return {
            servers: [
                { id: 1, icon: 'https://via.placeholder.com/40' },
                { id: 2, icon: 'https://via.placeholder.com/40' },
                { id: 3, icon: 'https://via.placeholder.com/40' },
            ],
            channels: [
                {
                    label: '텍스트 채널',
                    items: [
                        { label: '# general', icon: 'pi pi-comments' },
                        { label: '# random', icon: 'pi pi-comment' },
                    ],
                },
                {
                    label: '보이스 채널',
                    items: [
                        { label: 'General', icon: 'pi pi-microphone' },
                        { label: 'Gaming', icon: 'pi pi-gamepad' },
                    ],
                },
            ],
        };
    },
};
</script>

<style scoped>
.sidebar {
    width: 250px;
    background-color: #2f3136;
    display: flex;
    flex-direction: column;
}

.server-list {
    display: flex;
    flex-direction: column;
    padding: 10px;
}

.server-item {
    width: 50px;
    height: 50px;
    margin-bottom: 10px;
    border-radius: 25px;
    overflow: hidden;
    cursor: pointer;
    border: 2px solid transparent;
}

.server-item:hover {
    border-color: #7289da;
}

.server-item img {
    width: 100%;
    height: 100%;
}

.channel-list {
    flex: 1;
    padding: 10px;
    overflow-y: auto;
}
</style>
