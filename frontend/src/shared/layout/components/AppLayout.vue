<!-- src.shared.layout.components.AppLayout.vue -->

<template>
  <!-- 그리드 컨테이너 -->
  <div class="grid-container">
    <slot />
  </div>
</template>

<script>
export default {
  name: "AppLayout",
};
</script>

<style scoped lang="scss">
/* 기본 Desktop 설정 */
.grid-container {
  margin: 0 auto;
  padding-left: 40px;
  padding-right: 40px;
  max-width: 1200px;
  display: grid;
  grid-template-columns: repeat(12, 1fr);
  gap: 24px; // 거터

  /* Tablet 반응형 설정 */
  @media (max-width: 820px) {
    padding-left: 24px;
    padding-right: 24px;
    max-width: 768px;
    grid-template-columns: repeat(6, 1fr);
  }

  /* Mobile 반응형 설정 */
  @media (max-width: 414px) {
    padding-left: 16px;
    padding-right: 16px;
    max-width: 360px;
    grid-template-columns: repeat(4, 1fr);
  }
}
</style>
