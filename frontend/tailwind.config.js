module.exports = {
    content: ['./index.html', './**/*.{vue,js,ts,jsx,tsx}'],
    darkMode: 'class', // 다크 모드를 클래스 기반으로 활성화
    theme: {
        extend: {
            colors: {
                primary: 'var(--color-primary)',
                secondary: 'var(--color-secondary)',
            },
        },
    },
    plugins: [],
};
