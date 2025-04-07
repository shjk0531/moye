<template>
    <span :class="iconClasses" @click="handleClick" :title="itemTitle"></span>
</template>

<script>
import {
    navigateToItem,
    isItemActive,
    isItemTypeSupported,
} from '../services/navigationService';

export default {
    name: 'StudyNoticeIcon',
    props: {
        item: {
            type: Object,
            required: true,
        },
        isMemberListVisible: {
            type: Boolean,
            default: false,
        },
    },
    computed: {
        iconClasses() {
            // 기본 클래스 문자열 구성
            let classes = `mdi ${this.item.icon} ${this.item.hover} cursor-pointer `;

            // member 아이콘의 경우 isMemberListVisible 값에 따라 색상 변경
            if (this.item.type === 'member') {
                classes += this.isMemberListVisible
                    ? this.item.active
                    : this.item.color;
            } else if (
                this.item.type === 'channel' ||
                this.item.type === 'calendar' ||
                this.item.type === 'post' ||
                this.item.type === 'note'
            ) {
                // 현재 라우트를 확인하여 활성화된 아이콘 스타일 적용
                const currentPath = this.$route.path;
                // 아이템 ID가 있을 경우 라우트 객체와 함께 전달
                const activeItemId =
                    this.$store.state.activeItems[this.$route.params.studyId]?.[
                        this.item.type
                    ];
                const active = isItemActive(
                    this.item.type,
                    currentPath,
                    this.$route,
                    activeItemId,
                );
                classes += active ? this.item.active : this.item.color;
            } else {
                // 지원되지 않는 기능인 경우 다른 스타일 적용 가능
                const isSupported = isItemTypeSupported(this.item.type);
                classes += isSupported
                    ? this.item.color
                    : 'text-gray-600 opacity-50';
            }

            return classes;
        },
        itemTitle() {
            // 기본 타이틀
            const typeNameMap = {
                channel: '채널',
                calendar: '캘린더',
                post: '게시글',
                note: '노트',
                member: '멤버',
            };

            const typeName = typeNameMap[this.item.type] || this.item.type;

            // 지원되지 않는 기능인 경우 알림 추가
            if (
                !isItemTypeSupported(this.item.type) &&
                this.item.type !== 'member'
            ) {
                return `${typeName} (준비 중)`;
            }

            return typeName;
        },
    },
    methods: {
        async handleClick() {
            // member 아이콘 클릭 시 toggle 이벤트 emit
            if (this.item.type === 'member') {
                this.$emit('toggle-member-list');
                return;
            }

            // 지원되지 않는 기능은 별도 처리
            if (!isItemTypeSupported(this.item.type)) {
                this.$emit('icon-click', this.item);
                console.log(`${this.item.type} 기능은 현재 개발 중입니다.`);
                return;
            }

            // 채널, 캘린더, 게시글, 노트 아이콘 클릭 시
            if (
                ['channel', 'calendar', 'post', 'note'].includes(this.item.type)
            ) {
                const store = this.$store;
                const router = this.$router;
                const studyId = this.$route.params.studyId;

                if (!studyId) {
                    console.error('유효한 studyId가 없습니다.');
                    return;
                }

                await navigateToItem(this.item.type, studyId, store, router);
                return;
            }

            // 그 외 타입의 아이콘에 대한 처리
            this.$emit('icon-click', this.item);
        },
    },
};
</script>

<style scoped>
/* 필요한 스타일 추가 */
</style>
