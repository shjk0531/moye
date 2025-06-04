<!-- src/entities/lounge/components/LoungeList.vue -->
<template>
    <div
        class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4 p-4"
    >
        <LoungeCard
            v-for="lounge in lounges"
            :key="lounge.id"
            :lounge="lounge"
            @select="openModal"
        />
    </div>

    <!-- selectedLounge가 있을 때만 모달을 렌더링 -->
    <LoungeDetailModal
        v-if="selectedLounge"
        :lounge="selectedLounge"
        @close="closeModal"
    />
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { fetchLoungeList, type LoungeInfo } from '@/entities';
import LoungeCard from './LoungeCard.vue';
import LoungeDetailModal from './LoungeDetailModal.vue';

const lounges = ref<LoungeInfo[]>([]);
// 모달에 전달할 “선택된” 라운지 정보
const selectedLounge = ref<LoungeInfo | null>(null);

const getLoungeList = async () => {
    const response = await fetchLoungeList();
    lounges.value = response.lounges;
};

onMounted(() => {
    getLoungeList();
});

/**
 * 카드 클릭 시 호출출
 * @param lounge 클릭된 LoungeInfo 객체
 */
const openModal = (lounge: LoungeInfo) => {
    selectedLounge.value = lounge;
};

/**
 * 모달이 닫힐 때 호출
 */
const closeModal = () => {
    selectedLounge.value = null;
};
</script>
