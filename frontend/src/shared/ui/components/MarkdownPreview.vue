<template>
    <div class="markdown-body p-3" v-html="parsedHtml"></div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import MarkdownIt from 'markdown-it';
import hljs from 'highlight.js';

const md: MarkdownIt = new MarkdownIt({
    html: true,
    linkify: true,
    typographer: true,
    breaks: true,
    highlight: function (this: MarkdownIt, str: string, lang: string): string {
        if (lang && hljs.getLanguage(lang)) {
            try {
                return `<pre class="hljs"><code>${
                    hljs.highlight(str, {
                        language: lang,
                        ignoreIllegals: true,
                    }).value
                }</code></pre>`;
            } catch (_error) {}
        }
        return `<pre class="hljs"><code>${md.utils.escapeHtml(
            str,
        )}</code></pre>`;
    },
});

const props = defineProps<{ content: string }>();
const parsedHtml = computed(() => md.render(props.content));
</script>

<style lang="scss">
.markdown-body {
    font-family: ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont,
        'Segoe UI', Roboto, 'Helvetica Neue', Arial, 'Noto Sans', sans-serif,
        'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol',
        'Noto Color Emoji';
    color: #1f2937;
    line-height: 1.625;
}
.markdown-body h1 {
    margin-top: 1.5rem;
    margin-bottom: 0.5rem;
    font-size: 1.875rem;
    font-weight: 700;
}
.markdown-body h2 {
    margin-top: 1.25rem;
    margin-bottom: 0.5rem;
    font-size: 1.5rem;
    font-weight: 700;
}
.markdown-body h3 {
    margin-top: 1rem;
    margin-bottom: 0.5rem;
    font-size: 1.25rem;
    font-weight: 600;
}
.markdown-body h4 {
    margin-top: 0.75rem;
    margin-bottom: 0.5rem;
    font-size: 1.125rem;
    font-weight: 600;
}
.markdown-body h5,
.markdown-body h6 {
    margin-top: 0.5rem;
    margin-bottom: 0.5rem;
    font-size: 1rem;
    font-weight: 600;
}
.markdown-body p {
    margin: 0.5rem 0;
}
.markdown-body a {
    color: #2563eb;
    text-decoration: none;
}
.markdown-body a:hover {
    text-decoration: underline;
}
.markdown-body strong {
    font-weight: 700;
}
.markdown-body em {
    font-style: italic;
}
.markdown-body code {
    background-color: #f3f4f6;
    padding: 0.125rem 0.25rem;
    border-radius: 0.25rem;
    font-size: 0.875rem;
    font-family: Menlo, Monaco, 'Courier New', monospace;
}
.markdown-body pre {
    background-color: #f3f4f6;
    padding: 1rem;
    border-radius: 0.25rem;
    overflow-x: auto;
    margin: 1rem 0;
}
.markdown-body blockquote {
    border-left: 4px solid #d1d5db;
    padding-left: 1rem;
    font-style: italic;
    color: #4b5563;
    margin: 1rem 0;
}
.markdown-body hr {
    border-top: 1px solid #d1d5db;
    margin: 1.5rem 0;
}
.markdown-body ul,
.markdown-body ol {
    list-style-position: inside;
    margin: 0.5rem 0;
}
.markdown-body ul {
    list-style-type: disc;
}
.markdown-body ol {
    list-style-type: decimal;
}
.markdown-body li {
    margin: 0.25rem 0;
}
.markdown-body table {
    width: 100%;
    border-collapse: collapse;
    margin: 1rem 0;
}
.markdown-body th,
.markdown-body td {
    border: 1px solid #d1d5db;
    padding: 0.25rem 0.75rem;
}
.markdown-body thead {
    background-color: #f3f4f6;
}
.markdown-body img {
    max-width: 100%;
    border-radius: 0.25rem;
    margin: 1rem 0;
}
</style>

<!-- KaTeX math 스타일 추가 -->
<style lang="scss">
.markdown-body .katex {
    display: inline-block;
    vertical-align: middle;
    white-space: nowrap;
}
.markdown-body .katex-display {
    display: block;
    text-align: center;
    margin: 1em 0;
}
</style>

<!-- 다크모드 스타일 추가 -->
<style lang="scss">
@media (prefers-color-scheme: dark) {
    .markdown-body {
        color: #e5e7eb;
        background-color: transparent;
    }
    .markdown-body a {
        color: #3b82f6;
    }
    .markdown-body code {
        background-color: #1f2937;
        color: #e5e7eb;
    }
    .markdown-body pre {
        background-color: #1f2937;
        color: #e5e7eb;
    }
    .markdown-body blockquote {
        border-left-color: #4b5563;
        color: #9ca3af;
    }
    .markdown-body hr {
        border-color: #4b5563;
    }
    .markdown-body th,
    .markdown-body td {
        border-color: #4b5563;
    }
    .markdown-body thead {
        background-color: #1f2937;
    }
}

/* 애플리케이션 테마 클래스 기반 다크모드 */
:root.dark .markdown-body,
.dark .markdown-body {
    color: #e5e7eb;
    background-color: transparent;
}
:root.dark .markdown-body a,
.dark .markdown-body a {
    color: #3b82f6;
}
:root.dark .markdown-body code,
.dark .markdown-body code {
    background-color: #1f2937;
    color: #e5e7eb;
}
:root.dark .markdown-body pre,
.dark .markdown-body pre {
    background-color: #1f2937;
    color: #e5e7eb;
}
:root.dark .markdown-body blockquote,
.dark .markdown-body blockquote {
    border-left-color: #4b5563;
    color: #9ca3af;
}
:root.dark .markdown-body hr,
.dark .markdown-body hr {
    border-color: #4b5563;
}
:root.dark .markdown-body th,
:root.dark .markdown-body td,
.dark .markdown-body th,
.dark .markdown-body td {
    border-color: #4b5563;
}
:root.dark .markdown-body thead,
.dark .markdown-body thead {
    background-color: #1f2937;
}
</style>
