<!-- src/pages/lounge/LoungeListPage.vue -->
<template>
    <PageLayout>
        <template #notice>
            <LoungeListNotice />
        </template>
        <template #content>
            <LoungeList />
        </template>
    </PageLayout>
</template>

<script setup lang="ts">
import { LoungeList } from '@/entities';
import { PageLayout } from '@/shared/layout';
import { LoungeListNotice } from '@/widgets';
</script>
