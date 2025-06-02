<!-- src/entities/channel/components/ChannelList.vue -->
<template>
    <div
        class="flex flex-1 flex-col w-[var(--custom-item-list-width)] text-gray-50"
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
                class="panel-item cursor-pointer flex items-center hover:bg-gray-750 text-sm p-2 rounded-lg mx-2 my-0.5 text-gray-100 justify-between"
            >
                <span>{{ item.group!.name }}</span>
                <span class="mr-2">
                    <i
                        :class="
                            expandedGroups[item.group!.id]
                                ? 'mdi mdi-chevron-down text-gray-50'
                                : 'mdi mdi-chevron-right text-gray-50'
                        "
                    ></i>
                </span>
            </div>

            <!-- 2) 해당 그룹이 펼쳐진 경우에만, 그룹 내부 채널들 보여주기 -->
            <div
                v-if="item.item_type === 'group' && expandedGroups[item.group!.id]"
            >
                <div
                    v-for="ch in item.group!.channels"
                    :key="ch.id"
                    class="panel-item cursor-pointer flex items-center hover:bg-gray-750 text-sm p-2 rounded-lg mx-2 my-0.5 text-gray-100"
                    :class="{
                        'bg-gray-750': currentChannelId === ch.id,
                        'bg-gray-950': currentChannelId !== ch.id,
                    }"
                    @click="router.push(`/study/${studyId}/channel/${ch.id}`)"
                >
                    {{ ch.name }}
                </div>
            </div>

            <!-- 3) item_type==='channel'이면서, groupedChannelIds에 포함되지 않은 채널만 렌더링 -->
            <div
                v-else-if="item.item_type === 'channel'"
                class="panel-item cursor-pointer flex items-center hover:bg-gray-750 text-sm p-2 rounded-lg mx-2 my-0.5 text-gray-100"
                :class="{
                    'bg-gray-750': currentChannelId === item.channel!.id,
                    'bg-gray-950': currentChannelId !== item.channel!.id,
                }"
                @click="
                    router.push(`/study/${studyId}/channel/${item.channel!.id}`)
                "
            >
                {{ item.channel!.name }}
            </div>
        </div>

        <!-- 4) items 자체가 비어 있을 때 (그룹도 채널도 하나도 없을 경우) -->
        <div
            v-if="response?.items.length === 0"
            class="mt-6 text-center text-gray-400"
        >
            아직 생성된 채널이 없습니다.
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useStudyStore } from '@/store/study';
import type { StudyChannelResponse } from '@/entities/channel/models/types';

// 1) route에서 studyId 가져오기
const route = useRoute();
const router = useRouter();
const studyId = computed(() => route.params.studyId as string);
const currentChannelId = computed(() => route.params.channelId as string);

// 2) 스터디 스토어/응답 객체 정의
const studyStore = useStudyStore();
const response = ref<StudyChannelResponse>({ items: [] });

// 3) “그룹 내부 채널” ID들을 모아둘 Set
//    그룹에 속하지 않은 채널만 별도로 렌더링하기 위해 사용
const groupedChannelIds = ref<Set<string>>(new Set());

// 4) 어떤 그룹이 펼쳐진 상태인지 boolean으로 관리
//    { [groupId]: true/false }
const expandedGroups = reactive<Record<string, boolean>>({});

/**
 * response.value가 채워지면,
 * - expandedGroups에 모든 그룹 ID를 등록하고 기본값을 false로 세팅
 * - groupedChannelIds에 “모든 그룹 내부 채널 ID”를 모아둔다
 */
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

async function fetchChannels(studyId: string) {
    response.value = await studyStore.loadChannels(studyId);
    processResponse();
}

/** 그룹 헤더를 클릭했을 때, 해당 그룹의 boolean을 토글 */
function toggleGroup(groupId: string) {
    expandedGroups[groupId] = !expandedGroups[groupId];
}

// 5) 컴포넌트 마운트 시 한 번만 호출
onMounted(async () => {
    await fetchChannels(studyId.value);
});

watch(studyId, (newStudyId, oldStudyId) => {
    if (newStudyId && newStudyId !== oldStudyId) {
        fetchChannels(newStudyId);
    }
});
</script>

<style scoped>
/* 
  • flex 항목 간격이나 색상은 필요에 따라 조정하세요.
  • Tailwind 클래스(hover:bg-gray-750 등)도 프로젝트 컬러 팔레트에 맞게 바꾸셔도 됩니다. 
*/
</style>
