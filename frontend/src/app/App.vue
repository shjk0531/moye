<!-- src/app/App.vue -->

<template>
  <div id="app">
    <!-- 네비게이션 바 (필요 시) -->
    <header>
      <AppLayout>
        <nav>
          <router-link to="/">Home</router-link>
          <router-link to="/login">Login</router-link>
        </nav>
      </AppLayout>
    </header>
    <!-- 현재 라우트에 따른 컴포넌트 출력 -->
    <main>
      <AppLayout>
        <div class="grid-item" v-for="n in 12" :key="n">
          <p>셀 {{ n }}</p>
        </div>
        <router-view />
      </AppLayout>
    </main>
    <!-- 푸터 등 추가 UI 요소들 -->
    <footer>
      <p>© 2025 Your Company</p>
    </footer>
  </div>
</template>

<script setup lang="ts">
import AppLayout from "./components/AppLayout.vue";

// App.vue에서는 라우트 등록이 필요 없습니다.
// main.ts에서 createApp(App).use(router).mount('#app')로 이미 등록되어 있습니다.
</script>

<style>
#html,
body,
#app {
  margin: 0;
  font-family: Arial, sans-serif;
}
</style>

<style scoped>
#app {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  text-align: center;
}

header {
  /* background-color: #f8f9fa; */
  padding: 1rem;
}

nav {
  display: flex;
  gap: 1rem;
  justify-content: center;
  grid-column: span 12;
}

main {
  flex: 1;
}

footer {
  /* background-color: #f1f3f5; */
  padding: 1rem;
}
</style>
