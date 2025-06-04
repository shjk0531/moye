<template>
    <!-- LoungeCreationWizard 위젯을 래핑하고 이벤트 핸들러를 바인딩 -->
    <LoungeCreationWizard @cancel="onCancel" @create="onCreate" />
</template>

<script setup lang="ts">
import { LoungeCreationWizard } from '@/widgets/LoungeCreationWizard';
import { useRouter } from 'vue-router';
import { useAppStore } from '@/store';
import { createLounge } from '@/entities/lounge/api/loungeApi';
import { useToast } from 'primevue/usetoast';
import type { LoungeCreatePayload } from '@/entities/lounge/models/types';

const router = useRouter();
const appStore = useAppStore();
const toast = useToast();

// 취소 시 마지막으로 저장된 라우트로 이동
function onCancel() {
    const lastRoute = appStore.getLastRoute || '/';
    console.log('lastRoute', lastRoute);
    router.push(lastRoute);
}

// 생성 시 API 호출 후 성공/실패 토스트 출력 및 라우팅
async function onCreate(payload: {
    content: string;
    tags: string[];
    intro: string;
    title: string;
    thumbnailUrl: string;
}) {
    try {
        // 폼 데이터를 백엔드 API에 맞게 변환
        const loungePayload: LoungeCreatePayload = {
            name: payload.title,
            profile_url: payload.thumbnailUrl,
            description: payload.intro,
            content: payload.content,
            tags: payload.tags,
        };

        await createLounge(loungePayload);

        toast.add({
            severity: 'success',
            summary: '라운지 생성',
            detail: '라운지가 성공적으로 생성되었습니다.',
            life: 3000,
        });

        // 생성 후 라운지 목록 페이지로 이동
        router.push('/lounge');
    } catch (error) {
        console.error('라운지 생성 실패:', error);
        toast.add({
            severity: 'error',
            summary: '생성 실패',
            detail: '라운지 생성에 실패했습니다.',
            life: 5000,
        });
    }
}
</script>
