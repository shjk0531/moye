<template>
    <div
        ref="container"
        class="flex flex-1 flex-col w-[var(--custom-item-list-width)] text-gray-50 relative"
        @contextmenu.prevent="onContextMenu"
        @click="closeContextMenu"
    >
        <!-- response.items를 순회하면서, 그룹/채널 구분하여 렌더링 -->
        <div
            v-for="item in response?.items"
            :key="item.item_type === 'group' ? item.group!.id : item.channel!.id"
        >
            <!-- 1) 그룹 헤더 -->
            <div
                v-if="item.item_type === 'group'"
                @click="toggleGroup(item.group!.id)"
                @contextmenu.stop
                class="panel-item cursor-pointer flex items-center text-sm p-2 rounded-lg mx-2 my-0.5 justify-between text-gray-300 hover:text-gray-100"
            >
                <span>
                    {{ item.group!.name }}
                    <span class="mr-2">
                        <i
                            :class="
                expandedGroups[item.group!.id]
                  ? 'mdi mdi-chevron-down'
                  : 'mdi mdi-chevron-right'
              "
                        ></i>
                    </span>
                </span>
                <span class="ml-2">
                    <i
                        class="mdi mdi-plus"
                        @click.stop="addChannelToGroup(item.group!.id)"
                    ></i>
                </span>
            </div>

            <!-- 2) 그룹 내부 채널 렌더링 -->
            <div v-if="item.item_type === 'group'">
                <!-- 그룹이 펼쳐진 상태 -->
                <div v-if="expandedGroups[item.group!.id]">
                    <div
                        v-for="ch in item.group!.channels"
                        :key="ch.id"
                        class="panel-item cursor-pointer flex items-center text-sm p-2 rounded-lg mx-2 my-0.5"
                        :class="{
                            'text-gray-100 bg-gray-750':
                                currentChannelId === ch.id,
                            'text-gray-300 bg-gray-950 hover:bg-gray-750 hover:text-gray-100':
                                currentChannelId !== ch.id,
                        }"
                        @click="emit('selectChannel', ch.id)"
                        @contextmenu.stop
                    >
                        {{ ch.name }}
                    </div>
                </div>

                <!-- 그룹은 접혀 있는데, 현재 채널이 이 그룹에 속해 있는 경우 -->
                <div
                    v-else-if="
            item.group!.channels.some((ch) => ch.id === currentChannelId)
          "
                >
                    <div
                        v-for="ch in item.group!.channels.filter((ch) => ch.id === currentChannelId)"
                        :key="ch.id"
                        class="panel-item cursor-pointer flex items-center text-sm p-2 rounded-lg mx-2 my-0.5"
                        :class="{
                            'text-gray-100 bg-gray-750':
                                currentChannelId === ch.id,
                            'text-gray-300 bg-gray-950 hover:bg-gray-750 hover:text-gray-100':
                                currentChannelId !== ch.id,
                        }"
                        @click="emit('selectChannel', ch.id)"
                        @contextmenu.stop
                    >
                        {{ ch.name }}
                    </div>
                </div>
            </div>

            <!-- 3) item_type==='channel' -->
            <div
                v-else-if="item.item_type === 'channel'"
                class="panel-item cursor-pointer flex items-center text-sm p-2 rounded-lg mx-2 my-0.5"
                :class="{
          'text-gray-100 bg-gray-750': currentChannelId === item.channel!.id,
          'text-gray-300 bg-gray-950 hover:bg-gray-750 hover:text-gray-100':
            currentChannelId !== item.channel!.id,
        }"
                @click="emit('selectChannel', item.channel!.id)"
                @contextmenu.stop
            >
                {{ item.channel!.name }}
            </div>
        </div>

        <!-- 4) items 자체가 비어 있을 때 -->
        <div
            v-if="response?.items.length === 0"
            class="mt-6 text-center text-gray-400 flex-1"
        >
            아직 생성된 채널이 없습니다.
        </div>

        <!-- Context Menu -->
        <div
            v-if="contextMenuVisible"
            class="absolute bg-gray-950 shadow-lg rounded-md z-500 border border-gray-750 text-sm p-2"
            :style="{ top: `${contextMenuY}px`, left: `${contextMenuX}px` }"
            @click.stop
        >
            <ul>
                <li
                    class="px-4 py-2 hover:bg-gray-800 cursor-pointer text-gray-300 bg-gray-950 border border-gray-950 rounded-md hover:text-gray-100"
                    @click="createChannel"
                >
                    채널 만들기
                </li>
                <li
                    class="px-4 py-2 hover:bg-gray-800 cursor-pointer text-gray-300 bg-gray-950 border border-gray-950 rounded-md hover:text-gray-100"
                    @click="createGroup"
                >
                    그룹 만들기
                </li>
            </ul>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useLoungeStore } from '@/store/lounge';
import type { LoungeChannelResponse } from '@/entities/channel/models/types';

// 1) route에서 loungeId, channelId 가져오기
const route = useRoute();
const router = useRouter();
const loungeId = computed(() => route.params.loungeId as string);
const currentChannelId = computed(() => route.params.channelId as string);

// 2) 라운지 스토어/응답 객체 정의
const loungeStore = useLoungeStore();
const response = ref<LoungeChannelResponse>({ items: [] });

// 3) “그룹 내부 채널” ID들을 모아둘 Set
const groupedChannelIds = ref<Set<string>>(new Set());

// 4) 어떤 그룹이 펼쳐진 상태인지 boolean으로 관리
const expandedGroups = reactive<Record<string, boolean>>({});

// Context menu 상태
const contextMenuVisible = ref(false);
const contextMenuX = ref(0);
const contextMenuY = ref(0);

const emit = defineEmits<{
    (e: 'selectChannel', channelId: string): void;
}>();

function processResponse() {
    groupedChannelIds.value.clear();
    response.value.items.forEach((item) => {
        if (item.item_type === 'group' && item.group) {
            expandedGroups[item.group.id] = false;
            item.group.channels.forEach((ch) => {
                groupedChannelIds.value.add(ch.id);
            });
        }
    });
}

async function fetchChannels(loungeId: string) {
    response.value = await loungeStore.loadChannels(loungeId);
    processResponse();
}

function toggleGroup(groupId: string) {
    expandedGroups[groupId] = !expandedGroups[groupId];
}

function addChannelToGroup(groupId: string) {
    console.log('addChannelToGroup', groupId);
}

/** Context menu 핸들러 */
function onContextMenu(event: MouseEvent) {
    // 패널 아이템 위가 아니면 메뉴 열기
    const target = event.target as HTMLElement;
    if (!target.closest('.panel-item')) {
        contextMenuX.value = event.offsetX;
        contextMenuY.value = event.offsetY;
        contextMenuVisible.value = true;
    }
}

function closeContextMenu() {
    contextMenuVisible.value = false;
}

function createChannel() {
    // 채널 생성 로직 호출
    console.log('채널 만들기 클릭');
    contextMenuVisible.value = false;
}

function createGroup() {
    // 그룹 생성 로직 호출
    console.log('그룹 만들기 클릭');
    contextMenuVisible.value = false;
}

onMounted(async () => {
    await fetchChannels(loungeId.value);
});

watch(loungeId, (newLoungeId, oldLoungeId) => {
    if (newLoungeId && newLoungeId !== oldLoungeId) {
        fetchChannels(newLoungeId);
    }
});
</script>

<style scoped>
/* Context menu 기본 스타일 (필요에 따라 수정) */
.absolute ul {
    list-style: none;
    margin: 0;
    padding: 0;
}
.absolute li {
    white-space: nowrap;
}
</style>
