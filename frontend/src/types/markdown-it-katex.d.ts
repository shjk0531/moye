declare module 'markdown-it-katex' {
    import type { PluginWithOptions } from 'markdown-it';
    interface KatexOptions {
        throwOnError?: boolean;
        errorColor?: string;
    }
    const markdownItKatex: PluginWithOptions<KatexOptions>;
    export default markdownItKatex;
}
