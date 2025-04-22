<!-- src/app/App.vue -->

<template>
    <Toast />
    <router-view />
</template>

<script setup lang="ts"></script>

<style>
#html,
body,
#app {
    margin: 0;
    font-family: Arial, sans-serif;
    width: 100vw;
    height: 100vh;
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
